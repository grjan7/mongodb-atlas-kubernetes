package status

import (
	"context"
	"encoding/json"

	mdbv1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type patchValue struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

// PatchUpdateStatus performs the JSONPatch patch update to the Atlas Custom Resource.
// The "jsonPatch" merge allows to update only status field so is more
func PatchUpdateStatus(kubeClient client.Client, resource mdbv1.AtlasCustomResource) error {
	return doPatch(kubeClient, resource, resource.GetStatus())
}

func doPatch(kubeClient client.Client, resource runtime.Object, statusValue interface{}) error {
	payload := []patchValue{{
		Op:    "replace",
		Path:  "/status",
		Value: statusValue,
	}}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	patch := client.RawPatch(types.JSONPatchType, data)
	return kubeClient.Status().Patch(context.Background(), resource, patch)
}