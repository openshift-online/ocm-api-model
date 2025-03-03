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

// Model is the representation of the set of services, each with a set of versions.
type Model struct {
	// All the services of the model, indexed by name:
	services map[string]*Service
}

// NewModel creates a new model containing only the built-in types.
func NewModel() *Model {
	model := new(Model)
	model.services = make(map[string]*Service)
	return model
}

// Services returns the list of services that are part of this model.
func (m *Model) Services() ServiceSlice {
	count := len(m.services)
	services := make(ServiceSlice, count)
	index := 0
	for _, service := range m.services {
		services[index] = service
		index++
	}
	sort.Sort(services)
	return services
}

// FindService returns the service with the given name, or nil of there is no such service.
func (m *Model) FindService(name *names.Name) *Service {
	if name == nil {
		return nil
	}
	return m.services[name.String()]
}

// AddService adds the given service to the model.
func (m *Model) AddService(service *Service) {
	if service != nil {
		m.services[service.Name().String()] = service
		service.SetOwner(m)
	}
}

// AddServices adds the given services to the model.
func (m *Model) AddServices(services []*Service) {
	for _, service := range services {
		m.AddService(service)
	}
}
