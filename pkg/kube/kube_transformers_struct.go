package kube

import (
	"bytes"

	netV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
)

// Transformers struct
type Transformers struct{}

var tx = &Transformers{}

// Int32ToPtr method
func (t *Transformers) Int32ToPtr(i int32) *int32 { return &i }

// Int64ToPtr method
func (t *Transformers) Int64ToPtr(i int64) *int64 { return &i }

// PathType creates a new ptr to a PathType enum
func (t *Transformers) PathType(v netV1.PathType) *netV1.PathType {
	return &v
}

// IntOrStringI method
func (t *Transformers) IntOrStringI(i int32) *intstr.IntOrString {
	return &intstr.IntOrString{Type: intstr.Int, IntVal: i}
}

// IntOrStringI method
func (t *Transformers) IntOrStringS(s string) intstr.IntOrString {
	return intstr.IntOrString{Type: intstr.String, StrVal: s}
}

func (t *Transformers) ResourcesToYAML(resources []runtime.Object) ([]byte, error) {
	s := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	yaml := []byte{}

	for _, v := range resources {
		buffer := new(bytes.Buffer)
		if err := s.Encode(v, buffer); err != nil {
			return yaml, err
		}
		yaml = append(yaml, []byte("---\n")...)
		yaml = append(yaml, buffer.Bytes()...)
	}
	return yaml, nil
}
