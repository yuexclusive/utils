// package web

// import (

// 	// "github.com/micro/go-micro/v2/registry"
// 	// "github.com/micro/go-micro/v2/registry/etcd"

// 	"github.com/gin-gonic/gin"
// 	// "github.com/micro/go-micro/v2/metadata"
// )

// type Starter interface {
// 	Start(engine *gin.Engine)
// }

// type Options struct {
// 	IsOpenSwagger bool
// 	// 是否记录日志到ES 默认为false
// 	IsLogToES bool
// 	// 是否使用opentrace(jaeger)
// 	IsTrace bool
// 	//是否允许跨域 默认为false,因为micro api默认做了跨域处理
// 	IsAllowOrigin bool
// 	//是否限流 默认为true
// 	IsRateLimite bool
// 	//端口 默认为空
// 	Port string
// 	//是否监控
// 	IsMonitor bool
// }

// type Option func(ops *Options)

// // 	router := gin.Default()

// // 	if options.IsAllowOrigin {
// // 		router.Use(
// // 			AllowOrigin(),
// // 		)
// // 		logrus.Infoln("开启跨域")
// // 	}

// // 	if options.IsRateLimite {
// // 		router.Use(
// // 			RateLimite(),
// // 		)
// // 		logrus.Infoln("开启限流")
// // 	}

// // 	if options.IsTrace {
// // 		_, closer, err := trace.NewTracer(name, appconfig.MustGet().JaegerAddress)

// // 		if err != nil {
// // 			logrus.Fatal(err)
// // 			return nil
// // 		}
// // 		defer closer.Close()

// // 		router.Use(TracerWrapper)
// // 		logrus.Infoln("开启链路追踪")
// // 	}

// // 	head := strings.TrimPrefix(name, "go.micro.api.")

// // 	if options.IsOpenSwagger {
// // 		var swaggerPath, swaggerURL string
// // 		swaggerPath = fmt.Sprintf("/%s/swagger/*any", head)
// // 		swaggerURL = fmt.Sprintf("http://%s:%s/%s/swagger/doc.json", config.HostIP, config.APIPort, head)

// // 		UseSwagger(swaggerPath, swaggerURL, router)
// // 	}

// // 	if options.IsMonitor {
// // 		router.Use(
// // 			Prometheus(head),
// // 		)
// // 	}

// // 	starter.Start(router)

// // 	service.Handle("/", router)

// // 	// run service
// // 	if err := service.Run(); err != nil {
// // 		return fmt.Errorf("服务运行错误:%w", err)
// // 	}

// // 	return nil
// // }

// // TracerWrapper tracer 中间件
