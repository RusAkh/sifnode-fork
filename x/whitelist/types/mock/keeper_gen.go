// Code generated by MockGen. DO NOT EDIT.
// Source: ../expected_keepers.go

// Package whitelistmocks is a generated GoMock package.
package whitelistmocks

import (
	reflect "reflect"

	types "github.com/Sifchain/sifnode/x/whitelist/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types1 "github.com/tendermint/tendermint/abci/types"
)

// MockKeeper is a mock of Keeper interface.
type MockKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockKeeperMockRecorder
}

// MockKeeperMockRecorder is the mock recorder for MockKeeper.
type MockKeeperMockRecorder struct {
	mock *MockKeeper
}

// NewMockKeeper creates a new mock instance.
func NewMockKeeper(ctrl *gomock.Controller) *MockKeeper {
	mock := &MockKeeper{ctrl: ctrl}
	mock.recorder = &MockKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeeper) EXPECT() *MockKeeperMockRecorder {
	return m.recorder
}

// ExportGenesis mocks base method.
func (m *MockKeeper) ExportGenesis(ctx types0.Context) *types.GenesisState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", ctx)
	ret0, _ := ret[0].(*types.GenesisState)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis.
func (mr *MockKeeperMockRecorder) ExportGenesis(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockKeeper)(nil).ExportGenesis), ctx)
}

// GetDenom mocks base method.
func (m *MockKeeper) GetDenom(ctx types0.Context, denom string) types.DenomWhitelistEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenom", ctx, denom)
	ret0, _ := ret[0].(types.DenomWhitelistEntry)
	return ret0
}

// GetDenom indicates an expected call of GetDenom.
func (mr *MockKeeperMockRecorder) GetDenom(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenom", reflect.TypeOf((*MockKeeper)(nil).GetDenom), ctx, denom)
}

// InitGenesis mocks base method.
func (m *MockKeeper) InitGenesis(ctx types0.Context, state types.GenesisState) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", ctx, state)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis.
func (mr *MockKeeperMockRecorder) InitGenesis(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockKeeper)(nil).InitGenesis), ctx, state)
}

// IsAdminAccount mocks base method.
func (m *MockKeeper) IsAdminAccount(ctx types0.Context, adminAccount types0.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAdminAccount", ctx, adminAccount)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAdminAccount indicates an expected call of IsAdminAccount.
func (mr *MockKeeperMockRecorder) IsAdminAccount(ctx, adminAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdminAccount", reflect.TypeOf((*MockKeeper)(nil).IsAdminAccount), ctx, adminAccount)
}

// IsDenomWhitelisted mocks base method.
func (m *MockKeeper) IsDenomWhitelisted(ctx types0.Context, denom string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDenomWhitelisted", ctx, denom)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDenomWhitelisted indicates an expected call of IsDenomWhitelisted.
func (mr *MockKeeperMockRecorder) IsDenomWhitelisted(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDenomWhitelisted", reflect.TypeOf((*MockKeeper)(nil).IsDenomWhitelisted), ctx, denom)
}

// SetDenom mocks base method.
func (m *MockKeeper) SetDenom(ctx types0.Context, denom string, exp int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDenom", ctx, denom, exp)
}

// SetDenom indicates an expected call of SetDenom.
func (mr *MockKeeperMockRecorder) SetDenom(ctx, denom, exp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDenom", reflect.TypeOf((*MockKeeper)(nil).SetDenom), ctx, denom, exp)
}
