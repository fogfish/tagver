//
// Copyright (C) 2010 - 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/dynamo
//

package tagver

import "strings"

// Identifier a deployment
// - for a change proposal (TEST)
// - latest snapshot (MAIN)
// - from tag/release (LIVE).
type Version string

// Is Version identifying a deployment for a change proposal.
func IsTest(vsn Version) bool {
	return strings.HasPrefix(string(vsn), "pr")
}

// Is Version identifying a production deployment from tag/release.
func IsLive(vsn Version) bool {
	return strings.HasPrefix(string(vsn), "v")
}

// Is Version identifying a deployment of latest snapshot from branch.
func IsMain(vsn Version) bool {
	return !IsTest(vsn) && !IsLive(vsn)
}

// Tag cloud resource
func (vsn Version) Tag(name string) string {
	if vsn == "" {
		return name
	}

	return name + "-" + string(vsn)
}

// Software MAY establish dependencies between stacks (e.g. api layer depends on the database).
// Versions are established as key-value mapping
type Versions map[string]Version

const divider_vsn = "@"
const divider_dep = ":"

func NewVersions(v string) Versions {
	vsn := Versions{}

	for _, stack := range strings.Split(v, divider_dep) {
		seq := strings.Split(stack, divider_vsn)
		if len(seq) == 2 {
			vsn[seq[0]] = Version(seq[1])
		}
	}

	return vsn
}

func (vsn Versions) Get(key, def string) Version {
	if v, has := vsn[key]; has {
		return v
	}

	return Version(def)
}
