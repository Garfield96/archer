// Copyright 2023 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package f5

import (
	"net/http"
	"testing"

	"github.com/f5devcentral/go-bigip"
	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/fixture"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"

	"github.com/sapcc/archer/internal/agent/f5/as3"
	"github.com/sapcc/archer/internal/neutron"
	"github.com/sapcc/archer/models"
)

func TestAgent_EnsureSelfIPs_Create(t *testing.T) {
	dbMock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dbMock.Close()
	}()

	th.SetupPersistentPortHTTP(t, 8931)
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/v2.0/subnets/e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e", "GET", "",
		GetSubnetResponseFixture, http.StatusOK)
	fixture.SetupHandler(t, "/v2.0/ports", "GET", "",
		GetPortListResponseFixture, http.StatusOK)

	neutronClient := neutron.NeutronClient{ServiceClient: fake.ServiceClient()}
	neutronClient.InitCache()

	bigIPMock := as3.NewMockBigIPIface(t)
	// we don't have the selfip yet, let it create it
	bigIPMock.EXPECT().
		SelfIP("selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4").
		Return(nil, nil)
	bigIPMock.EXPECT().
		CreateSelfIP(&bigip.SelfIP{
			Name:    "selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4",
			Address: "42.42.42.42%1234/8",
			Vlan:    "/Common/vlan-1234",
		}).
		Return(nil)

	a := &Agent{
		jobQueue: nil,
		pool:     dbMock,
		neutron:  &neutronClient,
		bigips:   []*as3.BigIP{{Host: "dummybigiphost", BigIPIface: bigIPMock}},
		vcmps:    []*as3.BigIP{},
		bigip:    &as3.BigIP{Host: "dummybigiphost", BigIPIface: bigIPMock},
	}

	epPort := ports.Port{
		ID: "af813e79-8038-4a0a-bf57-9c0bb8d65c2a",
		FixedIPs: []ports.IP{
			{
				IPAddress: "6.6.6.6",
				SubnetID:  "e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e",
			},
		},
	}

	assert.Nil(t, a.EnsureSelfIPs(1234, &epPort), "EnsureSelfIPs() should not return an error")
}

func TestAgent_EnsureSelfIPs_NoOp(t *testing.T) {
	dbMock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dbMock.Close()
	}()

	th.SetupPersistentPortHTTP(t, 8931)
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/v2.0/subnets/e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e", "GET", "",
		GetSubnetResponseFixture, http.StatusOK)
	fixture.SetupHandler(t, "/v2.0/ports", "GET", "",
		GetPortListResponseFixture, http.StatusOK)

	neutronClient := neutron.NeutronClient{ServiceClient: fake.ServiceClient()}
	neutronClient.InitCache()

	bigIPMock := as3.NewMockBigIPIface(t)
	// we don't have the selfip yet, let it create it
	bigIPMock.EXPECT().
		SelfIP("selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4").
		Return(&bigip.SelfIP{
			Name:    "selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4",
			Address: "42.42.42.42%1234/8",
			Vlan:    "/Common/vlan-1234",
		}, nil)

	a := &Agent{
		jobQueue: nil,
		pool:     dbMock,
		neutron:  &neutronClient,
		bigips:   []*as3.BigIP{{Host: "dummybigiphost", BigIPIface: bigIPMock}},
		vcmps:    []*as3.BigIP{},
		bigip:    &as3.BigIP{Host: "dummybigiphost", BigIPIface: bigIPMock},
	}

	epPort := ports.Port{
		ID: "af813e79-8038-4a0a-bf57-9c0bb8d65c2a",
		FixedIPs: []ports.IP{
			{
				IPAddress: "6.6.6.6",
				SubnetID:  "e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e",
			},
		},
	}

	assert.Nil(t, a.EnsureSelfIPs(1234, &epPort), "EnsureSelfIPs() should not return an error")
}

func TestAgent_CleanupSelfIPs(t *testing.T) {
	dbMock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dbMock.Close()
	}()

	th.SetupPersistentPortHTTP(t, 8931)
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/v2.0/subnets/e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e", "GET", "",
		GetSubnetResponseFixture, http.StatusOK)
	fixture.SetupHandler(t, "/v2.0/ports", "GET", "",
		GetPortListResponseFixture, http.StatusOK)
	fixture.SetupHandler(t, "/v2.0/ports/5a8ad669-4ffe-4133-b9f9-6de62cd654a4", "DELETE", "",
		"", http.StatusAccepted)

	neutronClient := neutron.NeutronClient{ServiceClient: fake.ServiceClient()}
	neutronClient.InitCache()

	bigIPMock := as3.NewMockBigIPIface(t)
	// we don't have the selfip yet, let it create it
	bigIPMock.EXPECT().
		SelfIP("selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4").
		Return(&bigip.SelfIP{
			Name:    "selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4",
			Address: "42.42.42.42%1234/8",
			Vlan:    "/Common/vlan-1234",
		}, nil)
	bigIPMock.EXPECT().
		DeleteSelfIP("selfip-5a8ad669-4ffe-4133-b9f9-6de62cd654a4").
		Return(nil)

	a := &Agent{
		jobQueue: nil,
		pool:     dbMock,
		neutron:  &neutronClient,
		bigips:   []*as3.BigIP{{Host: "dummybigiphost", BigIPIface: bigIPMock}},
		vcmps:    []*as3.BigIP{},
		bigip:    &as3.BigIP{Host: "dummybigiphost", BigIPIface: bigIPMock},
	}

	epPort := ports.Port{
		ID: "af813e79-8038-4a0a-bf57-9c0bb8d65c2a",
		FixedIPs: []ports.IP{
			{
				IPAddress: "6.6.6.6",
				SubnetID:  "e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e",
			},
		},
	}

	allPorts := []*as3.ExtendedEndpoint{{
		Endpoint: models.Endpoint{Status: models.EndpointStatusPENDINGDELETE},
		Port:     &epPort,
	}}

	// Port should be deleted
	assert.Nil(t, a.CleanupSelfIPs(&epPort, allPorts), "CleanupSelfIPs() should not return an error")
}

func TestAgent_CleanupSelfIPs_NoOp(t *testing.T) {
	dbMock, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dbMock.Close()
	}()

	th.SetupPersistentPortHTTP(t, 8931)
	defer th.TeardownHTTP()
	neutronClient := neutron.NeutronClient{ServiceClient: fake.ServiceClient()}
	neutronClient.InitCache()

	bigIPMock := as3.NewMockBigIPIface(t)
	a := &Agent{
		jobQueue: nil,
		pool:     dbMock,
		neutron:  &neutronClient,
		bigips:   []*as3.BigIP{{Host: "dummybigiphost", BigIPIface: bigIPMock}},
		vcmps:    []*as3.BigIP{},
		bigip:    &as3.BigIP{Host: "dummybigiphost", BigIPIface: bigIPMock},
	}

	epPort := ports.Port{
		ID: "af813e79-8038-4a0a-bf57-9c0bb8d65c2a",
		FixedIPs: []ports.IP{
			{
				IPAddress: "6.6.6.6",
				SubnetID:  "e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e",
			},
		},
	}

	allPorts := []*as3.ExtendedEndpoint{{
		Endpoint: models.Endpoint{Status: models.EndpointStatusPENDINGDELETE},
		Port:     &epPort,
	}, {
		Endpoint: models.Endpoint{Status: models.EndpointStatusACTIVE},
		Port: &ports.Port{
			FixedIPs: []ports.IP{{SubnetID: "e0e0e0e0-e0e0-4e0e-8e0e-0e0e0e0e0e0e"}},
		},
	}}

	// SelfIP Port should not be deleted, since we have an existing endpoint in the same subnet
	assert.Nil(t, a.CleanupSelfIPs(&epPort, allPorts), "CleanupSelfIPs() should not return an error")
	assert.Nil(t, bigIPMock.Calls)
}
