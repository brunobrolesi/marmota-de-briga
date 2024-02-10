// Code generated by mockery v2.40.3. DO NOT EDIT.

package mock_gateway

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

// MockTransactionRepository is an autogenerated mock type for the TransactionRepository type
type MockTransactionRepository struct {
	mock.Mock
}

type MockTransactionRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransactionRepository) EXPECT() *MockTransactionRepository_Expecter {
	return &MockTransactionRepository_Expecter{mock: &_m.Mock}
}

// CreateTransaction provides a mock function with given fields: ctx, clientID, value, transactionType, description
func (_m *MockTransactionRepository) CreateTransaction(ctx context.Context, clientID int, value int, transactionType model.TransactionType, description string) (*model.Transaction, error) {
	ret := _m.Called(ctx, clientID, value, transactionType, description)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, model.TransactionType, string) (*model.Transaction, error)); ok {
		return rf(ctx, clientID, value, transactionType, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, model.TransactionType, string) *model.Transaction); ok {
		r0 = rf(ctx, clientID, value, transactionType, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, model.TransactionType, string) error); ok {
		r1 = rf(ctx, clientID, value, transactionType, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactionRepository_CreateTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransaction'
type MockTransactionRepository_CreateTransaction_Call struct {
	*mock.Call
}

// CreateTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - clientID int
//   - value int
//   - transactionType model.TransactionType
//   - description string
func (_e *MockTransactionRepository_Expecter) CreateTransaction(ctx interface{}, clientID interface{}, value interface{}, transactionType interface{}, description interface{}) *MockTransactionRepository_CreateTransaction_Call {
	return &MockTransactionRepository_CreateTransaction_Call{Call: _e.mock.On("CreateTransaction", ctx, clientID, value, transactionType, description)}
}

func (_c *MockTransactionRepository_CreateTransaction_Call) Run(run func(ctx context.Context, clientID int, value int, transactionType model.TransactionType, description string)) *MockTransactionRepository_CreateTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(model.TransactionType), args[4].(string))
	})
	return _c
}

func (_c *MockTransactionRepository_CreateTransaction_Call) Return(_a0 *model.Transaction, _a1 error) *MockTransactionRepository_CreateTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactionRepository_CreateTransaction_Call) RunAndReturn(run func(context.Context, int, int, model.TransactionType, string) (*model.Transaction, error)) *MockTransactionRepository_CreateTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransactionRepository creates a new instance of MockTransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactionRepository {
	mock := &MockTransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
