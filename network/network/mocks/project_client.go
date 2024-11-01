// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	"context"

	"github.com/ignite/network/x/project/types"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// ProjectClient is an autogenerated mock type for the ProjectClient type
type ProjectClient struct {
	mock.Mock
}

type ProjectClient_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectClient) EXPECT() *ProjectClient_Expecter {
	return &ProjectClient_Expecter{mock: &_m.Mock}
}

// GetMainnetAccount provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) GetMainnetAccount(ctx context.Context, in *types.QueryGetMainnetAccountRequest, opts ...grpc.CallOption) (*types.QueryGetMainnetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMainnetAccount")
	}

	var r0 *types.QueryGetMainnetAccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) (*types.QueryGetMainnetAccountResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) *types.QueryGetMainnetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMainnetAccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_GetMainnetAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMainnetAccount'
type ProjectClient_GetMainnetAccount_Call struct {
	*mock.Call
}

// GetMainnetAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetMainnetAccountRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) GetMainnetAccount(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_GetMainnetAccount_Call {
	return &ProjectClient_GetMainnetAccount_Call{Call: _e.mock.On("GetMainnetAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_GetMainnetAccount_Call) Run(run func(ctx context.Context, in *types.QueryGetMainnetAccountRequest, opts ...grpc.CallOption)) *ProjectClient_GetMainnetAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetMainnetAccountRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_GetMainnetAccount_Call) Return(_a0 *types.QueryGetMainnetAccountResponse, _a1 error) *ProjectClient_GetMainnetAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_GetMainnetAccount_Call) RunAndReturn(run func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) (*types.QueryGetMainnetAccountResponse, error)) *ProjectClient_GetMainnetAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetProject provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) GetProject(ctx context.Context, in *types.QueryGetProjectRequest, opts ...grpc.CallOption) (*types.QueryGetProjectResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProject")
	}

	var r0 *types.QueryGetProjectResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetProjectRequest, ...grpc.CallOption) (*types.QueryGetProjectResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetProjectRequest, ...grpc.CallOption) *types.QueryGetProjectResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetProjectResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetProjectRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_GetProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProject'
type ProjectClient_GetProject_Call struct {
	*mock.Call
}

// GetProject is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetProjectRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) GetProject(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_GetProject_Call {
	return &ProjectClient_GetProject_Call{Call: _e.mock.On("GetProject",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_GetProject_Call) Run(run func(ctx context.Context, in *types.QueryGetProjectRequest, opts ...grpc.CallOption)) *ProjectClient_GetProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetProjectRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_GetProject_Call) Return(_a0 *types.QueryGetProjectResponse, _a1 error) *ProjectClient_GetProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_GetProject_Call) RunAndReturn(run func(context.Context, *types.QueryGetProjectRequest, ...grpc.CallOption) (*types.QueryGetProjectResponse, error)) *ProjectClient_GetProject_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectChains provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) GetProjectChains(ctx context.Context, in *types.QueryGetProjectChainsRequest, opts ...grpc.CallOption) (*types.QueryGetProjectChainsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectChains")
	}

	var r0 *types.QueryGetProjectChainsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetProjectChainsRequest, ...grpc.CallOption) (*types.QueryGetProjectChainsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetProjectChainsRequest, ...grpc.CallOption) *types.QueryGetProjectChainsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetProjectChainsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetProjectChainsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_GetProjectChains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectChains'
type ProjectClient_GetProjectChains_Call struct {
	*mock.Call
}

// GetProjectChains is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetProjectChainsRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) GetProjectChains(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_GetProjectChains_Call {
	return &ProjectClient_GetProjectChains_Call{Call: _e.mock.On("GetProjectChains",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_GetProjectChains_Call) Run(run func(ctx context.Context, in *types.QueryGetProjectChainsRequest, opts ...grpc.CallOption)) *ProjectClient_GetProjectChains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetProjectChainsRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_GetProjectChains_Call) Return(_a0 *types.QueryGetProjectChainsResponse, _a1 error) *ProjectClient_GetProjectChains_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_GetProjectChains_Call) RunAndReturn(run func(context.Context, *types.QueryGetProjectChainsRequest, ...grpc.CallOption) (*types.QueryGetProjectChainsResponse, error)) *ProjectClient_GetProjectChains_Call {
	_c.Call.Return(run)
	return _c
}

// ListMainnetAccount provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) ListMainnetAccount(ctx context.Context, in *types.QueryAllMainnetAccountRequest, opts ...grpc.CallOption) (*types.QueryAllMainnetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMainnetAccount")
	}

	var r0 *types.QueryAllMainnetAccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) (*types.QueryAllMainnetAccountResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) *types.QueryAllMainnetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMainnetAccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_ListMainnetAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMainnetAccount'
type ProjectClient_ListMainnetAccount_Call struct {
	*mock.Call
}

// ListMainnetAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllMainnetAccountRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) ListMainnetAccount(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_ListMainnetAccount_Call {
	return &ProjectClient_ListMainnetAccount_Call{Call: _e.mock.On("ListMainnetAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_ListMainnetAccount_Call) Run(run func(ctx context.Context, in *types.QueryAllMainnetAccountRequest, opts ...grpc.CallOption)) *ProjectClient_ListMainnetAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllMainnetAccountRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_ListMainnetAccount_Call) Return(_a0 *types.QueryAllMainnetAccountResponse, _a1 error) *ProjectClient_ListMainnetAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_ListMainnetAccount_Call) RunAndReturn(run func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) (*types.QueryAllMainnetAccountResponse, error)) *ProjectClient_ListMainnetAccount_Call {
	_c.Call.Return(run)
	return _c
}

// ListMainnetAccountBalance provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) ListMainnetAccountBalance(ctx context.Context, in *types.QueryListMainnetAccountBalanceRequest, opts ...grpc.CallOption) (*types.QueryListMainnetAccountBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMainnetAccountBalance")
	}

	var r0 *types.QueryListMainnetAccountBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryListMainnetAccountBalanceRequest, ...grpc.CallOption) (*types.QueryListMainnetAccountBalanceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryListMainnetAccountBalanceRequest, ...grpc.CallOption) *types.QueryListMainnetAccountBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryListMainnetAccountBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryListMainnetAccountBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_ListMainnetAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMainnetAccountBalance'
type ProjectClient_ListMainnetAccountBalance_Call struct {
	*mock.Call
}

// ListMainnetAccountBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryListMainnetAccountBalanceRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) ListMainnetAccountBalance(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_ListMainnetAccountBalance_Call {
	return &ProjectClient_ListMainnetAccountBalance_Call{Call: _e.mock.On("ListMainnetAccountBalance",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_ListMainnetAccountBalance_Call) Run(run func(ctx context.Context, in *types.QueryListMainnetAccountBalanceRequest, opts ...grpc.CallOption)) *ProjectClient_ListMainnetAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryListMainnetAccountBalanceRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_ListMainnetAccountBalance_Call) Return(_a0 *types.QueryListMainnetAccountBalanceResponse, _a1 error) *ProjectClient_ListMainnetAccountBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_ListMainnetAccountBalance_Call) RunAndReturn(run func(context.Context, *types.QueryListMainnetAccountBalanceRequest, ...grpc.CallOption) (*types.QueryListMainnetAccountBalanceResponse, error)) *ProjectClient_ListMainnetAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// ListProject provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) ListProject(ctx context.Context, in *types.QueryAllProjectRequest, opts ...grpc.CallOption) (*types.QueryAllProjectResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProject")
	}

	var r0 *types.QueryAllProjectResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllProjectRequest, ...grpc.CallOption) (*types.QueryAllProjectResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllProjectRequest, ...grpc.CallOption) *types.QueryAllProjectResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllProjectResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllProjectRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_ListProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProject'
type ProjectClient_ListProject_Call struct {
	*mock.Call
}

// ListProject is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllProjectRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) ListProject(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_ListProject_Call {
	return &ProjectClient_ListProject_Call{Call: _e.mock.On("ListProject",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_ListProject_Call) Run(run func(ctx context.Context, in *types.QueryAllProjectRequest, opts ...grpc.CallOption)) *ProjectClient_ListProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllProjectRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_ListProject_Call) Return(_a0 *types.QueryAllProjectResponse, _a1 error) *ProjectClient_ListProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_ListProject_Call) RunAndReturn(run func(context.Context, *types.QueryAllProjectRequest, ...grpc.CallOption) (*types.QueryAllProjectResponse, error)) *ProjectClient_ListProject_Call {
	_c.Call.Return(run)
	return _c
}

// MainnetAccountBalance provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) MainnetAccountBalance(ctx context.Context, in *types.QueryMainnetAccountBalanceRequest, opts ...grpc.CallOption) (*types.QueryMainnetAccountBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MainnetAccountBalance")
	}

	var r0 *types.QueryMainnetAccountBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMainnetAccountBalanceRequest, ...grpc.CallOption) (*types.QueryMainnetAccountBalanceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMainnetAccountBalanceRequest, ...grpc.CallOption) *types.QueryMainnetAccountBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMainnetAccountBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMainnetAccountBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_MainnetAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MainnetAccountBalance'
type ProjectClient_MainnetAccountBalance_Call struct {
	*mock.Call
}

// MainnetAccountBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryMainnetAccountBalanceRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) MainnetAccountBalance(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_MainnetAccountBalance_Call {
	return &ProjectClient_MainnetAccountBalance_Call{Call: _e.mock.On("MainnetAccountBalance",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_MainnetAccountBalance_Call) Run(run func(ctx context.Context, in *types.QueryMainnetAccountBalanceRequest, opts ...grpc.CallOption)) *ProjectClient_MainnetAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryMainnetAccountBalanceRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_MainnetAccountBalance_Call) Return(_a0 *types.QueryMainnetAccountBalanceResponse, _a1 error) *ProjectClient_MainnetAccountBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_MainnetAccountBalance_Call) RunAndReturn(run func(context.Context, *types.QueryMainnetAccountBalanceRequest, ...grpc.CallOption) (*types.QueryMainnetAccountBalanceResponse, error)) *ProjectClient_MainnetAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Params")
	}

	var r0 *types.QueryParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) (*types.QueryParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_Params_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Params'
type ProjectClient_Params_Call struct {
	*mock.Call
}

// Params is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryParamsRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) Params(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_Params_Call {
	return &ProjectClient_Params_Call{Call: _e.mock.On("Params",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_Params_Call) Run(run func(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption)) *ProjectClient_Params_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryParamsRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_Params_Call) Return(_a0 *types.QueryParamsResponse, _a1 error) *ProjectClient_Params_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_Params_Call) RunAndReturn(run func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) (*types.QueryParamsResponse, error)) *ProjectClient_Params_Call {
	_c.Call.Return(run)
	return _c
}

// SpecialAllocationsBalance provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) SpecialAllocationsBalance(ctx context.Context, in *types.QuerySpecialAllocationsBalanceRequest, opts ...grpc.CallOption) (*types.QuerySpecialAllocationsBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SpecialAllocationsBalance")
	}

	var r0 *types.QuerySpecialAllocationsBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpecialAllocationsBalanceRequest, ...grpc.CallOption) (*types.QuerySpecialAllocationsBalanceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpecialAllocationsBalanceRequest, ...grpc.CallOption) *types.QuerySpecialAllocationsBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySpecialAllocationsBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySpecialAllocationsBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_SpecialAllocationsBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpecialAllocationsBalance'
type ProjectClient_SpecialAllocationsBalance_Call struct {
	*mock.Call
}

// SpecialAllocationsBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QuerySpecialAllocationsBalanceRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) SpecialAllocationsBalance(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_SpecialAllocationsBalance_Call {
	return &ProjectClient_SpecialAllocationsBalance_Call{Call: _e.mock.On("SpecialAllocationsBalance",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_SpecialAllocationsBalance_Call) Run(run func(ctx context.Context, in *types.QuerySpecialAllocationsBalanceRequest, opts ...grpc.CallOption)) *ProjectClient_SpecialAllocationsBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QuerySpecialAllocationsBalanceRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_SpecialAllocationsBalance_Call) Return(_a0 *types.QuerySpecialAllocationsBalanceResponse, _a1 error) *ProjectClient_SpecialAllocationsBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_SpecialAllocationsBalance_Call) RunAndReturn(run func(context.Context, *types.QuerySpecialAllocationsBalanceRequest, ...grpc.CallOption) (*types.QuerySpecialAllocationsBalanceResponse, error)) *ProjectClient_SpecialAllocationsBalance_Call {
	_c.Call.Return(run)
	return _c
}

// TotalShares provides a mock function with given fields: ctx, in, opts
func (_m *ProjectClient) TotalShares(ctx context.Context, in *types.QueryTotalSharesRequest, opts ...grpc.CallOption) (*types.QueryTotalSharesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TotalShares")
	}

	var r0 *types.QueryTotalSharesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) (*types.QueryTotalSharesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) *types.QueryTotalSharesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryTotalSharesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectClient_TotalShares_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TotalShares'
type ProjectClient_TotalShares_Call struct {
	*mock.Call
}

// TotalShares is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryTotalSharesRequest
//   - opts ...grpc.CallOption
func (_e *ProjectClient_Expecter) TotalShares(ctx interface{}, in interface{}, opts ...interface{}) *ProjectClient_TotalShares_Call {
	return &ProjectClient_TotalShares_Call{Call: _e.mock.On("TotalShares",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ProjectClient_TotalShares_Call) Run(run func(ctx context.Context, in *types.QueryTotalSharesRequest, opts ...grpc.CallOption)) *ProjectClient_TotalShares_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryTotalSharesRequest), variadicArgs...)
	})
	return _c
}

func (_c *ProjectClient_TotalShares_Call) Return(_a0 *types.QueryTotalSharesResponse, _a1 error) *ProjectClient_TotalShares_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectClient_TotalShares_Call) RunAndReturn(run func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) (*types.QueryTotalSharesResponse, error)) *ProjectClient_TotalShares_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectClient creates a new instance of ProjectClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectClient {
	mock := &ProjectClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}