/*
Copyright (c) 2019 Red Hat, Inc.

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

package concepts

import (
	"sort"

	"github.com/openshift-online/ocm-api-metamodel/pkg/names"
)

// Service is the representation of service, containing potentiall mutiple versions, for example the
// clusters management service.
type Service struct {
	documentedSupport
	annotatedSupport
	namedSupport

	// Model that owns this service:
	owner *Model

	// All the versions of the service, indexed by name:
	versions map[string]*Version
}

// NewService creates a new empty service.
func NewService() *Service {
	service := new(Service)
	service.versions = make(map[string]*Version)
	return service
}

// Owner returns the model that owns this service.
func (s *Service) Owner() *Model {
	return s.owner
}

// SetOwner sets the model that owns this service.
func (s *Service) SetOwner(value *Model) {
	s.owner = value
}

// Versions returns the list of versions of the service.
func (s *Service) Versions() VersionSlice {
	count := len(s.versions)
	versions := make(VersionSlice, count)
	index := 0
	for _, version := range s.versions {
		versions[index] = version
		index++
	}
	sort.Sort(versions)
	return versions
}

// FindVersion returns the version with the given name, or nil of there is no such version.
func (s *Service) FindVersion(name *names.Name) *Version {
	if name == nil {
		return nil
	}
	return s.versions[name.String()]
}

// AddVersion adds the given version to the service.
func (s *Service) AddVersion(version *Version) {
	if version != nil {
		s.versions[version.Name().String()] = version
		version.SetOwner(s)
	}
}

// AddVersions adds the given versions to the service.
func (s *Service) AddVersions(versions []*Version) {
	if len(versions) > 0 {
		for _, version := range versions {
			s.AddVersion(version)
		}
	}
}

// ServiceSlice is used to simplify sorting of slices of services by name.
type ServiceSlice []*Service

func (s ServiceSlice) Len() int {
	return len(s)
}

func (s ServiceSlice) Less(i, j int) bool {
	return names.Compare(s[i].name, s[j].name) == -1
}

func (s ServiceSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
