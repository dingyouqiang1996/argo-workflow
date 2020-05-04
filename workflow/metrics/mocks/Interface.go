// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/argoproj/argo/workflow/common"

	mock "github.com/stretchr/testify/mock"

	prometheus "github.com/prometheus/client_golang/prometheus"

	time "time"

	v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Collect provides a mock function with given fields: _a0
func (_m *Interface) Collect(_a0 chan<- prometheus.Metric) {
	_m.Called(_a0)
}

// DeleteExpiredMetrics provides a mock function with given fields: ttl
func (_m *Interface) DeleteExpiredMetrics(ttl time.Duration) {
	_m.Called(ttl)
}

// Describe provides a mock function with given fields: _a0
func (_m *Interface) Describe(_a0 chan<- *prometheus.Desc) {
	_m.Called(_a0)
}

// GetCustom provides a mock function with given fields: desc
func (_m *Interface) GetCustom(desc string) prometheus.Metric {
	ret := _m.Called(desc)

	var r0 prometheus.Metric
	if rf, ok := ret.Get(0).(func(string) prometheus.Metric); ok {
		r0 = rf(desc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(prometheus.Metric)
		}
	}

	return r0
}

// PodChanged provides a mock function with given fields: significant
func (_m *Interface) PodChanged(significant bool) {
	_m.Called(significant)
}

// PodProcessed provides a mock function with given fields:
func (_m *Interface) PodProcessed() {
	_m.Called()
}

// PodResourceVersionRepeated provides a mock function with given fields:
func (_m *Interface) PodResourceVersionRepeated() {
	_m.Called()
}

// SetCustom provides a mock function with given fields: desc, metric
func (_m *Interface) SetCustom(desc string, metric common.Metric) {
	_m.Called(desc, metric)
}

// UpdatesPersisted provides a mock function with given fields:
func (_m *Interface) UpdatesPersisted() {
	_m.Called()
}

// UpdatesReapplied provides a mock function with given fields:
func (_m *Interface) UpdatesReapplied() {
	_m.Called()
}

// WorkflowAdded provides a mock function with given fields: phase
func (_m *Interface) WorkflowAdded(phase v1alpha1.NodePhase) {
	_m.Called(phase)
}

// WorkflowDeleted provides a mock function with given fields: phase
func (_m *Interface) WorkflowDeleted(phase v1alpha1.NodePhase) {
	_m.Called(phase)
}

// WorkflowProcessed provides a mock function with given fields:
func (_m *Interface) WorkflowProcessed() {
	_m.Called()
}

// WorkflowResourceVersionRepeated provides a mock function with given fields:
func (_m *Interface) WorkflowResourceVersionRepeated() {
	_m.Called()
}

// WorkflowUpdated provides a mock function with given fields: from, to
func (_m *Interface) WorkflowUpdated(from v1alpha1.NodePhase, to v1alpha1.NodePhase) {
	_m.Called(from, to)
}
