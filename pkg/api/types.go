/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package api

import (
	"github.com/blackducksoftware/horizon/pkg/components"
	routev1 "github.com/openshift/api/route/v1"
)

// Route defines the route component
type Route struct {
	Namespace          string
	Name               string
	Kind               string
	ServiceName        string
	PortName           string
	Labels             map[string]string
	TLSTerminationType routev1.TLSTerminationType
}

// ComponentList defines the list of components for an app
type ComponentList struct {
	ReplicationControllers    []*components.ReplicationController
	Services                  []*components.Service
	ConfigMaps                []*components.ConfigMap
	ServiceAccounts           []*components.ServiceAccount
	ClusterRoleBindings       []*components.ClusterRoleBinding
	ClusterRoles              []*components.ClusterRole
	RoleBindings              []*components.RoleBinding
	Roles                     []*components.Role
	Deployments               []*components.Deployment
	Secrets                   []*components.Secret
	PersistentVolumeClaims    []*components.PersistentVolumeClaim
	Routes                    []*Route
	CustomResourceDefinitions []*components.CustomResourceDefinition
}

// GetKubeInterfaces returns a list of kube components as interfaces
func (clist *ComponentList) GetKubeInterfaces() []interface{} {
	components := []interface{}{}
	for _, rc := range clist.ReplicationControllers {
		components = append(components, rc.ReplicationController)
	}
	for _, svc := range clist.Services {
		components = append(components, svc.Service)
	}
	for _, cm := range clist.ConfigMaps {
		components = append(components, cm.ConfigMap)
	}
	for _, sa := range clist.ServiceAccounts {
		components = append(components, sa.ServiceAccount)
	}
	for _, crb := range clist.ClusterRoleBindings {
		components = append(components, crb.ClusterRoleBinding)
	}
	for _, cr := range clist.ClusterRoles {
		components = append(components, cr.ClusterRole)
	}
	for _, d := range clist.Deployments {
		components = append(components, d.Deployment)
	}
	for _, sec := range clist.Secrets {
		components = append(components, sec.Secret)
	}
	for _, pvc := range clist.PersistentVolumeClaims {
		components = append(components, pvc.PersistentVolumeClaim)
	}
	return components
}

// RegistryConfiguration contains the registry configuration
type RegistryConfiguration struct {
	Registry    string   `json:"registry"`
	PullSecrets []string `json:"pullSecrets"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfiguration) DeepCopyInto(out *RegistryConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlackduckSpec.
func (in *RegistryConfiguration) DeepCopy() *RegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(RegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// SecurityContext will contain the specifications of a security contexts
type SecurityContext struct {
	FsGroup    *int64 `json:"fsGroup"`
	RunAsUser  *int64 `json:"runAsUser"`
	RunAsGroup *int64 `json:"runAsGroup"`
}

// DeploymentResource contains the specification of a pod replica's, container memory and cpu limits and requests and container JVM heap max memory
type DeploymentResource struct {
	Replicas      *int32    `json:"replicas,omitempty"`
	Resources     Resources `json:"resources,omitempty"`
	HeapMaxMemory *string   `json:"heapMaxMemory,omitempty"`
}

// Resources contains specification for container requests and limits
type Resources struct {
	Requests Resource `json:"requests,omitempty"`
	Limits   Resource `json:"limits,omitempty"`
}

// Resource contains specification for container cpu and memory
type Resource struct {
	CPU    *string `json:"cpu,omitempty"`
	Memory *string `json:"memory,omitempty"`
}
