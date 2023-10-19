// Code generated by mockery v2.36.0. DO NOT EDIT.

package as3

import (
	bigip "github.com/f5devcentral/go-bigip"
	mock "github.com/stretchr/testify/mock"
)

// MockBigIPIface is an autogenerated mock type for the BigIPIface type
type MockBigIPIface struct {
	mock.Mock
}

type MockBigIPIface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBigIPIface) EXPECT() *MockBigIPIface_Expecter {
	return &MockBigIPIface_Expecter{mock: &_m.Mock}
}

// APICall provides a mock function with given fields: options
func (_m *MockBigIPIface) APICall(options *bigip.APIRequest) ([]byte, error) {
	ret := _m.Called(options)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*bigip.APIRequest) ([]byte, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(*bigip.APIRequest) []byte); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*bigip.APIRequest) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBigIPIface_APICall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'APICall'
type MockBigIPIface_APICall_Call struct {
	*mock.Call
}

// APICall is a helper method to define mock.On call
//   - options *bigip.APIRequest
func (_e *MockBigIPIface_Expecter) APICall(options interface{}) *MockBigIPIface_APICall_Call {
	return &MockBigIPIface_APICall_Call{Call: _e.mock.On("APICall", options)}
}

func (_c *MockBigIPIface_APICall_Call) Run(run func(options *bigip.APIRequest)) *MockBigIPIface_APICall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bigip.APIRequest))
	})
	return _c
}

func (_c *MockBigIPIface_APICall_Call) Return(_a0 []byte, _a1 error) *MockBigIPIface_APICall_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBigIPIface_APICall_Call) RunAndReturn(run func(*bigip.APIRequest) ([]byte, error)) *MockBigIPIface_APICall_Call {
	_c.Call.Return(run)
	return _c
}

// AddInterfaceToVlan provides a mock function with given fields: vlan, iface, tagged
func (_m *MockBigIPIface) AddInterfaceToVlan(vlan string, iface string, tagged bool) error {
	ret := _m.Called(vlan, iface, tagged)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool) error); ok {
		r0 = rf(vlan, iface, tagged)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_AddInterfaceToVlan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddInterfaceToVlan'
type MockBigIPIface_AddInterfaceToVlan_Call struct {
	*mock.Call
}

// AddInterfaceToVlan is a helper method to define mock.On call
//   - vlan string
//   - iface string
//   - tagged bool
func (_e *MockBigIPIface_Expecter) AddInterfaceToVlan(vlan interface{}, iface interface{}, tagged interface{}) *MockBigIPIface_AddInterfaceToVlan_Call {
	return &MockBigIPIface_AddInterfaceToVlan_Call{Call: _e.mock.On("AddInterfaceToVlan", vlan, iface, tagged)}
}

func (_c *MockBigIPIface_AddInterfaceToVlan_Call) Run(run func(vlan string, iface string, tagged bool)) *MockBigIPIface_AddInterfaceToVlan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockBigIPIface_AddInterfaceToVlan_Call) Return(_a0 error) *MockBigIPIface_AddInterfaceToVlan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_AddInterfaceToVlan_Call) RunAndReturn(run func(string, string, bool) error) *MockBigIPIface_AddInterfaceToVlan_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSelfIP provides a mock function with given fields: config
func (_m *MockBigIPIface) CreateSelfIP(config *bigip.SelfIP) error {
	ret := _m.Called(config)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bigip.SelfIP) error); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_CreateSelfIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSelfIP'
type MockBigIPIface_CreateSelfIP_Call struct {
	*mock.Call
}

// CreateSelfIP is a helper method to define mock.On call
//   - config *bigip.SelfIP
func (_e *MockBigIPIface_Expecter) CreateSelfIP(config interface{}) *MockBigIPIface_CreateSelfIP_Call {
	return &MockBigIPIface_CreateSelfIP_Call{Call: _e.mock.On("CreateSelfIP", config)}
}

func (_c *MockBigIPIface_CreateSelfIP_Call) Run(run func(config *bigip.SelfIP)) *MockBigIPIface_CreateSelfIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bigip.SelfIP))
	})
	return _c
}

func (_c *MockBigIPIface_CreateSelfIP_Call) Return(_a0 error) *MockBigIPIface_CreateSelfIP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_CreateSelfIP_Call) RunAndReturn(run func(*bigip.SelfIP) error) *MockBigIPIface_CreateSelfIP_Call {
	_c.Call.Return(run)
	return _c
}

// CreateVlan provides a mock function with given fields: config
func (_m *MockBigIPIface) CreateVlan(config *bigip.Vlan) error {
	ret := _m.Called(config)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bigip.Vlan) error); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_CreateVlan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateVlan'
type MockBigIPIface_CreateVlan_Call struct {
	*mock.Call
}

// CreateVlan is a helper method to define mock.On call
//   - config *bigip.Vlan
func (_e *MockBigIPIface_Expecter) CreateVlan(config interface{}) *MockBigIPIface_CreateVlan_Call {
	return &MockBigIPIface_CreateVlan_Call{Call: _e.mock.On("CreateVlan", config)}
}

func (_c *MockBigIPIface_CreateVlan_Call) Run(run func(config *bigip.Vlan)) *MockBigIPIface_CreateVlan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bigip.Vlan))
	})
	return _c
}

func (_c *MockBigIPIface_CreateVlan_Call) Return(_a0 error) *MockBigIPIface_CreateVlan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_CreateVlan_Call) RunAndReturn(run func(*bigip.Vlan) error) *MockBigIPIface_CreateVlan_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRouteDomain provides a mock function with given fields: name
func (_m *MockBigIPIface) DeleteRouteDomain(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_DeleteRouteDomain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRouteDomain'
type MockBigIPIface_DeleteRouteDomain_Call struct {
	*mock.Call
}

// DeleteRouteDomain is a helper method to define mock.On call
//   - name string
func (_e *MockBigIPIface_Expecter) DeleteRouteDomain(name interface{}) *MockBigIPIface_DeleteRouteDomain_Call {
	return &MockBigIPIface_DeleteRouteDomain_Call{Call: _e.mock.On("DeleteRouteDomain", name)}
}

func (_c *MockBigIPIface_DeleteRouteDomain_Call) Run(run func(name string)) *MockBigIPIface_DeleteRouteDomain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBigIPIface_DeleteRouteDomain_Call) Return(_a0 error) *MockBigIPIface_DeleteRouteDomain_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_DeleteRouteDomain_Call) RunAndReturn(run func(string) error) *MockBigIPIface_DeleteRouteDomain_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSelfIP provides a mock function with given fields: name
func (_m *MockBigIPIface) DeleteSelfIP(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_DeleteSelfIP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSelfIP'
type MockBigIPIface_DeleteSelfIP_Call struct {
	*mock.Call
}

// DeleteSelfIP is a helper method to define mock.On call
//   - name string
func (_e *MockBigIPIface_Expecter) DeleteSelfIP(name interface{}) *MockBigIPIface_DeleteSelfIP_Call {
	return &MockBigIPIface_DeleteSelfIP_Call{Call: _e.mock.On("DeleteSelfIP", name)}
}

func (_c *MockBigIPIface_DeleteSelfIP_Call) Run(run func(name string)) *MockBigIPIface_DeleteSelfIP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBigIPIface_DeleteSelfIP_Call) Return(_a0 error) *MockBigIPIface_DeleteSelfIP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_DeleteSelfIP_Call) RunAndReturn(run func(string) error) *MockBigIPIface_DeleteSelfIP_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteVlan provides a mock function with given fields: name
func (_m *MockBigIPIface) DeleteVlan(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_DeleteVlan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteVlan'
type MockBigIPIface_DeleteVlan_Call struct {
	*mock.Call
}

// DeleteVlan is a helper method to define mock.On call
//   - name string
func (_e *MockBigIPIface_Expecter) DeleteVlan(name interface{}) *MockBigIPIface_DeleteVlan_Call {
	return &MockBigIPIface_DeleteVlan_Call{Call: _e.mock.On("DeleteVlan", name)}
}

func (_c *MockBigIPIface_DeleteVlan_Call) Run(run func(name string)) *MockBigIPIface_DeleteVlan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBigIPIface_DeleteVlan_Call) Return(_a0 error) *MockBigIPIface_DeleteVlan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_DeleteVlan_Call) RunAndReturn(run func(string) error) *MockBigIPIface_DeleteVlan_Call {
	_c.Call.Return(run)
	return _c
}

// GetDevices provides a mock function with given fields:
func (_m *MockBigIPIface) GetDevices() ([]bigip.Device, error) {
	ret := _m.Called()

	var r0 []bigip.Device
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]bigip.Device, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []bigip.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bigip.Device)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBigIPIface_GetDevices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDevices'
type MockBigIPIface_GetDevices_Call struct {
	*mock.Call
}

// GetDevices is a helper method to define mock.On call
func (_e *MockBigIPIface_Expecter) GetDevices() *MockBigIPIface_GetDevices_Call {
	return &MockBigIPIface_GetDevices_Call{Call: _e.mock.On("GetDevices")}
}

func (_c *MockBigIPIface_GetDevices_Call) Run(run func()) *MockBigIPIface_GetDevices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBigIPIface_GetDevices_Call) Return(_a0 []bigip.Device, _a1 error) *MockBigIPIface_GetDevices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBigIPIface_GetDevices_Call) RunAndReturn(run func() ([]bigip.Device, error)) *MockBigIPIface_GetDevices_Call {
	_c.Call.Return(run)
	return _c
}

// GetVlanInterfaces provides a mock function with given fields: vlan
func (_m *MockBigIPIface) GetVlanInterfaces(vlan string) (*bigip.VlanInterfaces, error) {
	ret := _m.Called(vlan)

	var r0 *bigip.VlanInterfaces
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*bigip.VlanInterfaces, error)); ok {
		return rf(vlan)
	}
	if rf, ok := ret.Get(0).(func(string) *bigip.VlanInterfaces); ok {
		r0 = rf(vlan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bigip.VlanInterfaces)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(vlan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBigIPIface_GetVlanInterfaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVlanInterfaces'
type MockBigIPIface_GetVlanInterfaces_Call struct {
	*mock.Call
}

// GetVlanInterfaces is a helper method to define mock.On call
//   - vlan string
func (_e *MockBigIPIface_Expecter) GetVlanInterfaces(vlan interface{}) *MockBigIPIface_GetVlanInterfaces_Call {
	return &MockBigIPIface_GetVlanInterfaces_Call{Call: _e.mock.On("GetVlanInterfaces", vlan)}
}

func (_c *MockBigIPIface_GetVlanInterfaces_Call) Run(run func(vlan string)) *MockBigIPIface_GetVlanInterfaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBigIPIface_GetVlanInterfaces_Call) Return(_a0 *bigip.VlanInterfaces, _a1 error) *MockBigIPIface_GetVlanInterfaces_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBigIPIface_GetVlanInterfaces_Call) RunAndReturn(run func(string) (*bigip.VlanInterfaces, error)) *MockBigIPIface_GetVlanInterfaces_Call {
	_c.Call.Return(run)
	return _c
}

// PostAs3Bigip provides a mock function with given fields: as3NewJson, tenantFilter
func (_m *MockBigIPIface) PostAs3Bigip(as3NewJson string, tenantFilter string) (error, string, string) {
	ret := _m.Called(as3NewJson, tenantFilter)

	var r0 error
	var r1 string
	var r2 string
	if rf, ok := ret.Get(0).(func(string, string) (error, string, string)); ok {
		return rf(as3NewJson, tenantFilter)
	}
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(as3NewJson, tenantFilter)
	} else {
		r0 = ret.Error(0)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(as3NewJson, tenantFilter)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) string); ok {
		r2 = rf(as3NewJson, tenantFilter)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// MockBigIPIface_PostAs3Bigip_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostAs3Bigip'
type MockBigIPIface_PostAs3Bigip_Call struct {
	*mock.Call
}

// PostAs3Bigip is a helper method to define mock.On call
//   - as3NewJson string
//   - tenantFilter string
func (_e *MockBigIPIface_Expecter) PostAs3Bigip(as3NewJson interface{}, tenantFilter interface{}) *MockBigIPIface_PostAs3Bigip_Call {
	return &MockBigIPIface_PostAs3Bigip_Call{Call: _e.mock.On("PostAs3Bigip", as3NewJson, tenantFilter)}
}

func (_c *MockBigIPIface_PostAs3Bigip_Call) Run(run func(as3NewJson string, tenantFilter string)) *MockBigIPIface_PostAs3Bigip_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockBigIPIface_PostAs3Bigip_Call) Return(_a0 error, _a1 string, _a2 string) *MockBigIPIface_PostAs3Bigip_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockBigIPIface_PostAs3Bigip_Call) RunAndReturn(run func(string, string) (error, string, string)) *MockBigIPIface_PostAs3Bigip_Call {
	_c.Call.Return(run)
	return _c
}

// SelfIPs provides a mock function with given fields:
func (_m *MockBigIPIface) SelfIPs() (*bigip.SelfIPs, error) {
	ret := _m.Called()

	var r0 *bigip.SelfIPs
	var r1 error
	if rf, ok := ret.Get(0).(func() (*bigip.SelfIPs, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *bigip.SelfIPs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bigip.SelfIPs)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBigIPIface_SelfIPs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelfIPs'
type MockBigIPIface_SelfIPs_Call struct {
	*mock.Call
}

// SelfIPs is a helper method to define mock.On call
func (_e *MockBigIPIface_Expecter) SelfIPs() *MockBigIPIface_SelfIPs_Call {
	return &MockBigIPIface_SelfIPs_Call{Call: _e.mock.On("SelfIPs")}
}

func (_c *MockBigIPIface_SelfIPs_Call) Run(run func()) *MockBigIPIface_SelfIPs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBigIPIface_SelfIPs_Call) Return(_a0 *bigip.SelfIPs, _a1 error) *MockBigIPIface_SelfIPs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBigIPIface_SelfIPs_Call) RunAndReturn(run func() (*bigip.SelfIPs, error)) *MockBigIPIface_SelfIPs_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateVcmpGuest provides a mock function with given fields: name, config
func (_m *MockBigIPIface) UpdateVcmpGuest(name string, config *bigip.VcmpGuest) error {
	ret := _m.Called(name, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *bigip.VcmpGuest) error); ok {
		r0 = rf(name, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBigIPIface_UpdateVcmpGuest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateVcmpGuest'
type MockBigIPIface_UpdateVcmpGuest_Call struct {
	*mock.Call
}

// UpdateVcmpGuest is a helper method to define mock.On call
//   - name string
//   - config *bigip.VcmpGuest
func (_e *MockBigIPIface_Expecter) UpdateVcmpGuest(name interface{}, config interface{}) *MockBigIPIface_UpdateVcmpGuest_Call {
	return &MockBigIPIface_UpdateVcmpGuest_Call{Call: _e.mock.On("UpdateVcmpGuest", name, config)}
}

func (_c *MockBigIPIface_UpdateVcmpGuest_Call) Run(run func(name string, config *bigip.VcmpGuest)) *MockBigIPIface_UpdateVcmpGuest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*bigip.VcmpGuest))
	})
	return _c
}

func (_c *MockBigIPIface_UpdateVcmpGuest_Call) Return(_a0 error) *MockBigIPIface_UpdateVcmpGuest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBigIPIface_UpdateVcmpGuest_Call) RunAndReturn(run func(string, *bigip.VcmpGuest) error) *MockBigIPIface_UpdateVcmpGuest_Call {
	_c.Call.Return(run)
	return _c
}

// Vlans provides a mock function with given fields:
func (_m *MockBigIPIface) Vlans() (*bigip.Vlans, error) {
	ret := _m.Called()

	var r0 *bigip.Vlans
	var r1 error
	if rf, ok := ret.Get(0).(func() (*bigip.Vlans, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *bigip.Vlans); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bigip.Vlans)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBigIPIface_Vlans_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Vlans'
type MockBigIPIface_Vlans_Call struct {
	*mock.Call
}

// Vlans is a helper method to define mock.On call
func (_e *MockBigIPIface_Expecter) Vlans() *MockBigIPIface_Vlans_Call {
	return &MockBigIPIface_Vlans_Call{Call: _e.mock.On("Vlans")}
}

func (_c *MockBigIPIface_Vlans_Call) Run(run func()) *MockBigIPIface_Vlans_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBigIPIface_Vlans_Call) Return(_a0 *bigip.Vlans, _a1 error) *MockBigIPIface_Vlans_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBigIPIface_Vlans_Call) RunAndReturn(run func() (*bigip.Vlans, error)) *MockBigIPIface_Vlans_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBigIPIface creates a new instance of MockBigIPIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBigIPIface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBigIPIface {
	mock := &MockBigIPIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}