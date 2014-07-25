// Automatically generated by MockGen. DO NOT EDIT!
// Source: io/docker.go

package io

import (
	time "time"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of DockerCli interface
type MockDockerCli struct {
	ctrl     *gomock.Controller
	recorder *_MockDockerCliRecorder
}

// Recorder for MockDockerCli (not exported)
type _MockDockerCliRecorder struct {
	mock *MockDockerCli
}

func NewMockDockerCli(ctrl *gomock.Controller) *MockDockerCli {
	mock := &MockDockerCli{ctrl: ctrl}
	mock.recorder = &_MockDockerCliRecorder{mock}
	return mock
}

func (_m *MockDockerCli) EXPECT() *_MockDockerCliRecorder {
	return _m.recorder
}

func (_m *MockDockerCli) CmdRun(_param0 bool, _param1 ...string) error {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdRun", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdRun(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdRun", _s...)
}

func (_m *MockDockerCli) CmdPs(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdPs", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdPs(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdPs", arg0...)
}

func (_m *MockDockerCli) CmdTag(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdTag", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdTag(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdTag", arg0...)
}

func (_m *MockDockerCli) CmdCommit(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdCommit", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdCommit(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdCommit", arg0...)
}

func (_m *MockDockerCli) CmdInspect(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdInspect", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdInspect(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdInspect", arg0...)
}

func (_m *MockDockerCli) CmdBuild(_param0 bool, _param1 ...string) error {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdBuild", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdBuild(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdBuild", _s...)
}

func (_m *MockDockerCli) CmdCp(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdCp", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdCp(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdCp", arg0...)
}

func (_m *MockDockerCli) CmdWait(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdWait", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdWait(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdWait", arg0...)
}

func (_m *MockDockerCli) CmdAttach(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdAttach", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdAttach(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdAttach", arg0...)
}

func (_m *MockDockerCli) CmdStop(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdStop", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdStop(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdStop", arg0...)
}

func (_m *MockDockerCli) Stdout() string {
	ret := _m.ctrl.Call(_m, "Stdout")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDockerCliRecorder) Stdout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stdout")
}

func (_m *MockDockerCli) LastLineOfStdout() string {
	ret := _m.ctrl.Call(_m, "LastLineOfStdout")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDockerCliRecorder) LastLineOfStdout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastLineOfStdout")
}

func (_m *MockDockerCli) Stderr() string {
	ret := _m.ctrl.Call(_m, "Stderr")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDockerCliRecorder) Stderr() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stderr")
}

func (_m *MockDockerCli) EmptyOutput() bool {
	ret := _m.ctrl.Call(_m, "EmptyOutput")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDockerCliRecorder) EmptyOutput() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EmptyOutput")
}

func (_m *MockDockerCli) DecodeInspect(_param0 ...string) (Inspected, error) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DecodeInspect", _s...)
	ret0, _ := ret[0].(Inspected)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerCliRecorder) DecodeInspect(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeInspect", arg0...)
}

func (_m *MockDockerCli) DumpErrOutput() {
	_m.ctrl.Call(_m, "DumpErrOutput")
}

func (_mr *_MockDockerCliRecorder) DumpErrOutput() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DumpErrOutput")
}

// Mock of Inspected interface
type MockInspected struct {
	ctrl     *gomock.Controller
	recorder *_MockInspectedRecorder
}

// Recorder for MockInspected (not exported)
type _MockInspectedRecorder struct {
	mock *MockInspected
}

func NewMockInspected(ctrl *gomock.Controller) *MockInspected {
	mock := &MockInspected{ctrl: ctrl}
	mock.recorder = &_MockInspectedRecorder{mock}
	return mock
}

func (_m *MockInspected) EXPECT() *_MockInspectedRecorder {
	return _m.recorder
}

func (_m *MockInspected) CreatedTime() time.Time {
	ret := _m.ctrl.Call(_m, "CreatedTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockInspectedRecorder) CreatedTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreatedTime")
}

func (_m *MockInspected) ContainerName() string {
	ret := _m.ctrl.Call(_m, "ContainerName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockInspectedRecorder) ContainerName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerName")
}

func (_m *MockInspected) Running() bool {
	ret := _m.ctrl.Call(_m, "Running")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockInspectedRecorder) Running() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Running")
}
