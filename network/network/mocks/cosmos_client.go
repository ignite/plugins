// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/cosmos/cosmos-sdk/client"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"

	cosmosaccount "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"

	cosmosclient "github.com/ignite/cli/v28/ignite/pkg/cosmosclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CosmosClient is an autogenerated mock type for the CosmosClient type
type CosmosClient struct {
	mock.Mock
}

type CosmosClient_Expecter struct {
	mock *mock.Mock
}

func (_m *CosmosClient) EXPECT() *CosmosClient_Expecter {
	return &CosmosClient_Expecter{mock: &_m.Mock}
}

// BroadcastTx provides a mock function with given fields: ctx, account, msgs
func (_m *CosmosClient) BroadcastTx(ctx context.Context, account cosmosaccount.Account, msgs ...types.Msg) (cosmosclient.Response, error) {
	_va := make([]interface{}, len(msgs))
	for _i := range msgs {
		_va[_i] = msgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, account)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BroadcastTx")
	}

	var r0 cosmosclient.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmosaccount.Account, ...types.Msg) (cosmosclient.Response, error)); ok {
		return rf(ctx, account, msgs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmosaccount.Account, ...types.Msg) cosmosclient.Response); ok {
		r0 = rf(ctx, account, msgs...)
	} else {
		r0 = ret.Get(0).(cosmosclient.Response)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmosaccount.Account, ...types.Msg) error); ok {
		r1 = rf(ctx, account, msgs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CosmosClient_BroadcastTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BroadcastTx'
type CosmosClient_BroadcastTx_Call struct {
	*mock.Call
}

// BroadcastTx is a helper method to define mock.On call
//   - ctx context.Context
//   - account cosmosaccount.Account
//   - msgs ...types.Msg
func (_e *CosmosClient_Expecter) BroadcastTx(ctx interface{}, account interface{}, msgs ...interface{}) *CosmosClient_BroadcastTx_Call {
	return &CosmosClient_BroadcastTx_Call{Call: _e.mock.On("BroadcastTx",
		append([]interface{}{ctx, account}, msgs...)...)}
}

func (_c *CosmosClient_BroadcastTx_Call) Run(run func(ctx context.Context, account cosmosaccount.Account, msgs ...types.Msg)) *CosmosClient_BroadcastTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]types.Msg, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(types.Msg)
			}
		}
		run(args[0].(context.Context), args[1].(cosmosaccount.Account), variadicArgs...)
	})
	return _c
}

func (_c *CosmosClient_BroadcastTx_Call) Return(_a0 cosmosclient.Response, _a1 error) *CosmosClient_BroadcastTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CosmosClient_BroadcastTx_Call) RunAndReturn(run func(context.Context, cosmosaccount.Account, ...types.Msg) (cosmosclient.Response, error)) *CosmosClient_BroadcastTx_Call {
	_c.Call.Return(run)
	return _c
}

// ConsensusInfo provides a mock function with given fields: ctx, height
func (_m *CosmosClient) ConsensusInfo(ctx context.Context, height int64) (cosmosclient.ConsensusInfo, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("no return value specified for ConsensusInfo")
	}

	var r0 cosmosclient.ConsensusInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (cosmosclient.ConsensusInfo, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) cosmosclient.ConsensusInfo); ok {
		r0 = rf(ctx, height)
	} else {
		r0 = ret.Get(0).(cosmosclient.ConsensusInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CosmosClient_ConsensusInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsensusInfo'
type CosmosClient_ConsensusInfo_Call struct {
	*mock.Call
}

// ConsensusInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - height int64
func (_e *CosmosClient_Expecter) ConsensusInfo(ctx interface{}, height interface{}) *CosmosClient_ConsensusInfo_Call {
	return &CosmosClient_ConsensusInfo_Call{Call: _e.mock.On("ConsensusInfo", ctx, height)}
}

func (_c *CosmosClient_ConsensusInfo_Call) Run(run func(ctx context.Context, height int64)) *CosmosClient_ConsensusInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *CosmosClient_ConsensusInfo_Call) Return(_a0 cosmosclient.ConsensusInfo, _a1 error) *CosmosClient_ConsensusInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CosmosClient_ConsensusInfo_Call) RunAndReturn(run func(context.Context, int64) (cosmosclient.ConsensusInfo, error)) *CosmosClient_ConsensusInfo_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *CosmosClient) Context() client.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 client.Context
	if rf, ok := ret.Get(0).(func() client.Context); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(client.Context)
	}

	return r0
}

// CosmosClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type CosmosClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *CosmosClient_Expecter) Context() *CosmosClient_Context_Call {
	return &CosmosClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *CosmosClient_Context_Call) Run(run func()) *CosmosClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CosmosClient_Context_Call) Return(_a0 client.Context) *CosmosClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CosmosClient_Context_Call) RunAndReturn(run func() client.Context) *CosmosClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields: ctx
func (_m *CosmosClient) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 *coretypes.ResultStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*coretypes.ResultStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *coretypes.ResultStatus); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CosmosClient_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type CosmosClient_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CosmosClient_Expecter) Status(ctx interface{}) *CosmosClient_Status_Call {
	return &CosmosClient_Status_Call{Call: _e.mock.On("Status", ctx)}
}

func (_c *CosmosClient_Status_Call) Run(run func(ctx context.Context)) *CosmosClient_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CosmosClient_Status_Call) Return(_a0 *coretypes.ResultStatus, _a1 error) *CosmosClient_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CosmosClient_Status_Call) RunAndReturn(run func(context.Context) (*coretypes.ResultStatus, error)) *CosmosClient_Status_Call {
	_c.Call.Return(run)
	return _c
}

// NewCosmosClient creates a new instance of CosmosClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCosmosClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *CosmosClient {
	mock := &CosmosClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}