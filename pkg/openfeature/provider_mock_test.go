// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/openfeature/provider.go

// Package openfeature is a generated GoMock package.
package openfeature

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeatureProvider is a mock of FeatureProvider interface.
type MockFeatureProvider struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureProviderMockRecorder
}

// MockFeatureProviderMockRecorder is the mock recorder for MockFeatureProvider.
type MockFeatureProviderMockRecorder struct {
	mock *MockFeatureProvider
}

// NewMockFeatureProvider creates a new mock instance.
func NewMockFeatureProvider(ctrl *gomock.Controller) *MockFeatureProvider {
	mock := &MockFeatureProvider{ctrl: ctrl}
	mock.recorder = &MockFeatureProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureProvider) EXPECT() *MockFeatureProviderMockRecorder {
	return m.recorder
}

// GetBooleanEvaluation mocks base method.
func (m *MockFeatureProvider) GetBooleanEvaluation(flag string, defaultValue bool, evalCtx EvaluationContext, options EvaluationOptions) BoolResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooleanEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(BoolResolutionDetail)
	return ret0
}

// GetBooleanEvaluation indicates an expected call of GetBooleanEvaluation.
func (mr *MockFeatureProviderMockRecorder) GetBooleanEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooleanEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).GetBooleanEvaluation), flag, defaultValue, evalCtx, options)
}

// GetNumberEvaluation mocks base method.
func (m *MockFeatureProvider) GetNumberEvaluation(flag string, defaultValue float64, evalCtx EvaluationContext, options EvaluationOptions) NumberResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumberEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(NumberResolutionDetail)
	return ret0
}

// GetNumberEvaluation indicates an expected call of GetNumberEvaluation.
func (mr *MockFeatureProviderMockRecorder) GetNumberEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumberEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).GetNumberEvaluation), flag, defaultValue, evalCtx, options)
}

// GetObjectEvaluation mocks base method.
func (m *MockFeatureProvider) GetObjectEvaluation(flag string, defaultValue interface{}, evalCtx EvaluationContext, options EvaluationOptions) ResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(ResolutionDetail)
	return ret0
}

// GetObjectEvaluation indicates an expected call of GetObjectEvaluation.
func (mr *MockFeatureProviderMockRecorder) GetObjectEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).GetObjectEvaluation), flag, defaultValue, evalCtx, options)
}

// GetStringEvaluation mocks base method.
func (m *MockFeatureProvider) GetStringEvaluation(flag, defaultValue string, evalCtx EvaluationContext, options EvaluationOptions) StringResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(StringResolutionDetail)
	return ret0
}

// GetStringEvaluation indicates an expected call of GetStringEvaluation.
func (mr *MockFeatureProviderMockRecorder) GetStringEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).GetStringEvaluation), flag, defaultValue, evalCtx, options)
}

// Metadata mocks base method.
func (m *MockFeatureProvider) Metadata() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Metadata indicates an expected call of Metadata.
func (mr *MockFeatureProviderMockRecorder) Metadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockFeatureProvider)(nil).Metadata))
}
