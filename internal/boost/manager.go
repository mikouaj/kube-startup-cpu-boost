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
	"context"
	"errors"
	"sync"

	corev1 "k8s.io/api/core/v1"

	autoscaling "github.com/mikouaj/kube-startup-cpu-boost/api/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
)

var (
	errStartupCPUBoostAlreadyExists = errors.New("startupCPUBoost already exists")
	errInvalidStartupCPUBoostSpec   = errors.New("invalid startupCPUBoost spec")
)

type Manager interface {
	AddStartupCPUBoost(ctx context.Context, boost *autoscaling.StartupCPUBoost) error
	DeleteStartupCPUBoost(boost *autoscaling.StartupCPUBoost)
	StartupCPUBoostForPod(pod *corev1.Pod) (string, bool)
}

type managerImpl struct {
	sync.RWMutex
	startupCPUBoosts map[string]map[string]*startupCPUBoost
}

func NewManager() Manager {
	return &managerImpl{
		startupCPUBoosts: make(map[string]map[string]*startupCPUBoost),
	}
}

func (m *managerImpl) AddStartupCPUBoost(ctx context.Context, boost *autoscaling.StartupCPUBoost) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.getStartupCPUBoostWithNS(boost.Namespace, boost.Name); ok {
		return errStartupCPUBoostAlreadyExists
	}
	boostImpl, err := newStartupCPUBoost(boost)
	if err != nil {
		return errInvalidStartupCPUBoostSpec
	}
	m.addStartupCPUBoostWithNS(boostImpl)
	return nil
}

func (m *managerImpl) DeleteStartupCPUBoost(boost *autoscaling.StartupCPUBoost) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.getStartupCPUBoostWithNS(boost.Namespace, boost.Name); !ok {
		return
	}
	m.deleteStartupCPUBoostWithNS(boost.Namespace, boost.Name)
}

func (m *managerImpl) StartupCPUBoostForPod(pod *corev1.Pod) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	nsBoosts, ok := m.startupCPUBoosts[pod.Namespace]
	if !ok {
		return "", false
	}
	for _, boost := range nsBoosts {
		if boost.selector.Matches(labels.Set(pod.Labels)) {
			return boost.name, true
		}
	}
	return "", false
}

func (m *managerImpl) addStartupCPUBoostWithNS(boostImpl *startupCPUBoost) {
	nsBoosts, ok := m.startupCPUBoosts[boostImpl.namespace]
	if !ok {
		nsBoosts = make(map[string]*startupCPUBoost)
		m.startupCPUBoosts[boostImpl.namespace] = nsBoosts
	}
	nsBoosts[boostImpl.name] = boostImpl
}

func (m *managerImpl) getStartupCPUBoostWithNS(ns string, name string) (*startupCPUBoost, bool) {
	if nsboosts, ok := m.startupCPUBoosts[ns]; ok {
		return nsboosts[name], true
	}
	return nil, false
}

func (m *managerImpl) deleteStartupCPUBoostWithNS(ns string, name string) {
	if nsBoosts, ok := m.startupCPUBoosts[ns]; ok {
		delete(nsBoosts, name)
	}
}
