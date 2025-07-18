// Copyright 2025 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/util/netlink (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/agent/util/netlink/testing/mock_netlink_linux.go -package testing antrea.io/antrea/pkg/agent/util/netlink Interface
//

// Package testing is a generated GoMock package.
package testing

import (
	net "net"
	reflect "reflect"

	netlink "github.com/vishvananda/netlink"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AddrAdd mocks base method.
func (m *MockInterface) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrAdd", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrAdd indicates an expected call of AddrAdd.
func (mr *MockInterfaceMockRecorder) AddrAdd(link, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrAdd", reflect.TypeOf((*MockInterface)(nil).AddrAdd), link, addr)
}

// AddrDel mocks base method.
func (m *MockInterface) AddrDel(link netlink.Link, addr *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrDel", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrDel indicates an expected call of AddrDel.
func (mr *MockInterfaceMockRecorder) AddrDel(link, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrDel", reflect.TypeOf((*MockInterface)(nil).AddrDel), link, addr)
}

// AddrList mocks base method.
func (m *MockInterface) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrList", link, family)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddrList indicates an expected call of AddrList.
func (mr *MockInterfaceMockRecorder) AddrList(link, family any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrList", reflect.TypeOf((*MockInterface)(nil).AddrList), link, family)
}

// AddrReplace mocks base method.
func (m *MockInterface) AddrReplace(link netlink.Link, addr *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrReplace", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrReplace indicates an expected call of AddrReplace.
func (mr *MockInterfaceMockRecorder) AddrReplace(link, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrReplace", reflect.TypeOf((*MockInterface)(nil).AddrReplace), link, addr)
}

// ConntrackDeleteFilter mocks base method.
func (m *MockInterface) ConntrackDeleteFilter(table netlink.ConntrackTableType, family netlink.InetFamily, filter netlink.CustomConntrackFilter) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConntrackDeleteFilter", table, family, filter)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConntrackDeleteFilter indicates an expected call of ConntrackDeleteFilter.
func (mr *MockInterfaceMockRecorder) ConntrackDeleteFilter(table, family, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConntrackDeleteFilter", reflect.TypeOf((*MockInterface)(nil).ConntrackDeleteFilter), table, family, filter)
}

// LinkAddAltName mocks base method.
func (m *MockInterface) LinkAddAltName(link netlink.Link, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkAddAltName", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkAddAltName indicates an expected call of LinkAddAltName.
func (mr *MockInterfaceMockRecorder) LinkAddAltName(link, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkAddAltName", reflect.TypeOf((*MockInterface)(nil).LinkAddAltName), link, name)
}

// LinkByIndex mocks base method.
func (m *MockInterface) LinkByIndex(index int) (netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkByIndex", index)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByIndex indicates an expected call of LinkByIndex.
func (mr *MockInterfaceMockRecorder) LinkByIndex(index any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByIndex", reflect.TypeOf((*MockInterface)(nil).LinkByIndex), index)
}

// LinkByName mocks base method.
func (m *MockInterface) LinkByName(name string) (netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkByName", name)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByName indicates an expected call of LinkByName.
func (mr *MockInterfaceMockRecorder) LinkByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByName", reflect.TypeOf((*MockInterface)(nil).LinkByName), name)
}

// LinkDelAltName mocks base method.
func (m *MockInterface) LinkDelAltName(link netlink.Link, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkDelAltName", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDelAltName indicates an expected call of LinkDelAltName.
func (mr *MockInterfaceMockRecorder) LinkDelAltName(link, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDelAltName", reflect.TypeOf((*MockInterface)(nil).LinkDelAltName), link, name)
}

// LinkList mocks base method.
func (m *MockInterface) LinkList() ([]netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkList")
	ret0, _ := ret[0].([]netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkList indicates an expected call of LinkList.
func (mr *MockInterfaceMockRecorder) LinkList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkList", reflect.TypeOf((*MockInterface)(nil).LinkList))
}

// LinkSetAlias mocks base method.
func (m *MockInterface) LinkSetAlias(link netlink.Link, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetAlias", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetAlias indicates an expected call of LinkSetAlias.
func (mr *MockInterfaceMockRecorder) LinkSetAlias(link, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetAlias", reflect.TypeOf((*MockInterface)(nil).LinkSetAlias), link, name)
}

// LinkSetDown mocks base method.
func (m *MockInterface) LinkSetDown(link netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetDown", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetDown indicates an expected call of LinkSetDown.
func (mr *MockInterfaceMockRecorder) LinkSetDown(link any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetDown", reflect.TypeOf((*MockInterface)(nil).LinkSetDown), link)
}

// LinkSetHardwareAddr mocks base method.
func (m *MockInterface) LinkSetHardwareAddr(link netlink.Link, hwaddr net.HardwareAddr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetHardwareAddr", link, hwaddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetHardwareAddr indicates an expected call of LinkSetHardwareAddr.
func (mr *MockInterfaceMockRecorder) LinkSetHardwareAddr(link, hwaddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetHardwareAddr", reflect.TypeOf((*MockInterface)(nil).LinkSetHardwareAddr), link, hwaddr)
}

// LinkSetMTU mocks base method.
func (m *MockInterface) LinkSetMTU(link netlink.Link, mtu int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetMTU", link, mtu)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetMTU indicates an expected call of LinkSetMTU.
func (mr *MockInterfaceMockRecorder) LinkSetMTU(link, mtu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetMTU", reflect.TypeOf((*MockInterface)(nil).LinkSetMTU), link, mtu)
}

// LinkSetName mocks base method.
func (m *MockInterface) LinkSetName(link netlink.Link, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetName", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetName indicates an expected call of LinkSetName.
func (mr *MockInterfaceMockRecorder) LinkSetName(link, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetName", reflect.TypeOf((*MockInterface)(nil).LinkSetName), link, name)
}

// LinkSetNsFd mocks base method.
func (m *MockInterface) LinkSetNsFd(link netlink.Link, fd int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetNsFd", link, fd)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetNsFd indicates an expected call of LinkSetNsFd.
func (mr *MockInterfaceMockRecorder) LinkSetNsFd(link, fd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetNsFd", reflect.TypeOf((*MockInterface)(nil).LinkSetNsFd), link, fd)
}

// LinkSetUp mocks base method.
func (m *MockInterface) LinkSetUp(link netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetUp", link)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetUp indicates an expected call of LinkSetUp.
func (mr *MockInterfaceMockRecorder) LinkSetUp(link any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetUp", reflect.TypeOf((*MockInterface)(nil).LinkSetUp), link)
}

// NeighDel mocks base method.
func (m *MockInterface) NeighDel(neigh *netlink.Neigh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeighDel", neigh)
	ret0, _ := ret[0].(error)
	return ret0
}

// NeighDel indicates an expected call of NeighDel.
func (mr *MockInterfaceMockRecorder) NeighDel(neigh any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeighDel", reflect.TypeOf((*MockInterface)(nil).NeighDel), neigh)
}

// NeighList mocks base method.
func (m *MockInterface) NeighList(linkIndex, family int) ([]netlink.Neigh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeighList", linkIndex, family)
	ret0, _ := ret[0].([]netlink.Neigh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeighList indicates an expected call of NeighList.
func (mr *MockInterfaceMockRecorder) NeighList(linkIndex, family any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeighList", reflect.TypeOf((*MockInterface)(nil).NeighList), linkIndex, family)
}

// NeighListExecute mocks base method.
func (m *MockInterface) NeighListExecute(msg netlink.Ndmsg) ([]netlink.Neigh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeighListExecute", msg)
	ret0, _ := ret[0].([]netlink.Neigh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeighListExecute indicates an expected call of NeighListExecute.
func (mr *MockInterfaceMockRecorder) NeighListExecute(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeighListExecute", reflect.TypeOf((*MockInterface)(nil).NeighListExecute), msg)
}

// NeighSet mocks base method.
func (m *MockInterface) NeighSet(neigh *netlink.Neigh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeighSet", neigh)
	ret0, _ := ret[0].(error)
	return ret0
}

// NeighSet indicates an expected call of NeighSet.
func (mr *MockInterfaceMockRecorder) NeighSet(neigh any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeighSet", reflect.TypeOf((*MockInterface)(nil).NeighSet), neigh)
}

// RouteAdd mocks base method.
func (m *MockInterface) RouteAdd(route *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteAdd", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteAdd indicates an expected call of RouteAdd.
func (mr *MockInterfaceMockRecorder) RouteAdd(route any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteAdd", reflect.TypeOf((*MockInterface)(nil).RouteAdd), route)
}

// RouteDel mocks base method.
func (m *MockInterface) RouteDel(route *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteDel", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteDel indicates an expected call of RouteDel.
func (mr *MockInterfaceMockRecorder) RouteDel(route any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDel", reflect.TypeOf((*MockInterface)(nil).RouteDel), route)
}

// RouteList mocks base method.
func (m *MockInterface) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteList", link, family)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteList indicates an expected call of RouteList.
func (mr *MockInterfaceMockRecorder) RouteList(link, family any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteList", reflect.TypeOf((*MockInterface)(nil).RouteList), link, family)
}

// RouteListFiltered mocks base method.
func (m *MockInterface) RouteListFiltered(family int, filter *netlink.Route, filterMask uint64) ([]netlink.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteListFiltered", family, filter, filterMask)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteListFiltered indicates an expected call of RouteListFiltered.
func (mr *MockInterfaceMockRecorder) RouteListFiltered(family, filter, filterMask any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteListFiltered", reflect.TypeOf((*MockInterface)(nil).RouteListFiltered), family, filter, filterMask)
}

// RouteReplace mocks base method.
func (m *MockInterface) RouteReplace(route *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteReplace", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteReplace indicates an expected call of RouteReplace.
func (mr *MockInterfaceMockRecorder) RouteReplace(route any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteReplace", reflect.TypeOf((*MockInterface)(nil).RouteReplace), route)
}

// RuleAdd mocks base method.
func (m *MockInterface) RuleAdd(rule *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleAdd", rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleAdd indicates an expected call of RuleAdd.
func (mr *MockInterfaceMockRecorder) RuleAdd(rule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleAdd", reflect.TypeOf((*MockInterface)(nil).RuleAdd), rule)
}

// RuleDel mocks base method.
func (m *MockInterface) RuleDel(rule *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleDel", rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleDel indicates an expected call of RuleDel.
func (mr *MockInterfaceMockRecorder) RuleDel(rule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleDel", reflect.TypeOf((*MockInterface)(nil).RuleDel), rule)
}

// RuleList mocks base method.
func (m *MockInterface) RuleList(family int) ([]netlink.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleList", family)
	ret0, _ := ret[0].([]netlink.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RuleList indicates an expected call of RuleList.
func (mr *MockInterfaceMockRecorder) RuleList(family any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleList", reflect.TypeOf((*MockInterface)(nil).RuleList), family)
}
