package more_kube

// func TestKubeContainerRegistry(t *testing.T) {
// 	// given struct
// 	type given struct {
// 		input *kube.ContainerRegistryOptions
// 		err   string
// 	}

// 	scenario := make(map[string]given)

// 	scenario["when all required parameters are provided"] = given{
// 		input: &kube.ContainerRegistryOptions{
// 			Name:      "go-test",
// 			Namespace: "go-test",
// 			Registry:  "https://someurl.com",
// 			Username:  "test-user",
// 			Password:  "test-password",
// 		},
// 	}

// 	for when, given := range scenario {
// 		t.Run(when, func(t *testing.T) {
// 			resource, err := kube.NewContainerRegistry(given.input)
// 			if (err != nil) && (given.err != err.Error()) {
// 				t.Log(err.Error())
// 				t.Errorf("unexpected error with [%s]", err.Error())
// 				return
// 			}

// 			yaml, err := resource.ToYAML()
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			snaps.MatchSnapshot(t, string(yaml))
// 			// if !reflect.DeepEqual(*resource, given.expect) {
// 			// 	t.Errorf("%s got %v, but expected %v", when, *resource, given.expect)
// 			// }
// 		})
// 	}
// }
