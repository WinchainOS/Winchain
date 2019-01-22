// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
)

const (
	VersionMajor = 1          // Major version component of the current release
	VersionMinor = 8          // Minor version component of the current release
	VersionPatch = 16         // Patch version component of the current release
	VersionMeta  = "stable" // Version metadata to append to the version string

	//TODO【Winchain】================the part had been changed==================
	WinchainVersionMajor = 1
	WinchainVersionMinor = 0
	WinchainVersionPatch = 0
)

// Version holds the textual version string.
var Version = func() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}()

//TODO【Winchain】================the part had been changed==================
// Version holds the textual version string.
// 联盟链的版本信息
var WinchainVersion = func() string {
	return fmt.Sprintf("%d.%d.%d", WinchainVersionMajor, WinchainVersionMinor, WinchainVersionPatch)
}()

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := Version
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()

// ArchiveVersion holds the textual version string used for Geth archives.
// e.g. "1.8.11-dea1ce05" for stable releases, or
//      "1.8.13-unstable-21c059b6" for unstable releases
func ArchiveVersion(gitCommit string) string {
	vsn := Version
	if VersionMeta != "stable" {
		vsn += "-" + VersionMeta
	}
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}

func VersionWithCommit(gitCommit string) string {
	vsn := VersionWithMeta
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}

//TODO【Winchain】================the part had been changed==================
func WinchainVersionWithCommit(gitCommit string) string {
	vsn := WinchainVersion
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}


