// Code generated by mockery v2.43.0. DO NOT EDIT.

package queue

import (
	context "context"

	models "github.com/content-services/content-sources-backend/pkg/models"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// MockQueue is an autogenerated mock type for the Queue type
type MockQueue struct {
	mock.Mock
}

// Cancel provides a mock function with given fields: ctx, taskId
func (_m *MockQueue) Cancel(ctx context.Context, taskId uuid.UUID) error {
	ret := _m.Called(ctx, taskId)

	if len(ret) == 0 {
		panic("no return value specified for Cancel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, taskId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dequeue provides a mock function with given fields: ctx, taskTypes
func (_m *MockQueue) Dequeue(ctx context.Context, taskTypes []string) (*models.TaskInfo, error) {
	ret := _m.Called(ctx, taskTypes)

	if len(ret) == 0 {
		panic("no return value specified for Dequeue")
	}

	var r0 *models.TaskInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (*models.TaskInfo, error)); ok {
		return rf(ctx, taskTypes)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) *models.TaskInfo); ok {
		r0 = rf(ctx, taskTypes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TaskInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, taskTypes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Enqueue provides a mock function with given fields: task
func (_m *MockQueue) Enqueue(task *Task) (uuid.UUID, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for Enqueue")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*Task) (uuid.UUID, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(*Task) uuid.UUID); ok {
		r0 = rf(task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Finish provides a mock function with given fields: taskId, taskError
func (_m *MockQueue) Finish(taskId uuid.UUID, taskError error) error {
	ret := _m.Called(taskId, taskError)

	if len(ret) == 0 {
		panic("no return value specified for Finish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, error) error); ok {
		r0 = rf(taskId, taskError)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Heartbeats provides a mock function with given fields: olderThan
func (_m *MockQueue) Heartbeats(olderThan time.Duration) []uuid.UUID {
	ret := _m.Called(olderThan)

	if len(ret) == 0 {
		panic("no return value specified for Heartbeats")
	}

	var r0 []uuid.UUID
	if rf, ok := ret.Get(0).(func(time.Duration) []uuid.UUID); ok {
		r0 = rf(olderThan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uuid.UUID)
		}
	}

	return r0
}

// IdFromToken provides a mock function with given fields: token
func (_m *MockQueue) IdFromToken(token uuid.UUID) (uuid.UUID, bool, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for IdFromToken")
	}

	var r0 uuid.UUID
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (uuid.UUID, bool, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) uuid.UUID); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) bool); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(uuid.UUID) error); ok {
		r2 = rf(token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListenForCancel provides a mock function with given fields: ctx, taskID, cancelFunc
func (_m *MockQueue) ListenForCancel(ctx context.Context, taskID uuid.UUID, cancelFunc context.CancelCauseFunc) {
	_m.Called(ctx, taskID, cancelFunc)
}

// RefreshHeartbeat provides a mock function with given fields: token
func (_m *MockQueue) RefreshHeartbeat(token uuid.UUID) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for RefreshHeartbeat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Requeue provides a mock function with given fields: taskId
func (_m *MockQueue) Requeue(taskId uuid.UUID) error {
	ret := _m.Called(taskId)

	if len(ret) == 0 {
		panic("no return value specified for Requeue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(taskId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequeueFailedTasks provides a mock function with given fields: taskTypes
func (_m *MockQueue) RequeueFailedTasks(taskTypes []string) error {
	ret := _m.Called(taskTypes)

	if len(ret) == 0 {
		panic("no return value specified for RequeueFailedTasks")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(taskTypes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields: taskId
func (_m *MockQueue) Status(taskId uuid.UUID) (*models.TaskInfo, error) {
	ret := _m.Called(taskId)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 *models.TaskInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.TaskInfo, error)); ok {
		return rf(taskId)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.TaskInfo); ok {
		r0 = rf(taskId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TaskInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(taskId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayload provides a mock function with given fields: task, payload
func (_m *MockQueue) UpdatePayload(task *models.TaskInfo, payload interface{}) (*models.TaskInfo, error) {
	ret := _m.Called(task, payload)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePayload")
	}

	var r0 *models.TaskInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.TaskInfo, interface{}) (*models.TaskInfo, error)); ok {
		return rf(task, payload)
	}
	if rf, ok := ret.Get(0).(func(*models.TaskInfo, interface{}) *models.TaskInfo); ok {
		r0 = rf(task, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TaskInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.TaskInfo, interface{}) error); ok {
		r1 = rf(task, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockQueue creates a new instance of MockQueue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueue(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueue {
	mock := &MockQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
