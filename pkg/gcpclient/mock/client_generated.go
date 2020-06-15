// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	gcpclient "github.com/openshift/hive/pkg/gcpclient"
	compute "google.golang.org/api/compute/v1"
	dns "google.golang.org/api/dns/v1"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListManagedZones mocks base method
func (m *MockClient) ListManagedZones(opts gcpclient.ListManagedZonesOptions) (*dns.ManagedZonesListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListManagedZones", opts)
	ret0, _ := ret[0].(*dns.ManagedZonesListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListManagedZones indicates an expected call of ListManagedZones
func (mr *MockClientMockRecorder) ListManagedZones(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedZones", reflect.TypeOf((*MockClient)(nil).ListManagedZones), opts)
}

// ListResourceRecordSets mocks base method
func (m *MockClient) ListResourceRecordSets(managedZone string, opts gcpclient.ListResourceRecordSetsOptions) (*dns.ResourceRecordSetsListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRecordSets", managedZone, opts)
	ret0, _ := ret[0].(*dns.ResourceRecordSetsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRecordSets indicates an expected call of ListResourceRecordSets
func (mr *MockClientMockRecorder) ListResourceRecordSets(managedZone, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ListResourceRecordSets), managedZone, opts)
}

// AddResourceRecordSet mocks base method
func (m *MockClient) AddResourceRecordSet(managedZone string, recordSet *dns.ResourceRecordSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResourceRecordSet", managedZone, recordSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResourceRecordSet indicates an expected call of AddResourceRecordSet
func (mr *MockClientMockRecorder) AddResourceRecordSet(managedZone, recordSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResourceRecordSet", reflect.TypeOf((*MockClient)(nil).AddResourceRecordSet), managedZone, recordSet)
}

// DeleteResourceRecordSet mocks base method
func (m *MockClient) DeleteResourceRecordSet(managedZone string, recordSet *dns.ResourceRecordSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResourceRecordSet", managedZone, recordSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResourceRecordSet indicates an expected call of DeleteResourceRecordSet
func (mr *MockClientMockRecorder) DeleteResourceRecordSet(managedZone, recordSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceRecordSet", reflect.TypeOf((*MockClient)(nil).DeleteResourceRecordSet), managedZone, recordSet)
}

// UpdateResourceRecordSet mocks base method
func (m *MockClient) UpdateResourceRecordSet(managedZone string, addRecordSet, removeRecordSet *dns.ResourceRecordSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResourceRecordSet", managedZone, addRecordSet, removeRecordSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResourceRecordSet indicates an expected call of UpdateResourceRecordSet
func (mr *MockClientMockRecorder) UpdateResourceRecordSet(managedZone, addRecordSet, removeRecordSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceRecordSet", reflect.TypeOf((*MockClient)(nil).UpdateResourceRecordSet), managedZone, addRecordSet, removeRecordSet)
}

// GetManagedZone mocks base method
func (m *MockClient) GetManagedZone(managedZone string) (*dns.ManagedZone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedZone", managedZone)
	ret0, _ := ret[0].(*dns.ManagedZone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedZone indicates an expected call of GetManagedZone
func (mr *MockClientMockRecorder) GetManagedZone(managedZone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedZone", reflect.TypeOf((*MockClient)(nil).GetManagedZone), managedZone)
}

// CreateManagedZone mocks base method
func (m *MockClient) CreateManagedZone(managedZone *dns.ManagedZone) (*dns.ManagedZone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateManagedZone", managedZone)
	ret0, _ := ret[0].(*dns.ManagedZone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateManagedZone indicates an expected call of CreateManagedZone
func (mr *MockClientMockRecorder) CreateManagedZone(managedZone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateManagedZone", reflect.TypeOf((*MockClient)(nil).CreateManagedZone), managedZone)
}

// DeleteManagedZone mocks base method
func (m *MockClient) DeleteManagedZone(managedZone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteManagedZone", managedZone)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteManagedZone indicates an expected call of DeleteManagedZone
func (mr *MockClientMockRecorder) DeleteManagedZone(managedZone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteManagedZone", reflect.TypeOf((*MockClient)(nil).DeleteManagedZone), managedZone)
}

// ListComputeZones mocks base method
func (m *MockClient) ListComputeZones(arg0 gcpclient.ListComputeZonesOptions) (*compute.ZoneList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComputeZones", arg0)
	ret0, _ := ret[0].(*compute.ZoneList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComputeZones indicates an expected call of ListComputeZones
func (mr *MockClientMockRecorder) ListComputeZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComputeZones", reflect.TypeOf((*MockClient)(nil).ListComputeZones), arg0)
}

// ListComputeImages mocks base method
func (m *MockClient) ListComputeImages(arg0 gcpclient.ListComputeImagesOptions) (*compute.ImageList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComputeImages", arg0)
	ret0, _ := ret[0].(*compute.ImageList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComputeImages indicates an expected call of ListComputeImages
func (mr *MockClientMockRecorder) ListComputeImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComputeImages", reflect.TypeOf((*MockClient)(nil).ListComputeImages), arg0)
}
