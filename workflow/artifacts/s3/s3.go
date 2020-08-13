package s3

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/argoproj/pkg/file"
	argos3 "github.com/argoproj/pkg/s3"

	"github.com/argoproj/argo/errors"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/workflow/common"
)

// S3ArtifactDriver is a driver for AWS S3
type S3ArtifactDriver struct {
	Endpoint    string
	Region      string
	Secure      bool
	AccessKey   string
	SecretKey   string
	RoleARN     string
	UseSDKCreds bool
	Context     context.Context
}

// newMinioClient instantiates a new minio client object.
func (s3Driver *S3ArtifactDriver) newS3Client(ctx context.Context) (argos3.S3Client, error) {
	opts := argos3.S3ClientOpts{
		Endpoint:    s3Driver.Endpoint,
		Region:      s3Driver.Region,
		Secure:      s3Driver.Secure,
		AccessKey:   s3Driver.AccessKey,
		SecretKey:   s3Driver.SecretKey,
		RoleARN:     s3Driver.RoleARN,
		Trace:       os.Getenv(common.EnvVarArgoTrace) == "1",
		UseSDKCreds: s3Driver.UseSDKCreds,
	}
	return argos3.NewS3Client(ctx, opts)
}

// Load downloads artifacts from S3 compliant storage
func (s3Driver *S3ArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := wait.ExponentialBackoff(wait.Backoff{Duration: time.Second * 2, Factor: 2.0, Steps: 5, Jitter: 0.1},
		func() (bool, error) {
			log.WithFields(log.Fields{"path": path, "key": inputArtifact.S3.Key}).Info("S3 Load")
			s3cli, err := s3Driver.newS3Client(ctx)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Warn("Failed to create new S3 client")
				return false, nil
			}
			origErr := s3cli.GetFile(inputArtifact.S3.Bucket, inputArtifact.S3.Key, path)
			if origErr == nil {
				return true, nil
			}
			if !argos3.IsS3ErrCode(origErr, "NoSuchKey") {
				log.WithFields(log.Fields{"error": origErr}).Warn("Failed get file")
				return false, nil
			}
			// If we get here, the error was a NoSuchKey. The key might be a s3 "directory"
			isDir, err := s3cli.IsDirectory(inputArtifact.S3.Bucket, inputArtifact.S3.Key)
			if err != nil {
				log.WithFields(log.Fields{"path": inputArtifact.S3.Bucket, "error": err}).Warn("Failed to test if the path is a directory")
				return false, nil
			}
			if !isDir {
				// It's neither a file, nor a directory. Return the original NoSuchKey error
				return false, errors.New(errors.CodeNotFound, origErr.Error())
			}

			if err = s3cli.GetDirectory(inputArtifact.S3.Bucket, inputArtifact.S3.Key, path); err != nil {
				log.WithFields(log.Fields{"error": err}).Warn("Failed get directory")
				return false, nil
			}
			return true, nil
		})

	return err
}

// Save saves an artifact to S3 compliant storage
func (s3Driver *S3ArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := wait.ExponentialBackoff(wait.Backoff{Duration: time.Second * 2, Factor: 2.0, Steps: 5, Jitter: 0.1},
		func() (bool, error) {
			log.WithFields(log.Fields{"path": path, "key": outputArtifact.S3.Key}).Info("S3 Save")
			s3cli, err := s3Driver.newS3Client(ctx)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Warn("Failed to create new S3 client")
				return false, nil
			}
			isDir, err := file.IsDirectory(path)
			if err != nil {
				log.WithFields(log.Fields{"path": path, "error": err}).Warn("Failed to test if the path is a directory")
				return false, nil
			}
			if isDir {
				if err = s3cli.PutDirectory(outputArtifact.S3.Bucket, outputArtifact.S3.Key, path); err != nil {
					log.WithFields(log.Fields{"error": err}).Warn("Failed to put directory")
					return false, nil
				}
			} else {
				if err = s3cli.PutFile(outputArtifact.S3.Bucket, outputArtifact.S3.Key, path); err != nil {
					log.WithFields(log.Fields{"error": err}).Warn("Failed to put file")
					return false, nil
				}
			}
			return true, nil
		})
	return err
}
