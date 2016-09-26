/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package web

import (
	"testing"
)

func TestLookupTables(tt *testing.T) {
	t := newTest(tt)

	testPing(t)
	testGetSupportedPermissions(t)
	testGetSupportedClusterTypes(t)
	testGetSupportedEntityTypes(t)
}

func testPing(t *test) {
	expected := "hello"
	actual, err := t.svc.PingServer(t.su, expected)
	t.nil(err)
	t.ok(actual == expected, "ping mismatch")
}

func testGetSupportedPermissions(t *test) {
	expected := []string{
		"ManageCluster",
		"ManageEngine",
		"ManageIdentity",
		"ManageDataset",
		"ManageDatasource",
		"ManageModel",
		"ManageLabel",
		"ManageProject",
		"ManageRole",
		"ManageService",
		"ManageWorkgroup",
		"ViewCluster",
		"ViewEngine",
		"ViewIdentity",
		"ViewDataset",
		"ViewDatasource",
		"ViewModel",
		"ViewLabel",
		"ViewProject",
		"ViewRole",
		"ViewService",
		"ViewWorkgroup",
	}

	perms, err := t.svc.GetAllPermissions(t.su)
	t.nil(err)

	if len(perms) != len(expected) {
		t.fail("expected %d permissions", len(expected))
	}

	actual := make(map[string]bool)
	for _, p := range perms {
		actual[p.Code] = true
	}

	for _, e := range expected {
		if _, ok := actual[e]; !ok {
			t.fail("permission not found: %s", e)
		}
	}
}

func testGetSupportedClusterTypes(t *test) {
	expected := []string{
		"external",
		"yarn",
	}

	cts, err := t.svc.GetAllClusterTypes(t.su)
	t.nil(err)

	if len(cts) != len(expected) {
		t.fail("expected %d cluster types", len(expected))
	}

	actual := make(map[string]bool)
	for _, ct := range cts {
		actual[ct.Name] = true
	}

	for _, e := range expected {
		if _, ok := actual[e]; !ok {
			t.fail("cluster type not found: %s", e)
		}
	}
}

func testGetSupportedEntityTypes(t *test) {
	expected := []string{
		"role",
		"workgroup",
		"identity",
		"engine",
		"cluster",
		"project",
		"dataset",
		"datasource",
		"label",
		"model",
		"service",
	}

	ets, err := t.svc.GetAllEntityTypes(t.su)
	t.nil(err)

	if len(ets) != len(expected) {
		t.fail("expected %d entity types", len(expected))
	}

	actual := make(map[string]bool)
	for _, et := range ets {
		actual[et.Name] = true
	}

	for _, e := range expected {
		if _, ok := actual[e]; !ok {
			t.fail("entity type not found: %s", e)
		}
	}
}
