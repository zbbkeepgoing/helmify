package crossplane

import (
	"testing"

	"github.com/arttor/helmify/internal"
	"github.com/arttor/helmify/pkg/metadata"
	"github.com/stretchr/testify/assert"
)

const (
	strXRD = `apiVersion: apiextensions.crossplane.io/v1
kind: CompositeResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.7.0
  creationTimestamp: null
  name: cephvolumes.test.example.com
spec:
  group: test.example.com
  names:
    kind: CephVolume
    listKind: CephVolumeList
    plural: cephvolumes
    singular: cephvolume
  scope: Namespaced
`
)

func Test_xrd_Process(t *testing.T) {
	var testInstance xrd

	t.Run("processed", func(t *testing.T) {
		obj := internal.GenerateObj(strXRD)
		processed, _, err := testInstance.Process(&metadata.Service{}, obj)
		assert.NoError(t, err)
		assert.Equal(t, true, processed)
	})
	t.Run("skipped", func(t *testing.T) {
		obj := internal.TestNs
		processed, _, err := testInstance.Process(&metadata.Service{}, obj)
		assert.NoError(t, err)
		assert.Equal(t, false, processed)
	})
}
