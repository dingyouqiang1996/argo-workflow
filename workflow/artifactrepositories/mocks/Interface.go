// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	context "context"

	config "github.com/argoproj/argo-workflows/v3/config"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, ref
func (_m *Interface) Get(ctx context.Context, ref *v1alpha1.ArtifactRepositoryRefStatus) (*config.ArtifactRepository, error) {
	ret := _m.Called(ctx, ref)

	var r0 *config.ArtifactRepository
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ArtifactRepositoryRefStatus) *config.ArtifactRepository); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.ArtifactRepository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.ArtifactRepositoryRefStatus) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resolve provides a mock function with given fields: ctx, ref, workflowNamespace
func (_m *Interface) Resolve(ctx context.Context, ref *v1alpha1.ArtifactRepositoryRef, workflowNamespace string) (*v1alpha1.ArtifactRepositoryRefStatus, error) {
	ret := _m.Called(ctx, ref, workflowNamespace)

	var r0 *v1alpha1.ArtifactRepositoryRefStatus
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ArtifactRepositoryRef, string) *v1alpha1.ArtifactRepositoryRefStatus); ok {
		r0 = rf(ctx, ref, workflowNamespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ArtifactRepositoryRefStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.ArtifactRepositoryRef, string) error); ok {
		r1 = rf(ctx, ref, workflowNamespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
