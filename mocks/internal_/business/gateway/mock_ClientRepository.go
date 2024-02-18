// Code generated by mockery v2.40.3. DO NOT EDIT.

package mock_gateway

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

// MockClientRepository is an autogenerated mock type for the ClientRepository type
type MockClientRepository struct {
	mock.Mock
}

type MockClientRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientRepository) EXPECT() *MockClientRepository_Expecter {
	return &MockClientRepository_Expecter{mock: &_m.Mock}
}

// ACIDUpdateBalance provides a mock function with given fields: ctx, clientID, transactionValue, transactionType
func (_m *MockClientRepository) ACIDUpdateBalance(ctx context.Context, clientID int, transactionValue int, transactionType model.TransactionType) (*model.Client, error) {
	ret := _m.Called(ctx, clientID, transactionValue, transactionType)

	if len(ret) == 0 {
		panic("no return value specified for ACIDUpdateBalance")
	}

	var r0 *model.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, model.TransactionType) (*model.Client, error)); ok {
		return rf(ctx, clientID, transactionValue, transactionType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, model.TransactionType) *model.Client); ok {
		r0 = rf(ctx, clientID, transactionValue, transactionType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, model.TransactionType) error); ok {
		r1 = rf(ctx, clientID, transactionValue, transactionType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientRepository_ACIDUpdateBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ACIDUpdateBalance'
type MockClientRepository_ACIDUpdateBalance_Call struct {
	*mock.Call
}

// ACIDUpdateBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - clientID int
//   - transactionValue int
//   - transactionType model.TransactionType
func (_e *MockClientRepository_Expecter) ACIDUpdateBalance(ctx interface{}, clientID interface{}, transactionValue interface{}, transactionType interface{}) *MockClientRepository_ACIDUpdateBalance_Call {
	return &MockClientRepository_ACIDUpdateBalance_Call{Call: _e.mock.On("ACIDUpdateBalance", ctx, clientID, transactionValue, transactionType)}
}

func (_c *MockClientRepository_ACIDUpdateBalance_Call) Run(run func(ctx context.Context, clientID int, transactionValue int, transactionType model.TransactionType)) *MockClientRepository_ACIDUpdateBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(model.TransactionType))
	})
	return _c
}

func (_c *MockClientRepository_ACIDUpdateBalance_Call) Return(_a0 *model.Client, _a1 error) *MockClientRepository_ACIDUpdateBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientRepository_ACIDUpdateBalance_Call) RunAndReturn(run func(context.Context, int, int, model.TransactionType) (*model.Client, error)) *MockClientRepository_ACIDUpdateBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetClient provides a mock function with given fields: ctx, clientID
func (_m *MockClientRepository) GetClient(ctx context.Context, clientID int) (*model.Client, error) {
	ret := _m.Called(ctx, clientID)

	if len(ret) == 0 {
		panic("no return value specified for GetClient")
	}

	var r0 *model.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*model.Client, error)); ok {
		return rf(ctx, clientID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.Client); ok {
		r0 = rf(ctx, clientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientRepository_GetClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClient'
type MockClientRepository_GetClient_Call struct {
	*mock.Call
}

// GetClient is a helper method to define mock.On call
//   - ctx context.Context
//   - clientID int
func (_e *MockClientRepository_Expecter) GetClient(ctx interface{}, clientID interface{}) *MockClientRepository_GetClient_Call {
	return &MockClientRepository_GetClient_Call{Call: _e.mock.On("GetClient", ctx, clientID)}
}

func (_c *MockClientRepository_GetClient_Call) Run(run func(ctx context.Context, clientID int)) *MockClientRepository_GetClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockClientRepository_GetClient_Call) Return(_a0 *model.Client, _a1 error) *MockClientRepository_GetClient_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientRepository_GetClient_Call) RunAndReturn(run func(context.Context, int) (*model.Client, error)) *MockClientRepository_GetClient_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientRepository creates a new instance of MockClientRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientRepository {
	mock := &MockClientRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
