// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ProtonMail/proton-bridge/internal/bridge (interfaces: Configer,PreferenceProvider,PanicHandler,ClientManager,CredentialsStorer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	credentials "github.com/ProtonMail/proton-bridge/internal/bridge/credentials"
	pmapi "github.com/ProtonMail/proton-bridge/pkg/pmapi"
	gomock "github.com/golang/mock/gomock"
)

// MockConfiger is a mock of Configer interface
type MockConfiger struct {
	ctrl     *gomock.Controller
	recorder *MockConfigerMockRecorder
}

// MockConfigerMockRecorder is the mock recorder for MockConfiger
type MockConfigerMockRecorder struct {
	mock *MockConfiger
}

// NewMockConfiger creates a new mock instance
func NewMockConfiger(ctrl *gomock.Controller) *MockConfiger {
	mock := &MockConfiger{ctrl: ctrl}
	mock.recorder = &MockConfigerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfiger) EXPECT() *MockConfigerMockRecorder {
	return m.recorder
}

// ClearData mocks base method
func (m *MockConfiger) ClearData() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearData")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearData indicates an expected call of ClearData
func (mr *MockConfigerMockRecorder) ClearData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearData", reflect.TypeOf((*MockConfiger)(nil).ClearData))
}

// GetAPIConfig mocks base method
func (m *MockConfiger) GetAPIConfig() *pmapi.ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIConfig")
	ret0, _ := ret[0].(*pmapi.ClientConfig)
	return ret0
}

// GetAPIConfig indicates an expected call of GetAPIConfig
func (mr *MockConfigerMockRecorder) GetAPIConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIConfig", reflect.TypeOf((*MockConfiger)(nil).GetAPIConfig))
}

// GetDBDir mocks base method
func (m *MockConfiger) GetDBDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDBDir indicates an expected call of GetDBDir
func (mr *MockConfigerMockRecorder) GetDBDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBDir", reflect.TypeOf((*MockConfiger)(nil).GetDBDir))
}

// GetIMAPCachePath mocks base method
func (m *MockConfiger) GetIMAPCachePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIMAPCachePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIMAPCachePath indicates an expected call of GetIMAPCachePath
func (mr *MockConfigerMockRecorder) GetIMAPCachePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIMAPCachePath", reflect.TypeOf((*MockConfiger)(nil).GetIMAPCachePath))
}

// MockPreferenceProvider is a mock of PreferenceProvider interface
type MockPreferenceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockPreferenceProviderMockRecorder
}

// MockPreferenceProviderMockRecorder is the mock recorder for MockPreferenceProvider
type MockPreferenceProviderMockRecorder struct {
	mock *MockPreferenceProvider
}

// NewMockPreferenceProvider creates a new mock instance
func NewMockPreferenceProvider(ctrl *gomock.Controller) *MockPreferenceProvider {
	mock := &MockPreferenceProvider{ctrl: ctrl}
	mock.recorder = &MockPreferenceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPreferenceProvider) EXPECT() *MockPreferenceProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPreferenceProvider) Get(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockPreferenceProviderMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPreferenceProvider)(nil).Get), arg0)
}

// GetBool mocks base method
func (m *MockPreferenceProvider) GetBool(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBool", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool
func (mr *MockPreferenceProviderMockRecorder) GetBool(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockPreferenceProvider)(nil).GetBool), arg0)
}

// GetInt mocks base method
func (m *MockPreferenceProvider) GetInt(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt
func (mr *MockPreferenceProviderMockRecorder) GetInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockPreferenceProvider)(nil).GetInt), arg0)
}

// Set mocks base method
func (m *MockPreferenceProvider) Set(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set
func (mr *MockPreferenceProviderMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockPreferenceProvider)(nil).Set), arg0, arg1)
}

// MockPanicHandler is a mock of PanicHandler interface
type MockPanicHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPanicHandlerMockRecorder
}

// MockPanicHandlerMockRecorder is the mock recorder for MockPanicHandler
type MockPanicHandlerMockRecorder struct {
	mock *MockPanicHandler
}

// NewMockPanicHandler creates a new mock instance
func NewMockPanicHandler(ctrl *gomock.Controller) *MockPanicHandler {
	mock := &MockPanicHandler{ctrl: ctrl}
	mock.recorder = &MockPanicHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPanicHandler) EXPECT() *MockPanicHandlerMockRecorder {
	return m.recorder
}

// HandlePanic mocks base method
func (m *MockPanicHandler) HandlePanic() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandlePanic")
}

// HandlePanic indicates an expected call of HandlePanic
func (mr *MockPanicHandlerMockRecorder) HandlePanic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlePanic", reflect.TypeOf((*MockPanicHandler)(nil).HandlePanic))
}

// MockClientManager is a mock of ClientManager interface
type MockClientManager struct {
	ctrl     *gomock.Controller
	recorder *MockClientManagerMockRecorder
}

// MockClientManagerMockRecorder is the mock recorder for MockClientManager
type MockClientManagerMockRecorder struct {
	mock *MockClientManager
}

// NewMockClientManager creates a new mock instance
func NewMockClientManager(ctrl *gomock.Controller) *MockClientManager {
	mock := &MockClientManager{ctrl: ctrl}
	mock.recorder = &MockClientManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientManager) EXPECT() *MockClientManagerMockRecorder {
	return m.recorder
}

// AllowProxy mocks base method
func (m *MockClientManager) AllowProxy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AllowProxy")
}

// AllowProxy indicates an expected call of AllowProxy
func (mr *MockClientManagerMockRecorder) AllowProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowProxy", reflect.TypeOf((*MockClientManager)(nil).AllowProxy))
}

// DisallowProxy mocks base method
func (m *MockClientManager) DisallowProxy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisallowProxy")
}

// DisallowProxy indicates an expected call of DisallowProxy
func (mr *MockClientManagerMockRecorder) DisallowProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisallowProxy", reflect.TypeOf((*MockClientManager)(nil).DisallowProxy))
}

// GetAnonymousClient mocks base method
func (m *MockClientManager) GetAnonymousClient() pmapi.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnonymousClient")
	ret0, _ := ret[0].(pmapi.Client)
	return ret0
}

// GetAnonymousClient indicates an expected call of GetAnonymousClient
func (mr *MockClientManagerMockRecorder) GetAnonymousClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnonymousClient", reflect.TypeOf((*MockClientManager)(nil).GetAnonymousClient))
}

// GetBridgeAuthChannel mocks base method
func (m *MockClientManager) GetBridgeAuthChannel() chan pmapi.ClientAuth {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBridgeAuthChannel")
	ret0, _ := ret[0].(chan pmapi.ClientAuth)
	return ret0
}

// GetBridgeAuthChannel indicates an expected call of GetBridgeAuthChannel
func (mr *MockClientManagerMockRecorder) GetBridgeAuthChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBridgeAuthChannel", reflect.TypeOf((*MockClientManager)(nil).GetBridgeAuthChannel))
}

// GetClient mocks base method
func (m *MockClientManager) GetClient(arg0 string) pmapi.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", arg0)
	ret0, _ := ret[0].(pmapi.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient
func (mr *MockClientManagerMockRecorder) GetClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockClientManager)(nil).GetClient), arg0)
}

// MockCredentialsStorer is a mock of CredentialsStorer interface
type MockCredentialsStorer struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialsStorerMockRecorder
}

// MockCredentialsStorerMockRecorder is the mock recorder for MockCredentialsStorer
type MockCredentialsStorerMockRecorder struct {
	mock *MockCredentialsStorer
}

// NewMockCredentialsStorer creates a new mock instance
func NewMockCredentialsStorer(ctrl *gomock.Controller) *MockCredentialsStorer {
	mock := &MockCredentialsStorer{ctrl: ctrl}
	mock.recorder = &MockCredentialsStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCredentialsStorer) EXPECT() *MockCredentialsStorerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockCredentialsStorer) Add(arg0, arg1, arg2, arg3 string, arg4 []string) (*credentials.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*credentials.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockCredentialsStorerMockRecorder) Add(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCredentialsStorer)(nil).Add), arg0, arg1, arg2, arg3, arg4)
}

// Delete mocks base method
func (m *MockCredentialsStorer) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCredentialsStorerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCredentialsStorer)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockCredentialsStorer) Get(arg0 string) (*credentials.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*credentials.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCredentialsStorerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCredentialsStorer)(nil).Get), arg0)
}

// List mocks base method
func (m *MockCredentialsStorer) List() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCredentialsStorerMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCredentialsStorer)(nil).List))
}

// Logout mocks base method
func (m *MockCredentialsStorer) Logout(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockCredentialsStorerMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockCredentialsStorer)(nil).Logout), arg0)
}

// SwitchAddressMode mocks base method
func (m *MockCredentialsStorer) SwitchAddressMode(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwitchAddressMode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwitchAddressMode indicates an expected call of SwitchAddressMode
func (mr *MockCredentialsStorerMockRecorder) SwitchAddressMode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchAddressMode", reflect.TypeOf((*MockCredentialsStorer)(nil).SwitchAddressMode), arg0)
}

// UpdateEmails mocks base method
func (m *MockCredentialsStorer) UpdateEmails(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmails indicates an expected call of UpdateEmails
func (mr *MockCredentialsStorerMockRecorder) UpdateEmails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmails", reflect.TypeOf((*MockCredentialsStorer)(nil).UpdateEmails), arg0, arg1)
}

// UpdatePassword mocks base method
func (m *MockCredentialsStorer) UpdatePassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockCredentialsStorerMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockCredentialsStorer)(nil).UpdatePassword), arg0, arg1)
}

// UpdateToken mocks base method
func (m *MockCredentialsStorer) UpdateToken(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken
func (mr *MockCredentialsStorerMockRecorder) UpdateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockCredentialsStorer)(nil).UpdateToken), arg0, arg1)
}
