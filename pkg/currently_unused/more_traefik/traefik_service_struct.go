package more_traefik

// var test = []byte(`
// http:
//   middlewares:
//     forward-auth:
//       forwardAuth:
//         address: https://someplace.com
//         trustForwardHeader: true
//         authRequestHeaders:
//         - ""
//         authResponseHeaders:
//         - ""
//     rate-limit-100:
//       rateLimit:
//         sourceCriterion:
//           requestHeaderName: authorization
//         average: 100
//         burst: 100
//     rewrite-url:
//       stripPrefixRegex:
//         regex:
//         - ""
// `)

// // TraefikService struct
// type TraefikService struct {
// 	config *config.ConfigService
// 	k8s    *kube.Service
// 	log    *logrus.Entry
// }

// // GetTraefikConfig method
// func (t *TraefikService) GetTraefikConfig() (map[string]string, error) {
// 	config, err := t.k8s.GetConfigMap("k2s-traefik-options", t.config.SERVICE_NAME)
// 	if err != nil {
// 		return map[string]string{}, err
// 	}
// 	return config.Data, err
// }

// // OnModuleInit method
// func (t *TraefikService) OnModuleInit() error {
// 	Replicas, err := strconv.Atoi(t.config.TRAEFIK_REPLICAS)
// 	if err != nil {
// 		t.log.Fatal(err)
// 	}

// 	options := &ResourceOptions{
// 		Name:      "traefik",
// 		Namespace: t.config.SERVICE_NAME,
// 		Version:   t.config.TRAEFIK_VERSION,
// 		Replicas:  int32(Replicas),
// 	}

// 	if _, err := options.Validate(); err != nil {
// 		return err
// 	}

// 	traefik, err := NewTraefikResources(options, &TraefikConfig{})

// 	t.log.Info("deploying traefik.service-account")
// 	if _, err = t.k8s.CreateServiceAccount(options.Namespace, traefik.ServiceAccount); !apierrors.IsAlreadyExists(err) {
// 		return err
// 	}

// 	t.log.Info("deploying traefik.cluster-role")
// 	if _, err = t.k8s.ApplyClusterRole(traefik.ClusterRole); err != nil {
// 		return err
// 	}

// 	t.log.Info("deploying traefik.cluster-role-binding")
// 	if _, err = t.k8s.ApplyClusterRoleBinding(traefik.ClusterRoleBinding); err != nil {
// 		return err
// 	}

// 	t.log.Info("deploying traefik.service")
// 	if _, err = t.k8s.ApplyService(options.Namespace, traefik.Service); err != nil {
// 		return err
// 	}

// 	t.log.Info("deploying traefik.configmap")
// 	if _, err = t.k8s.ApplyConfigMap(options.Namespace, traefik.ConfigMap); err != nil {
// 		return err
// 	}

// 	t.log.Info("deploying traefik.deployment")
// 	if _, err = t.k8s.ApplyDeployment(options.Namespace, traefik.Deployment); err != nil {
// 		return err
// 	}

// 	return nil
// }
