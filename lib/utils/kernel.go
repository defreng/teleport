/*
Copyright 2019 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/gravitational/teleport"

	"github.com/gravitational/trace"

	"github.com/coreos/go-semver/semver"
)

// KernelVersion returns the kernel version of the host. This only returns
// something meaningful on Linux.
func KernelVersion() (*semver.Version, error) {
	if runtime.GOOS != teleport.LinuxOS {
		return nil, trace.BadParameter("requested kernel version on non-Linux host")
	}

	buf, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return nil, trace.Wrap(err)
	}

	ver, err := semver.NewVersion(strings.TrimSpace(string(buf)))
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return ver, nil
}