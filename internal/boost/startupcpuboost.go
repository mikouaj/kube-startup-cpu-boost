// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package boost

import (
	autoscaling "github.com/mikouaj/kube-startup-cpu-boost/api/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type startupCPUBoost struct {
	name      string
	namespace string
	selector  labels.Selector
}

func newStartupCPUBoost(boost *autoscaling.StartupCPUBoost) (*startupCPUBoost, error) {
	selector, err := metav1.LabelSelectorAsSelector(&boost.Selector)
	if err != nil {
		return nil, err
	}
	return &startupCPUBoost{
		name:      boost.Name,
		namespace: boost.Namespace,
		selector:  selector,
	}, nil
}

func Key(boost *autoscaling.StartupCPUBoost) string {
	return boost.Namespace + "/" + boost.Name
}
