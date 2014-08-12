// Automatically generated by MockGen. DO NOT EDIT!
// Source: io/helper.go

package io

import (
	io "io"
	os "os"
	time "time"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of Helper interface
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *_MockHelperRecorder
}

// Recorder for MockHelper (not exported)
type _MockHelperRecorder struct {
	mock *MockHelper
}

func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &_MockHelperRecorder{mock}
	return mock
}

func (_m *MockHelper) EXPECT() *_MockHelperRecorder {
	return _m.recorder
}

func (_m *MockHelper) Debug(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debug", _s...)
}

func (_mr *_MockHelperRecorder) Debug(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debug", _s...)
}

func (_m *MockHelper) OpenDockerfileRelative(dir string) (io.Reader, error) {
	ret := _m.ctrl.Call(_m, "OpenDockerfileRelative", dir)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHelperRecorder) OpenDockerfileRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenDockerfileRelative", arg0)
}

func (_m *MockHelper) OpenFileRelative(path string) (*os.File, error) {
	ret := _m.ctrl.Call(_m, "OpenFileRelative", path)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHelperRecorder) OpenFileRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenFileRelative", arg0)
}

func (_m *MockHelper) DirectoryRelative(dir string) string {
	ret := _m.ctrl.Call(_m, "DirectoryRelative", dir)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockHelperRecorder) DirectoryRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DirectoryRelative", arg0)
}

func (_m *MockHelper) Fatalf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatalf", _s...)
}

func (_mr *_MockHelperRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatalf", _s...)
}

func (_m *MockHelper) CheckFatal(_param0 error, _param1 string, _param2 ...interface{}) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "CheckFatal", _s...)
}

func (_mr *_MockHelperRecorder) CheckFatal(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckFatal", _s...)
}

func (_m *MockHelper) ConfigReader() io.Reader {
	ret := _m.ctrl.Call(_m, "ConfigReader")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

func (_mr *_MockHelperRecorder) ConfigReader() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigReader")
}

func (_m *MockHelper) ConfigFile() string {
	ret := _m.ctrl.Call(_m, "ConfigFile")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockHelperRecorder) ConfigFile() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigFile")
}

func (_m *MockHelper) LastTimeInDirRelative(_param0 string) (time.Time, error) {
	ret := _m.ctrl.Call(_m, "LastTimeInDirRelative", _param0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHelperRecorder) LastTimeInDirRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastTimeInDirRelative", arg0)
}

func (_m *MockHelper) LastTimeInDir(_param0 string) (time.Time, error) {
	ret := _m.ctrl.Call(_m, "LastTimeInDir", _param0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHelperRecorder) LastTimeInDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastTimeInDir", arg0)
}
