package registries

// // Module registries module
// func Module(app *gin.Engine, config *util.Config, k8s *kube.Service, logger *logrus.Entry) {
// 	log := logger.WithFields(logrus.Fields{"module": "registries"})

// 	registryService := &RegistryService{config: config, k8s: k8s, log: log}

// 	if err := registryService.OnModuleInit(); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Info("registries module loaded")
// }
