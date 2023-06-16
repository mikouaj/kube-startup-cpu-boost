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

package webhook

import (
	"context"
	"net/http"

	"github.com/mikouaj/kube-startup-cpu-boost/internal/boost"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/mutate-v1-pod,mutating=true,failurePolicy=fail,sideEffects=None,groups="",resources=pods,verbs=create,versions=v1,name=cpuboost.autoscaling.stefaniak.dev,admissionReviewVersions=v1

type podCPUBoostHandler struct {
	decoder *admission.Decoder
	manager boost.Manager
}

func NewPodCPUBoostWebHook(mgr boost.Manager) *webhook.Admission {
	return &webhook.Admission{
		Handler: &podCPUBoostHandler{
			manager: mgr,
		},
	}
}

func (h *podCPUBoostHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := corev1.Pod{}
	err := h.decoder.Decode(req, &pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	return admission.Allowed("good to go")
	/*
		  // mutate the fields in pod

		  marshaledPod, err := json.Marshal(pod)
		  if err != nil {
			  return admission.Errored(http.StatusInternalServerError, err)
		  }
		  return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
		}

		func (h *cpuBoostWebHook) InjectDecoder(d *admission.Decoder) error {
			h.decoder = d
			return nil
	*/
}
