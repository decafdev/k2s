package api

// // Middleware logs a gin HTTP request in JSON format, with some additional custom key/values
// func Middleware(config *config.ConfigService) gin.HandlerFunc {
// 	log := NewLogger(config)

// 	return func(context *gin.Context) {
// 		template := "request [status:%v] [method:%s] to [path:%s] in [ms:%v]"
// 		// Start timer
// 		start := time.Now()

// 		// Process Request
// 		context.Next()

// 		// Stop timer
// 		duration := time.Since(start)

// 		with := logrus.Fields{
// 			"method":   context.Request.Method,
// 			"path":     context.Request.RequestURI,
// 			"status":   context.Writer.Status(),
// 			"referrer": context.Request.Referer(),
// 			"duration": duration.Milliseconds(),
// 		}

// 		msg := fmt.Sprintf(template, with["status"], with["method"], with["path"], with["duration"])

// 		if context.Writer.Status() >= 500 {
// 			log.WithFields(with).Error(fmt.Sprintf("[errors: %s] %s", context.Errors.String(), msg))
// 		} else {
// 			log.Info(msg)
// 		}
// 	}
// }
