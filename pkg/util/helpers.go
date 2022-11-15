package util

import (
	"bytes"

	netV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
)

func Int32ToPtr(i int32) *int32 {
	return &i
}

func IntOrStringI(i int32) *intstr.IntOrString {
	return &intstr.IntOrString{Type: intstr.Int, IntVal: i}
}

func IntOrStringS(s string) intstr.IntOrString {
	return intstr.IntOrString{Type: intstr.String, StrVal: s}
}

func PathType(v netV1.PathType) *netV1.PathType {
	return &v
}

// StringDefault - returns the first string with a non-empty value
func StringDefault(a ...string) string {
	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			return a[i]
		}
	}
	return ""
}

func ResourcesToYAML(resources []runtime.Object) ([]byte, error) {
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