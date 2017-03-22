// Copyright 2017 CoreOS Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package torcx

import (
	"github.com/pkg/errors"
)

// archive is a single entry in the archives list. It contains
// an (image, version) tuple
type archive struct {
	Image string `json:"image"`
	Ref   string `json:"reference"`
}

// Archives contains a list of archives, part of a profile manifest
type Archives []archive

// CurrentProfileName returns the name of the currently running profile
func CurrentProfileName() (string, error) {
	var profile string

	meta, err := ReadMetadata(FUSE_PATH)
	if err != nil {
		return "", err
	}

	profile, ok := meta[FUSE_PROFILE_NAME]
	if !ok {
		return "", errors.New("unable to determine current profile name")
	}

	if profile == "" {
		return "", errors.New("invalid profile name")
	}

	return profile, nil
}

// CurrentProfilePath returns the path of the currently running profile
func CurrentProfilePath() (string, error) {
	var path string

	meta, err := ReadMetadata(FUSE_PATH)
	if err != nil {
		return "", err
	}

	path, ok := meta[FUSE_PROFILE_PATH]
	if !ok {
		return "", errors.New("unable to determine current profile path")
	}

	if path == "" {
		return "", errors.New("invalid profile path")
	}

	return path, nil
}

// ReadCurrentProfile returns the content of the currently running profile
func ReadCurrentProfile() (Archives, error) {
	pkglist := Archives{}

	_, err := CurrentProfilePath()
	if err != nil {
		return pkglist, err
	}

	// TODO(lucab): deserialize profile manifest from path

	return pkglist, nil
}
