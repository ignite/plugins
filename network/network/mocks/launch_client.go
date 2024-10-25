// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ignite/network/x/launch/types"
)

// LaunchClient is an autogenerated mock type for the LaunchClient type
type LaunchClient struct {
	mock.Mock
}

type LaunchClient_Expecter struct {
	mock *mock.Mock
}

func (_m *LaunchClient) EXPECT() *LaunchClient_Expecter {
	return &LaunchClient_Expecter{mock: &_m.Mock}
}

// GetChain provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) GetChain(ctx context.Context, in *types.QueryGetChainRequest, opts ...grpc.CallOption) (*types.QueryGetChainResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetChain")
	}

	var r0 *types.QueryGetChainResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetChainRequest, ...grpc.CallOption) (*types.QueryGetChainResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetChainRequest, ...grpc.CallOption) *types.QueryGetChainResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetChainResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetChainRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_GetChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChain'
type LaunchClient_GetChain_Call struct {
	*mock.Call
}

// GetChain is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetChainRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) GetChain(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_GetChain_Call {
	return &LaunchClient_GetChain_Call{Call: _e.mock.On("GetChain",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_GetChain_Call) Run(run func(ctx context.Context, in *types.QueryGetChainRequest, opts ...grpc.CallOption)) *LaunchClient_GetChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetChainRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_GetChain_Call) Return(_a0 *types.QueryGetChainResponse, _a1 error) *LaunchClient_GetChain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_GetChain_Call) RunAndReturn(run func(context.Context, *types.QueryGetChainRequest, ...grpc.CallOption) (*types.QueryGetChainResponse, error)) *LaunchClient_GetChain_Call {
	_c.Call.Return(run)
	return _c
}

// GetGenesisAccount provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) GetGenesisAccount(ctx context.Context, in *types.QueryGetGenesisAccountRequest, opts ...grpc.CallOption) (*types.QueryGetGenesisAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGenesisAccount")
	}

	var r0 *types.QueryGetGenesisAccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetGenesisAccountRequest, ...grpc.CallOption) (*types.QueryGetGenesisAccountResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetGenesisAccountRequest, ...grpc.CallOption) *types.QueryGetGenesisAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetGenesisAccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetGenesisAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_GetGenesisAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGenesisAccount'
type LaunchClient_GetGenesisAccount_Call struct {
	*mock.Call
}

// GetGenesisAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetGenesisAccountRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) GetGenesisAccount(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_GetGenesisAccount_Call {
	return &LaunchClient_GetGenesisAccount_Call{Call: _e.mock.On("GetGenesisAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_GetGenesisAccount_Call) Run(run func(ctx context.Context, in *types.QueryGetGenesisAccountRequest, opts ...grpc.CallOption)) *LaunchClient_GetGenesisAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetGenesisAccountRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_GetGenesisAccount_Call) Return(_a0 *types.QueryGetGenesisAccountResponse, _a1 error) *LaunchClient_GetGenesisAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_GetGenesisAccount_Call) RunAndReturn(run func(context.Context, *types.QueryGetGenesisAccountRequest, ...grpc.CallOption) (*types.QueryGetGenesisAccountResponse, error)) *LaunchClient_GetGenesisAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetGenesisValidator provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) GetGenesisValidator(ctx context.Context, in *types.QueryGetGenesisValidatorRequest, opts ...grpc.CallOption) (*types.QueryGetGenesisValidatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGenesisValidator")
	}

	var r0 *types.QueryGetGenesisValidatorResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetGenesisValidatorRequest, ...grpc.CallOption) (*types.QueryGetGenesisValidatorResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetGenesisValidatorRequest, ...grpc.CallOption) *types.QueryGetGenesisValidatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetGenesisValidatorResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetGenesisValidatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_GetGenesisValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGenesisValidator'
type LaunchClient_GetGenesisValidator_Call struct {
	*mock.Call
}

// GetGenesisValidator is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetGenesisValidatorRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) GetGenesisValidator(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_GetGenesisValidator_Call {
	return &LaunchClient_GetGenesisValidator_Call{Call: _e.mock.On("GetGenesisValidator",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_GetGenesisValidator_Call) Run(run func(ctx context.Context, in *types.QueryGetGenesisValidatorRequest, opts ...grpc.CallOption)) *LaunchClient_GetGenesisValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetGenesisValidatorRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_GetGenesisValidator_Call) Return(_a0 *types.QueryGetGenesisValidatorResponse, _a1 error) *LaunchClient_GetGenesisValidator_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_GetGenesisValidator_Call) RunAndReturn(run func(context.Context, *types.QueryGetGenesisValidatorRequest, ...grpc.CallOption) (*types.QueryGetGenesisValidatorResponse, error)) *LaunchClient_GetGenesisValidator_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequest provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) GetRequest(ctx context.Context, in *types.QueryGetRequestRequest, opts ...grpc.CallOption) (*types.QueryGetRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRequest")
	}

	var r0 *types.QueryGetRequestResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetRequestRequest, ...grpc.CallOption) (*types.QueryGetRequestResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetRequestRequest, ...grpc.CallOption) *types.QueryGetRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetRequestResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetRequestRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type LaunchClient_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetRequestRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) GetRequest(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_GetRequest_Call {
	return &LaunchClient_GetRequest_Call{Call: _e.mock.On("GetRequest",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_GetRequest_Call) Run(run func(ctx context.Context, in *types.QueryGetRequestRequest, opts ...grpc.CallOption)) *LaunchClient_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetRequestRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_GetRequest_Call) Return(_a0 *types.QueryGetRequestResponse, _a1 error) *LaunchClient_GetRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_GetRequest_Call) RunAndReturn(run func(context.Context, *types.QueryGetRequestRequest, ...grpc.CallOption) (*types.QueryGetRequestResponse, error)) *LaunchClient_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetVestingAccount provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) GetVestingAccount(ctx context.Context, in *types.QueryGetVestingAccountRequest, opts ...grpc.CallOption) (*types.QueryGetVestingAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetVestingAccount")
	}

	var r0 *types.QueryGetVestingAccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetVestingAccountRequest, ...grpc.CallOption) (*types.QueryGetVestingAccountResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetVestingAccountRequest, ...grpc.CallOption) *types.QueryGetVestingAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetVestingAccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetVestingAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_GetVestingAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVestingAccount'
type LaunchClient_GetVestingAccount_Call struct {
	*mock.Call
}

// GetVestingAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryGetVestingAccountRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) GetVestingAccount(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_GetVestingAccount_Call {
	return &LaunchClient_GetVestingAccount_Call{Call: _e.mock.On("GetVestingAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_GetVestingAccount_Call) Run(run func(ctx context.Context, in *types.QueryGetVestingAccountRequest, opts ...grpc.CallOption)) *LaunchClient_GetVestingAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryGetVestingAccountRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_GetVestingAccount_Call) Return(_a0 *types.QueryGetVestingAccountResponse, _a1 error) *LaunchClient_GetVestingAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_GetVestingAccount_Call) RunAndReturn(run func(context.Context, *types.QueryGetVestingAccountRequest, ...grpc.CallOption) (*types.QueryGetVestingAccountResponse, error)) *LaunchClient_GetVestingAccount_Call {
	_c.Call.Return(run)
	return _c
}

// ListChain provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) ListChain(ctx context.Context, in *types.QueryAllChainRequest, opts ...grpc.CallOption) (*types.QueryAllChainResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListChain")
	}

	var r0 *types.QueryAllChainResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllChainRequest, ...grpc.CallOption) (*types.QueryAllChainResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllChainRequest, ...grpc.CallOption) *types.QueryAllChainResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllChainResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllChainRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_ListChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListChain'
type LaunchClient_ListChain_Call struct {
	*mock.Call
}

// ListChain is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllChainRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) ListChain(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_ListChain_Call {
	return &LaunchClient_ListChain_Call{Call: _e.mock.On("ListChain",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_ListChain_Call) Run(run func(ctx context.Context, in *types.QueryAllChainRequest, opts ...grpc.CallOption)) *LaunchClient_ListChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllChainRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_ListChain_Call) Return(_a0 *types.QueryAllChainResponse, _a1 error) *LaunchClient_ListChain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_ListChain_Call) RunAndReturn(run func(context.Context, *types.QueryAllChainRequest, ...grpc.CallOption) (*types.QueryAllChainResponse, error)) *LaunchClient_ListChain_Call {
	_c.Call.Return(run)
	return _c
}

// ListGenesisAccount provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) ListGenesisAccount(ctx context.Context, in *types.QueryAllGenesisAccountRequest, opts ...grpc.CallOption) (*types.QueryAllGenesisAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGenesisAccount")
	}

	var r0 *types.QueryAllGenesisAccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllGenesisAccountRequest, ...grpc.CallOption) (*types.QueryAllGenesisAccountResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllGenesisAccountRequest, ...grpc.CallOption) *types.QueryAllGenesisAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllGenesisAccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllGenesisAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_ListGenesisAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListGenesisAccount'
type LaunchClient_ListGenesisAccount_Call struct {
	*mock.Call
}

// ListGenesisAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllGenesisAccountRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) ListGenesisAccount(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_ListGenesisAccount_Call {
	return &LaunchClient_ListGenesisAccount_Call{Call: _e.mock.On("ListGenesisAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_ListGenesisAccount_Call) Run(run func(ctx context.Context, in *types.QueryAllGenesisAccountRequest, opts ...grpc.CallOption)) *LaunchClient_ListGenesisAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllGenesisAccountRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_ListGenesisAccount_Call) Return(_a0 *types.QueryAllGenesisAccountResponse, _a1 error) *LaunchClient_ListGenesisAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_ListGenesisAccount_Call) RunAndReturn(run func(context.Context, *types.QueryAllGenesisAccountRequest, ...grpc.CallOption) (*types.QueryAllGenesisAccountResponse, error)) *LaunchClient_ListGenesisAccount_Call {
	_c.Call.Return(run)
	return _c
}

// ListGenesisValidator provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) ListGenesisValidator(ctx context.Context, in *types.QueryAllGenesisValidatorRequest, opts ...grpc.CallOption) (*types.QueryAllGenesisValidatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGenesisValidator")
	}

	var r0 *types.QueryAllGenesisValidatorResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllGenesisValidatorRequest, ...grpc.CallOption) (*types.QueryAllGenesisValidatorResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllGenesisValidatorRequest, ...grpc.CallOption) *types.QueryAllGenesisValidatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllGenesisValidatorResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllGenesisValidatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_ListGenesisValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListGenesisValidator'
type LaunchClient_ListGenesisValidator_Call struct {
	*mock.Call
}

// ListGenesisValidator is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllGenesisValidatorRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) ListGenesisValidator(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_ListGenesisValidator_Call {
	return &LaunchClient_ListGenesisValidator_Call{Call: _e.mock.On("ListGenesisValidator",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_ListGenesisValidator_Call) Run(run func(ctx context.Context, in *types.QueryAllGenesisValidatorRequest, opts ...grpc.CallOption)) *LaunchClient_ListGenesisValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllGenesisValidatorRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_ListGenesisValidator_Call) Return(_a0 *types.QueryAllGenesisValidatorResponse, _a1 error) *LaunchClient_ListGenesisValidator_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_ListGenesisValidator_Call) RunAndReturn(run func(context.Context, *types.QueryAllGenesisValidatorRequest, ...grpc.CallOption) (*types.QueryAllGenesisValidatorResponse, error)) *LaunchClient_ListGenesisValidator_Call {
	_c.Call.Return(run)
	return _c
}

// ListParamChange provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) ListParamChange(ctx context.Context, in *types.QueryAllParamChangeRequest, opts ...grpc.CallOption) (*types.QueryAllParamChangeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListParamChange")
	}

	var r0 *types.QueryAllParamChangeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllParamChangeRequest, ...grpc.CallOption) (*types.QueryAllParamChangeResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllParamChangeRequest, ...grpc.CallOption) *types.QueryAllParamChangeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllParamChangeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllParamChangeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_ListParamChange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListParamChange'
type LaunchClient_ListParamChange_Call struct {
	*mock.Call
}

// ListParamChange is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllParamChangeRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) ListParamChange(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_ListParamChange_Call {
	return &LaunchClient_ListParamChange_Call{Call: _e.mock.On("ListParamChange",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_ListParamChange_Call) Run(run func(ctx context.Context, in *types.QueryAllParamChangeRequest, opts ...grpc.CallOption)) *LaunchClient_ListParamChange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllParamChangeRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_ListParamChange_Call) Return(_a0 *types.QueryAllParamChangeResponse, _a1 error) *LaunchClient_ListParamChange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_ListParamChange_Call) RunAndReturn(run func(context.Context, *types.QueryAllParamChangeRequest, ...grpc.CallOption) (*types.QueryAllParamChangeResponse, error)) *LaunchClient_ListParamChange_Call {
	_c.Call.Return(run)
	return _c
}

// ListRequest provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) ListRequest(ctx context.Context, in *types.QueryAllRequestRequest, opts ...grpc.CallOption) (*types.QueryAllRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRequest")
	}

	var r0 *types.QueryAllRequestResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllRequestRequest, ...grpc.CallOption) (*types.QueryAllRequestResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllRequestRequest, ...grpc.CallOption) *types.QueryAllRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllRequestResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllRequestRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_ListRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRequest'
type LaunchClient_ListRequest_Call struct {
	*mock.Call
}

// ListRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllRequestRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) ListRequest(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_ListRequest_Call {
	return &LaunchClient_ListRequest_Call{Call: _e.mock.On("ListRequest",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_ListRequest_Call) Run(run func(ctx context.Context, in *types.QueryAllRequestRequest, opts ...grpc.CallOption)) *LaunchClient_ListRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllRequestRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_ListRequest_Call) Return(_a0 *types.QueryAllRequestResponse, _a1 error) *LaunchClient_ListRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_ListRequest_Call) RunAndReturn(run func(context.Context, *types.QueryAllRequestRequest, ...grpc.CallOption) (*types.QueryAllRequestResponse, error)) *LaunchClient_ListRequest_Call {
	_c.Call.Return(run)
	return _c
}

// ListVestingAccount provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) ListVestingAccount(ctx context.Context, in *types.QueryAllVestingAccountRequest, opts ...grpc.CallOption) (*types.QueryAllVestingAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListVestingAccount")
	}

	var r0 *types.QueryAllVestingAccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllVestingAccountRequest, ...grpc.CallOption) (*types.QueryAllVestingAccountResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllVestingAccountRequest, ...grpc.CallOption) *types.QueryAllVestingAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllVestingAccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllVestingAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchClient_ListVestingAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListVestingAccount'
type LaunchClient_ListVestingAccount_Call struct {
	*mock.Call
}

// ListVestingAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllVestingAccountRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) ListVestingAccount(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_ListVestingAccount_Call {
	return &LaunchClient_ListVestingAccount_Call{Call: _e.mock.On("ListVestingAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_ListVestingAccount_Call) Run(run func(ctx context.Context, in *types.QueryAllVestingAccountRequest, opts ...grpc.CallOption)) *LaunchClient_ListVestingAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllVestingAccountRequest), variadicArgs...)
	})
	return _c
}

func (_c *LaunchClient_ListVestingAccount_Call) Return(_a0 *types.QueryAllVestingAccountResponse, _a1 error) *LaunchClient_ListVestingAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_ListVestingAccount_Call) RunAndReturn(run func(context.Context, *types.QueryAllVestingAccountRequest, ...grpc.CallOption) (*types.QueryAllVestingAccountResponse, error)) *LaunchClient_ListVestingAccount_Call {
	_c.Call.Return(run)
	return _c
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *LaunchClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
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

// LaunchClient_Params_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Params'
type LaunchClient_Params_Call struct {
	*mock.Call
}

// Params is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryParamsRequest
//   - opts ...grpc.CallOption
func (_e *LaunchClient_Expecter) Params(ctx interface{}, in interface{}, opts ...interface{}) *LaunchClient_Params_Call {
	return &LaunchClient_Params_Call{Call: _e.mock.On("Params",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *LaunchClient_Params_Call) Run(run func(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption)) *LaunchClient_Params_Call {
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

func (_c *LaunchClient_Params_Call) Return(_a0 *types.QueryParamsResponse, _a1 error) *LaunchClient_Params_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LaunchClient_Params_Call) RunAndReturn(run func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) (*types.QueryParamsResponse, error)) *LaunchClient_Params_Call {
	_c.Call.Return(run)
	return _c
}

// NewLaunchClient creates a new instance of LaunchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLaunchClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *LaunchClient {
	mock := &LaunchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
