package registries

// // type RegistryServiceType interface {
// // 	GetContainerRegistry() (namespaces *coreV1.Secret, err error)
// // }

// // RegistryService struct
// type RegistryService struct {
// 	config           *config.ConfigService
// 	k8s              *kube.Service
// 	log              *logrus.Entry
// 	onNamespaceEvent rxgo.Observable
// 	onNamespaceAdded rxgo.Observable
// }

// // PrivateRegistry struct

// // ListPrivateRegisties method
// func (t *RegistryService) ListPrivateRegisties() (registries *[]kube.PrivateRegistryDTO, err error) {
// 	// registries, err := t.k8s.ListSecrets(t.config.SERVICE_NAME)
// 	// metadata := []PrivateRegistryDTO{}

// 	// for _, r := range registries.Items {
// 	// 	if r.ObjectMeta.Labels["managed-by"] == t.config.SERVICE_NAME {
// 	// 		metadata = append(metadata, PrivateRegistryDTO{
// 	// 			Name:      r.ObjectMeta.Name,
// 	// 			Namespace: r.ObjectMeta.Namespace,
// 	// 			Labels:    r.ObjectMeta.Labels,
// 	// 		})
// 	// 	}
// 	// }

// 	return registries, err
// }

// // CopyPrivateRegistry method
// func (t *RegistryService) CopyPrivateRegistry(stream rxgo.Observable) {
// 	stream.DoOnError(func(err error) {
// 		t.log.Error(err)
// 	})

// 	for item := range stream.Observe() {
// 		namespace := item.V.(*coreV1.Namespace)
// 		if t.config.PRIVATE_REGISTRY_ENABLED == "true" && !strings.HasPrefix(namespace.Name, "kube-") {
// 			t.log.Info("cloning private-registry to: " + namespace.Name)

// 			if _, err := t.k8s.CopyRegistry("private-registry", t.config.SERVICE_NAME, namespace.Name); err != nil {
// 				t.log.Error(err)
// 			}
// 		}
// 	}
// }

// // OnModuleInit function description
// func (t *RegistryService) OnModuleInit() error {
// 	rx := kube.Rx{}

// 	if t.config.PRIVATE_REGISTRY_ENABLED == "false" {
// 		return nil
// 	}

// 	_, err := t.k8s.CreateRegistrySecret(&kube.CreateRegistryDTO{
// 		Name:      "private-registry",
// 		Namespace: t.config.SERVICE_NAME,
// 		Registry:  t.config.PRIVATE_REGISTRY_URL,
// 		Username:  t.config.PRIVATE_REGISTRY_USER,
// 		Password:  t.config.PRIVATE_REGISTRY_PASS,
// 	})

// 	// listen to namespace events
// 	t.onNamespaceEvent = t.k8s.OnNamespaceEvent(metaV1.ListOptions{
// 		TypeMeta:            metaV1.TypeMeta{Kind: "Namespace"},
// 		Watch:               true,
// 		AllowWatchBookmarks: false,
// 	})

// 	// filter on namespace added events
// 	t.onNamespaceAdded = t.onNamespaceEvent.
// 		Filter(rx.EventTypeFilter(watch.Added)).
// 		Map(rx.NamespaceMap())

// 	go t.CopyPrivateRegistry(t.onNamespaceAdded)

// 	return err
// }
