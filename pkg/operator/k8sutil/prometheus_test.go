/*
Copyright 2020 The Rook Authors. All rights reserved.

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

// Package k8sutil for Kubernetes helpers.
package k8sutil

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rook/rook/pkg/util"
)

func TestGetServiceMonitor(t *testing.T) {
	projectRoot := util.PathToProjectRoot()
	filePath := path.Join(projectRoot, "/cluster/examples/kubernetes/ceph/monitoring/service-monitor.yaml")
	servicemonitor, err := GetServiceMonitor(filePath)
	assert.Nil(t, err)
	assert.Equal(t, "rook-ceph-mgr", servicemonitor.GetName())
	assert.Equal(t, "rook-ceph", servicemonitor.GetNamespace())
	assert.NotNil(t, servicemonitor.Spec.NamespaceSelector.MatchNames)
	assert.NotNil(t, servicemonitor.Spec.Endpoints)
	assert.Len(t, servicemonitor.Spec.Endpoints, 1)
	assert.NotNil(t, servicemonitor.Spec.Endpoints[0])
	assert.NotNil(t, servicemonitor.Spec.Endpoints[0].RelabelConfigs)
	assert.Len(t, servicemonitor.Spec.Endpoints[0].RelabelConfigs, 1)
	assert.Equal(t, []string{"rook_cluster"}, servicemonitor.Spec.Endpoints[0].RelabelConfigs[0].SourceLabels)
	assert.Equal(t, "replace", servicemonitor.Spec.Endpoints[0].RelabelConfigs[0].Action)
	assert.Equal(t, "cluster", servicemonitor.Spec.Endpoints[0].RelabelConfigs[0].TargetLabel)
}

func TestGetPrometheusRule(t *testing.T) {
	projectRoot := util.PathToProjectRoot()
	filePath := path.Join(projectRoot, "/cluster/examples/kubernetes/ceph/monitoring/prometheus-ceph-v14-rules.yaml")
	rules, err := GetPrometheusRule(filePath)
	assert.Nil(t, err)
	assert.Equal(t, "prometheus-ceph-rules", rules.GetName())
	assert.Equal(t, "rook-ceph", rules.GetNamespace())
	// Labels should be present as they are used by prometheus for identifying rules
	assert.NotNil(t, rules.GetLabels())
	assert.NotNil(t, rules.Spec.Groups)
}
