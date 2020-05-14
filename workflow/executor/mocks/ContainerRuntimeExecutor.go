// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"

// ContainerRuntimeExecutor is an autogenerated mock type for the ContainerRuntimeExecutor type
type ContainerRuntimeExecutor struct {
	mock.Mock
}

// CopyFile provides a mock function with given fields: containerID, sourcePath, destPath, compressionLevel
func (_m *ContainerRuntimeExecutor) CopyFile(containerID string, sourcePath string, destPath string, compressionLevel int) error {
	ret := _m.Called(containerID, sourcePath, destPath, compressionLevel)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int) error); ok {
		r0 = rf(containerID, sourcePath, destPath, compressionLevel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExitCode provides a mock function with given fields: containerID
func (_m *ContainerRuntimeExecutor) GetExitCode(containerID string) (string, error) {
	ret := _m.Called(containerID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(containerID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(containerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFileContents provides a mock function with given fields: containerID, sourcePath
func (_m *ContainerRuntimeExecutor) GetFileContents(containerID string, sourcePath string) (string, error) {
	ret := _m.Called(containerID, sourcePath)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(containerID, sourcePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(containerID, sourcePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOutputStream provides a mock function with given fields: containerID, combinedOutput
func (_m *ContainerRuntimeExecutor) GetOutputStream(containerID string, combinedOutput bool) (io.ReadCloser, error) {
	ret := _m.Called(containerID, combinedOutput)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, bool) io.ReadCloser); ok {
		r0 = rf(containerID, combinedOutput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(containerID, combinedOutput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Kill provides a mock function with given fields: containerIDs
func (_m *ContainerRuntimeExecutor) Kill(containerIDs []string) error {
	ret := _m.Called(containerIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(containerIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wait provides a mock function with given fields: containerID
func (_m *ContainerRuntimeExecutor) Wait(containerID string) error {
	ret := _m.Called(containerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(containerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitInit provides a mock function with given fields:
func (_m *ContainerRuntimeExecutor) WaitInit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
