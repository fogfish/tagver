//
// Copyright (C) 2010 - 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/dynamo
//

package tagver_test

import (
	"testing"

	"github.com/fogfish/tagver"
)

func TestIsTest(t *testing.T) {
	if !tagver.IsTest("pr1") {
		t.Errorf("pr1 MUST be TEST")
	}

	if tagver.IsTest("v1") {
		t.Errorf("v1 MUST NOT be TEST")
	}

	if tagver.IsTest("main") {
		t.Errorf("main MUST NOT be TEST")
	}
}

func TestIsLive(t *testing.T) {
	if tagver.IsLive("pr1") {
		t.Errorf("pr1 MUST NOT be LIVE")
	}

	if !tagver.IsLive("v1") {
		t.Errorf("v1 MUST be LIVE")
	}

	if tagver.IsLive("main") {
		t.Errorf("main MUST NOT be LIVE")
	}
}

func TestIsMain(t *testing.T) {
	if tagver.IsMain("pr1") {
		t.Errorf("pr1 MUST NOT be MAIN")
	}

	if tagver.IsMain("v1") {
		t.Errorf("v1 MUST NOT be MAIN")
	}

	if !tagver.IsMain("main") {
		t.Errorf("main MUST be MAIN")
	}
}

func TestVersions(t *testing.T) {
	for input, expected := range map[string]tagver.Versions{
		"api":                            {},
		"api@v1":                         {"api": "v1"},
		"api@v1:":                        {"api": "v1"},
		"api@v1:db@main":                 {"api": "v1", "db": "main"},
		"pfx-api@v1:pfx-db@main":         {"pfx-api": "v1", "pfx-db": "main"},
		"pfx-api@sfx-v1:pfx-db@sfx-main": {"pfx-api": "sfx-v1", "pfx-db": "sfx-main"},
		"api@v1db@main":                  {},
		"api:db@main":                    {"db": "main"},
		"api:db:":                        {},
	} {
		vsn := tagver.NewVersions(input)
		for k, v := range vsn {
			if s, has := expected[k]; !has || s != v {
				t.Errorf("failed to parse %v, %v expected (%s)", input, k, v)
			}
		}

		for k, v := range expected {
			if s, has := vsn[k]; !has || s != v {
				t.Errorf("failed to parse %v, %v expected (%s)", input, k, v)
			}
		}
	}
}

func TestVersionsGet(t *testing.T) {
	vsn := tagver.NewVersions("api@pr1:db@v1")

	if vsn.Get("api", "main") != "pr1" {
		t.Errorf("invalid api key value")
	}

	if vsn.Get("db", "main") != "v1" {
		t.Errorf("invalid db key value")
	}

	if vsn.Get("mw", "main") != "main" {
		t.Errorf("invalid mw key value")
	}
}
