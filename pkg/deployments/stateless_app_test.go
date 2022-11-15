package deployments

import (
	"testing"
)

func TestNewAPIApplication(t *testing.T) {
	// s := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	// given struct
	type given struct {
		input  *APIOptions
		expect APIResources
		err    string
	}

	scenario := make(map[string]given)

	scenario["when all required parameters are provided"] = given{
		input: &APIOptions{
			Name:        "whoami",
			Image:       "traefik/whoami:v1.7.1",
			Version:     "1.7.1",
			Port:        80,
			Replicas:    1,
			MemoryLimit: 64,
			CPULimit:    250,
			Middlewares: []string{"rewrite-url@file"},
		},
	}

	// scenario["when no parameters are provided"] = given{
	// 	input: &kube.APIOptions{},
	// 	err: strings.Join([]string{
	// 		"Key: 'APIOptions.Name' Error:Field validation for 'Name' failed on the 'required' tag",
	// 		"Key: 'APIOptions.Image' Error:Field validation for 'Image' failed on the 'required' tag",
	// 	}, "\n"),
	// }

	// for when, given := range scenario {
	// 	t.Run(when, func(t *testing.T) {
	// 		r, err := NewAPIApplication(given.input)

	// 		if (err != nil) && (given.err != err.Error()) {
	// 			t.Log(err.Error())
	// 			t.Errorf("unexpected error with [%s]", err.Error())
	// 			return
	// 		}
	// 		yaml, err := r.ToYAML()
	// 		if err != nil {
	// 			t.Error(err)
	// 		}
	// 		// os.WriteFile("./__snapshots__/kube_api-application_test.yaml", yaml, 0644)
	// 		snap, err := os.ReadFile("./__snapshots__/kube_api-application_test.yaml")
	// 		if err != nil {
	// 			t.Error(err)
	// 		}
	// 		assert.YAMLEq(t, string(snap), string(yaml))
	// 	})
	// }
}
