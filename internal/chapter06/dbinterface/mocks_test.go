// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/dbinterface (interfaces: DB,Transaction)

package dbinterface

import (
	sql "database/sql"

	gomock "github.com/golang/mock/gomock"
)

// Mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *_MockDBRecorder
}

// Recorder for MockDB (not exported)
type _MockDBRecorder struct {
	mock *MockDB
}

func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &_MockDBRecorder{mock}
	return mock
}

func (_m *MockDB) EXPECT() *_MockDBRecorder {
	return _m.recorder
}

func (_m *MockDB) Exec(_param0 string, _param1 ...interface{}) (sql.Result, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Exec", _s...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Exec(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exec", _s...)
}

func (_m *MockDB) Prepare(_param0 string) (*sql.Stmt, error) {
	ret := _m.ctrl.Call(_m, "Prepare", _param0)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Prepare(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0)
}

func (_m *MockDB) Query(_param0 string, _param1 ...interface{}) (*sql.Rows, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Query", _s...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", _s...)
}

func (_m *MockDB) QueryRow(_param0 string, _param1 ...interface{}) *sql.Row {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "QueryRow", _s...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

func (_mr *_MockDBRecorder) QueryRow(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryRow", _s...)
}

// Mock of Transaction interface
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *_MockTransactionRecorder
}

// Recorder for MockTransaction (not exported)
type _MockTransactionRecorder struct {
	mock *MockTransaction
}

func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &_MockTransactionRecorder{mock}
	return mock
}

func (_m *MockTransaction) EXPECT() *_MockTransactionRecorder {
	return _m.recorder
}

func (_m *MockTransaction) Commit() error {
	ret := _m.ctrl.Call(_m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTransactionRecorder) Commit() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Commit")
}

func (_m *MockTransaction) Exec(_param0 string, _param1 ...interface{}) (sql.Result, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Exec", _s...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTransactionRecorder) Exec(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exec", _s...)
}

func (_m *MockTransaction) Prepare(_param0 string) (*sql.Stmt, error) {
	ret := _m.ctrl.Call(_m, "Prepare", _param0)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTransactionRecorder) Prepare(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0)
}

func (_m *MockTransaction) Query(_param0 string, _param1 ...interface{}) (*sql.Rows, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Query", _s...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTransactionRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", _s...)
}

func (_m *MockTransaction) QueryRow(_param0 string, _param1 ...interface{}) *sql.Row {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "QueryRow", _s...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

func (_mr *_MockTransactionRecorder) QueryRow(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryRow", _s...)
}

func (_m *MockTransaction) Rollback() error {
	ret := _m.ctrl.Call(_m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTransactionRecorder) Rollback() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rollback")
}

func (_m *MockTransaction) Stmt(_param0 *sql.Stmt) *sql.Stmt {
	ret := _m.ctrl.Call(_m, "Stmt", _param0)
	ret0, _ := ret[0].(*sql.Stmt)
	return ret0
}

func (_mr *_MockTransactionRecorder) Stmt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stmt", arg0)
}
