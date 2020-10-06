// Copyright (c) 2018-2020 Splunk Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spltest "github.com/splunk/splunk-operator/pkg/splunk/test"
)

func TestApplyConfigMap(t *testing.T) {
	funcCalls := []spltest.MockFuncCall{{MetaName: "*v1.ConfigMap-test-defaults"}}
	createCalls := map[string][]spltest.MockFuncCall{"Get": funcCalls, "Create": funcCalls}
	updateCalls := map[string][]spltest.MockFuncCall{"Get": funcCalls, "Update": funcCalls}
	current := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "defaults",
			Namespace: "test",
		},
	}
	revised := current.DeepCopy()
	revised.Data = map[string]string{"a": "b"}
	reconcile := func(c *spltest.MockClient, cr interface{}) error {
		_, err := ApplyConfigMap(c, cr.(*corev1.ConfigMap))
		return err
	}
	spltest.ReconcileTester(t, "TestApplyConfigMap", &current, revised, createCalls, updateCalls, reconcile, false)
}
