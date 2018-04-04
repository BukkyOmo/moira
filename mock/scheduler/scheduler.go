// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira/notifier (interfaces: Scheduler)

// Package mock_scheduler is a generated GoMock package.
package mock_scheduler

import (
	gomock "github.com/golang/mock/gomock"
	moira "github.com/moira-alert/moira"
	reflect "reflect"
	time "time"
)

// MockScheduler is a mock of Scheduler interface
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// ScheduleNotification mocks base method
func (m *MockScheduler) ScheduleNotification(arg0 time.Time, arg1 moira.NotificationEvent, arg2 moira.TriggerData, arg3 moira.ContactData, arg4 bool, arg5 int) *moira.ScheduledNotification {
	ret := m.ctrl.Call(m, "ScheduleNotification", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*moira.ScheduledNotification)
	return ret0
}

// ScheduleNotification indicates an expected call of ScheduleNotification
func (mr *MockSchedulerMockRecorder) ScheduleNotification(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleNotification", reflect.TypeOf((*MockScheduler)(nil).ScheduleNotification), arg0, arg1, arg2, arg3, arg4, arg5)
}
