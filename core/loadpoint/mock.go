// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/evcc-io/evcc/core/loadpoint (interfaces: API)
//
// Generated by this command:
//
//	mockgen -package loadpoint -destination mock.go -mock_names API=MockAPI github.com/evcc-io/evcc/core/loadpoint API
//

// Package loadpoint is a generated GoMock package.
package loadpoint

import (
	reflect "reflect"
	time "time"

	api "github.com/evcc-io/evcc/api"
	gomock "go.uber.org/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ActivePhases mocks base method.
func (m *MockAPI) ActivePhases() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivePhases")
	ret0, _ := ret[0].(int)
	return ret0
}

// ActivePhases indicates an expected call of ActivePhases.
func (mr *MockAPIMockRecorder) ActivePhases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivePhases", reflect.TypeOf((*MockAPI)(nil).ActivePhases))
}

// EffectiveMaxPower mocks base method.
func (m *MockAPI) EffectiveMaxPower() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EffectiveMaxPower")
	ret0, _ := ret[0].(float64)
	return ret0
}

// EffectiveMaxPower indicates an expected call of EffectiveMaxPower.
func (mr *MockAPIMockRecorder) EffectiveMaxPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EffectiveMaxPower", reflect.TypeOf((*MockAPI)(nil).EffectiveMaxPower))
}

// EffectiveMinPower mocks base method.
func (m *MockAPI) EffectiveMinPower() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EffectiveMinPower")
	ret0, _ := ret[0].(float64)
	return ret0
}

// EffectiveMinPower indicates an expected call of EffectiveMinPower.
func (mr *MockAPIMockRecorder) EffectiveMinPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EffectiveMinPower", reflect.TypeOf((*MockAPI)(nil).EffectiveMinPower))
}

// EffectivePlanTime mocks base method.
func (m *MockAPI) EffectivePlanTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EffectivePlanTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// EffectivePlanTime indicates an expected call of EffectivePlanTime.
func (mr *MockAPIMockRecorder) EffectivePlanTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EffectivePlanTime", reflect.TypeOf((*MockAPI)(nil).EffectivePlanTime))
}

// EffectivePriority mocks base method.
func (m *MockAPI) EffectivePriority() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EffectivePriority")
	ret0, _ := ret[0].(int)
	return ret0
}

// EffectivePriority indicates an expected call of EffectivePriority.
func (mr *MockAPIMockRecorder) EffectivePriority() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EffectivePriority", reflect.TypeOf((*MockAPI)(nil).EffectivePriority))
}

// GetChargePower mocks base method.
func (m *MockAPI) GetChargePower() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChargePower")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetChargePower indicates an expected call of GetChargePower.
func (mr *MockAPIMockRecorder) GetChargePower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChargePower", reflect.TypeOf((*MockAPI)(nil).GetChargePower))
}

// GetChargePowerFlexibility mocks base method.
func (m *MockAPI) GetChargePowerFlexibility() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChargePowerFlexibility")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetChargePowerFlexibility indicates an expected call of GetChargePowerFlexibility.
func (mr *MockAPIMockRecorder) GetChargePowerFlexibility() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChargePowerFlexibility", reflect.TypeOf((*MockAPI)(nil).GetChargePowerFlexibility))
}

// GetCharger mocks base method.
func (m *MockAPI) GetCharger() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharger")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCharger indicates an expected call of GetCharger.
func (mr *MockAPIMockRecorder) GetCharger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharger", reflect.TypeOf((*MockAPI)(nil).GetCharger))
}

// GetDefaultVehicle mocks base method.
func (m *MockAPI) GetDefaultVehicle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultVehicle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDefaultVehicle indicates an expected call of GetDefaultVehicle.
func (mr *MockAPIMockRecorder) GetDefaultVehicle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultVehicle", reflect.TypeOf((*MockAPI)(nil).GetDefaultVehicle))
}

// GetCircuit mocks base method.
func (m *MockAPI) GetCircuit() api.Circuit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCircuit")
	ret0, _ := ret[0].(api.Circuit)
	return ret0
}

// GetCircuit indicates an expected call of GetCircuit.
func (mr *MockAPIMockRecorder) GetCircuit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCircuit", reflect.TypeOf((*MockAPI)(nil).GetCircuit))
}

// GetDisableThreshold mocks base method.
func (m *MockAPI) GetDisableThreshold() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDisableThreshold")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetDisableThreshold indicates an expected call of GetDisableThreshold.
func (mr *MockAPIMockRecorder) GetDisableThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisableThreshold", reflect.TypeOf((*MockAPI)(nil).GetDisableThreshold))
}

// GetEnableThreshold mocks base method.
func (m *MockAPI) GetEnableThreshold() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnableThreshold")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetEnableThreshold indicates an expected call of GetEnableThreshold.
func (mr *MockAPIMockRecorder) GetEnableThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnableThreshold", reflect.TypeOf((*MockAPI)(nil).GetEnableThreshold))
}

// GetLimitEnergy mocks base method.
func (m *MockAPI) GetLimitEnergy() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimitEnergy")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetLimitEnergy indicates an expected call of GetLimitEnergy.
func (mr *MockAPIMockRecorder) GetLimitEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimitEnergy", reflect.TypeOf((*MockAPI)(nil).GetLimitEnergy))
}

// GetLimitSoc mocks base method.
func (m *MockAPI) GetLimitSoc() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimitSoc")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetLimitSoc indicates an expected call of GetLimitSoc.
func (mr *MockAPIMockRecorder) GetLimitSoc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimitSoc", reflect.TypeOf((*MockAPI)(nil).GetLimitSoc))
}

// GetMaxCurrent mocks base method.
func (m *MockAPI) GetMaxCurrent() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxCurrent")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetMaxCurrent indicates an expected call of GetMaxCurrent.
func (mr *MockAPIMockRecorder) GetMaxCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxCurrent", reflect.TypeOf((*MockAPI)(nil).GetMaxCurrent))
}

// GetMeter mocks base method.
func (m *MockAPI) GetMeter() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeter")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMeter indicates an expected call of GetMeter.
func (mr *MockAPIMockRecorder) GetMeter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeter", reflect.TypeOf((*MockAPI)(nil).GetMeter))
}

// GetMaxPhaseCurrent mocks base method.
func (m *MockAPI) GetMaxPhaseCurrent() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxPhaseCurrent")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetMaxPhaseCurrent indicates an expected call of GetMaxPhaseCurrent.
func (mr *MockAPIMockRecorder) GetMaxPhaseCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxPhaseCurrent", reflect.TypeOf((*MockAPI)(nil).GetMaxPhaseCurrent))
}

// GetMinCurrent mocks base method.
func (m *MockAPI) GetMinCurrent() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinCurrent")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetMinCurrent indicates an expected call of GetMinCurrent.
func (mr *MockAPIMockRecorder) GetMinCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinCurrent", reflect.TypeOf((*MockAPI)(nil).GetMinCurrent))
}

// GetMode mocks base method.
func (m *MockAPI) GetMode() api.ChargeMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMode")
	ret0, _ := ret[0].(api.ChargeMode)
	return ret0
}

// GetMode indicates an expected call of GetMode.
func (mr *MockAPIMockRecorder) GetMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMode", reflect.TypeOf((*MockAPI)(nil).GetMode))
}

// GetPhases mocks base method.
func (m *MockAPI) GetPhases() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhases")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPhases indicates an expected call of GetPhases.
func (mr *MockAPIMockRecorder) GetPhases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhases", reflect.TypeOf((*MockAPI)(nil).GetPhases))
}

// GetPlan mocks base method.
func (m *MockAPI) GetPlan(arg0 time.Time, arg1 time.Duration) (api.Rates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlan", arg0, arg1)
	ret0, _ := ret[0].(api.Rates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlan indicates an expected call of GetPlan.
func (mr *MockAPIMockRecorder) GetPlan(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlan", reflect.TypeOf((*MockAPI)(nil).GetPlan), arg0, arg1)
}

// GetPlanEnergy mocks base method.
func (m *MockAPI) GetPlanEnergy() (time.Time, float64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanEnergy")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// GetPlanEnergy indicates an expected call of GetPlanEnergy.
func (mr *MockAPIMockRecorder) GetPlanEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanEnergy", reflect.TypeOf((*MockAPI)(nil).GetPlanEnergy))
}

// GetPlanGoal mocks base method.
func (m *MockAPI) GetPlanGoal() (float64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanGoal")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPlanGoal indicates an expected call of GetPlanGoal.
func (mr *MockAPIMockRecorder) GetPlanGoal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanGoal", reflect.TypeOf((*MockAPI)(nil).GetPlanGoal))
}

// GetPlanRequiredDuration mocks base method.
func (m *MockAPI) GetPlanRequiredDuration(arg0, arg1 float64) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanRequiredDuration", arg0, arg1)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetPlanRequiredDuration indicates an expected call of GetPlanRequiredDuration.
func (mr *MockAPIMockRecorder) GetPlanRequiredDuration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanRequiredDuration", reflect.TypeOf((*MockAPI)(nil).GetPlanRequiredDuration), arg0, arg1)
}

// GetPriority mocks base method.
func (m *MockAPI) GetPriority() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriority")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPriority indicates an expected call of GetPriority.
func (mr *MockAPIMockRecorder) GetPriority() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriority", reflect.TypeOf((*MockAPI)(nil).GetPriority))
}

// GetRemainingDuration mocks base method.
func (m *MockAPI) GetRemainingDuration() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemainingDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetRemainingDuration indicates an expected call of GetRemainingDuration.
func (mr *MockAPIMockRecorder) GetRemainingDuration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemainingDuration", reflect.TypeOf((*MockAPI)(nil).GetRemainingDuration))
}

// GetRemainingEnergy mocks base method.
func (m *MockAPI) GetRemainingEnergy() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemainingEnergy")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetRemainingEnergy indicates an expected call of GetRemainingEnergy.
func (mr *MockAPIMockRecorder) GetRemainingEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemainingEnergy", reflect.TypeOf((*MockAPI)(nil).GetRemainingEnergy))
}

// GetSmartCostLimit mocks base method.
func (m *MockAPI) GetSmartCostLimit() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSmartCostLimit")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetSmartCostLimit indicates an expected call of GetSmartCostLimit.
func (mr *MockAPIMockRecorder) GetSmartCostLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSmartCostLimit", reflect.TypeOf((*MockAPI)(nil).GetSmartCostLimit))
}

// GetSocConfig mocks base method.
func (m *MockAPI) GetSocConfig() SocConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSocConfig")
	ret0, _ := ret[0].(SocConfig)
	return ret0
}

// GetSocConfig indicates an expected call of GetSocConfig.
func (mr *MockAPIMockRecorder) GetSocConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSocConfig", reflect.TypeOf((*MockAPI)(nil).GetSocConfig))
}

// GetStatus mocks base method.
func (m *MockAPI) GetStatus() api.ChargeStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(api.ChargeStatus)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockAPIMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockAPI)(nil).GetStatus))
}

// GetThresholds mocks base method.
func (m *MockAPI) GetThresholds() ThresholdsConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThresholds")
	ret0, _ := ret[0].(ThresholdsConfig)
	return ret0
}

// GetThresholds indicates an expected call of GetThresholds.
func (mr *MockAPIMockRecorder) GetThresholds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThresholds", reflect.TypeOf((*MockAPI)(nil).GetThresholds))
}

// GetTitle mocks base method.
func (m *MockAPI) GetTitle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTitle indicates an expected call of GetTitle.
func (mr *MockAPIMockRecorder) GetTitle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockAPI)(nil).GetTitle))
}

// GetVehicle mocks base method.
func (m *MockAPI) GetVehicle() api.Vehicle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVehicle")
	ret0, _ := ret[0].(api.Vehicle)
	return ret0
}

// GetVehicle indicates an expected call of GetVehicle.
func (mr *MockAPIMockRecorder) GetVehicle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVehicle", reflect.TypeOf((*MockAPI)(nil).GetVehicle))
}

// HasChargeMeter mocks base method.
func (m *MockAPI) HasChargeMeter() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChargeMeter")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChargeMeter indicates an expected call of HasChargeMeter.
func (mr *MockAPIMockRecorder) HasChargeMeter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChargeMeter", reflect.TypeOf((*MockAPI)(nil).HasChargeMeter))
}

// IsFastChargingActive mocks base method.
func (m *MockAPI) IsFastChargingActive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFastChargingActive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFastChargingActive indicates an expected call of IsFastChargingActive.
func (mr *MockAPIMockRecorder) IsFastChargingActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFastChargingActive", reflect.TypeOf((*MockAPI)(nil).IsFastChargingActive))
}

// PublishEffectiveValues mocks base method.
func (m *MockAPI) PublishEffectiveValues() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishEffectiveValues")
}

// PublishEffectiveValues indicates an expected call of PublishEffectiveValues.
func (mr *MockAPIMockRecorder) PublishEffectiveValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishEffectiveValues", reflect.TypeOf((*MockAPI)(nil).PublishEffectiveValues))
}

// RemoteControl mocks base method.
func (m *MockAPI) RemoteControl(arg0 string, arg1 RemoteDemand) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoteControl", arg0, arg1)
}

// RemoteControl indicates an expected call of RemoteControl.
func (mr *MockAPIMockRecorder) RemoteControl(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteControl", reflect.TypeOf((*MockAPI)(nil).RemoteControl), arg0, arg1)
}

// SetDisableThreshold mocks base method.
func (m *MockAPI) SetDisableThreshold(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDisableThreshold", arg0)
}

// SetDisableThreshold indicates an expected call of SetDisableThreshold.
func (mr *MockAPIMockRecorder) SetDisableThreshold(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDisableThreshold", reflect.TypeOf((*MockAPI)(nil).SetDisableThreshold), arg0)
}

// SetEnableThreshold mocks base method.
func (m *MockAPI) SetEnableThreshold(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEnableThreshold", arg0)
}

// SetEnableThreshold indicates an expected call of SetEnableThreshold.
func (mr *MockAPIMockRecorder) SetEnableThreshold(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnableThreshold", reflect.TypeOf((*MockAPI)(nil).SetEnableThreshold), arg0)
}

// SetLimitEnergy mocks base method.
func (m *MockAPI) SetLimitEnergy(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLimitEnergy", arg0)
}

// SetLimitEnergy indicates an expected call of SetLimitEnergy.
func (mr *MockAPIMockRecorder) SetLimitEnergy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLimitEnergy", reflect.TypeOf((*MockAPI)(nil).SetLimitEnergy), arg0)
}

// SetLimitSoc mocks base method.
func (m *MockAPI) SetLimitSoc(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLimitSoc", arg0)
}

// SetLimitSoc indicates an expected call of SetLimitSoc.
func (mr *MockAPIMockRecorder) SetLimitSoc(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLimitSoc", reflect.TypeOf((*MockAPI)(nil).SetLimitSoc), arg0)
}

// SetMaxCurrent mocks base method.
func (m *MockAPI) SetMaxCurrent(arg0 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMaxCurrent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMaxCurrent indicates an expected call of SetMaxCurrent.
func (mr *MockAPIMockRecorder) SetMaxCurrent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxCurrent", reflect.TypeOf((*MockAPI)(nil).SetMaxCurrent), arg0)
}

// SetMinCurrent mocks base method.
func (m *MockAPI) SetMinCurrent(arg0 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMinCurrent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMinCurrent indicates an expected call of SetMinCurrent.
func (mr *MockAPIMockRecorder) SetMinCurrent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMinCurrent", reflect.TypeOf((*MockAPI)(nil).SetMinCurrent), arg0)
}

// SetMode mocks base method.
func (m *MockAPI) SetMode(arg0 api.ChargeMode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMode", arg0)
}

// SetMode indicates an expected call of SetMode.
func (mr *MockAPIMockRecorder) SetMode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMode", reflect.TypeOf((*MockAPI)(nil).SetMode), arg0)
}

// SetPhases mocks base method.
func (m *MockAPI) SetPhases(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPhases", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPhases indicates an expected call of SetPhases.
func (mr *MockAPIMockRecorder) SetPhases(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPhases", reflect.TypeOf((*MockAPI)(nil).SetPhases), arg0)
}

// SetPlanEnergy mocks base method.
func (m *MockAPI) SetPlanEnergy(arg0 time.Time, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlanEnergy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlanEnergy indicates an expected call of SetPlanEnergy.
func (mr *MockAPIMockRecorder) SetPlanEnergy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlanEnergy", reflect.TypeOf((*MockAPI)(nil).SetPlanEnergy), arg0, arg1)
}

// SetPriority mocks base method.
func (m *MockAPI) SetPriority(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPriority", arg0)
}

// SetPriority indicates an expected call of SetPriority.
func (mr *MockAPIMockRecorder) SetPriority(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPriority", reflect.TypeOf((*MockAPI)(nil).SetPriority), arg0)
}

// SetSmartCostLimit mocks base method.
func (m *MockAPI) SetSmartCostLimit(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSmartCostLimit", arg0)
}

// SetSmartCostLimit indicates an expected call of SetSmartCostLimit.
func (mr *MockAPIMockRecorder) SetSmartCostLimit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSmartCostLimit", reflect.TypeOf((*MockAPI)(nil).SetSmartCostLimit), arg0)
}

// SetSocConfig mocks base method.
func (m *MockAPI) SetSocConfig(arg0 SocConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSocConfig", arg0)
}

// SetSocConfig indicates an expected call of SetSocConfig.
func (mr *MockAPIMockRecorder) SetSocConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSocConfig", reflect.TypeOf((*MockAPI)(nil).SetSocConfig), arg0)
}

// SetThresholds mocks base method.
func (m *MockAPI) SetThresholds(arg0 ThresholdsConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetThresholds", arg0)
}

// SetThresholds indicates an expected call of SetThresholds.
func (mr *MockAPIMockRecorder) SetThresholds(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetThresholds", reflect.TypeOf((*MockAPI)(nil).SetThresholds), arg0)
}

// SetTitle mocks base method.
func (m *MockAPI) SetTitle(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTitle", arg0)
}

// SetTitle indicates an expected call of SetTitle.
func (mr *MockAPIMockRecorder) SetTitle(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockAPI)(nil).SetTitle), arg0)
}

// SetVehicle mocks base method.
func (m *MockAPI) SetVehicle(arg0 api.Vehicle) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVehicle", arg0)
}

// SetVehicle indicates an expected call of SetVehicle.
func (mr *MockAPIMockRecorder) SetVehicle(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVehicle", reflect.TypeOf((*MockAPI)(nil).SetVehicle), arg0)
}

// SocBasedPlanning mocks base method.
func (m *MockAPI) SocBasedPlanning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SocBasedPlanning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SocBasedPlanning indicates an expected call of SocBasedPlanning.
func (mr *MockAPIMockRecorder) SocBasedPlanning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SocBasedPlanning", reflect.TypeOf((*MockAPI)(nil).SocBasedPlanning))
}

// StartVehicleDetection mocks base method.
func (m *MockAPI) StartVehicleDetection() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartVehicleDetection")
}

// StartVehicleDetection indicates an expected call of StartVehicleDetection.
func (mr *MockAPIMockRecorder) StartVehicleDetection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartVehicleDetection", reflect.TypeOf((*MockAPI)(nil).StartVehicleDetection))
}
