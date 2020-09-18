package controller

import (
	"testing"

	"github.com/argoproj/argo/workflow/sync"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/config"
)

func TestUpdateConfig(t *testing.T) {
	cancel, controller := newController()
	defer cancel()
	controller.throttler = sync.NewThrottler(0, func(string) {})
	err := controller.updateConfig(config.Config{ExecutorImage: "argoexec:latest"})
	assert.NoError(t, err)
	assert.NotNil(t, controller.Config)
	assert.NotNil(t, controller.archiveLabelSelector)
	assert.NotNil(t, controller.wfArchive)
	assert.NotNil(t, controller.offloadNodeStatusRepo)
}
