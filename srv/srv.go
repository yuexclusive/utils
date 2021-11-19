package srv

// "github.com/micro/go-micro/v2"
// "github.com/micro/go-micro/v2/broker"
// "github.com/micro/go-micro/v2/broker/nats"
// "github.com/micro/go-micro/v2/registry"
// "github.com/micro/go-micro/v2/registry/etcd"

type Options struct {
	// 是否记录日志到ES 默认为false
	IsLogToES bool
	// 是否使用opentrace(jaeger)
	IsTrace bool
	//是否监控
	IsMonitor bool
}

type Option func(ops *Options)

// type Starter interface {
// 	Start(s micro.Service)
// }

// func Startup(name string, starter Starter, opts ...Option) error {
// options := &Options{
// 	IsLogToES: false,
// 	IsTrace:   false,
// }

// for _, opt := range opts {
// 	opt(options)
// }

// if options.IsLogToES {
// 	log.LogToES(name)
// }

// microOpts := []micro.Option{
// 	micro.Name(name),
// 	micro.Version("latest"),
// 	micro.Registry(etcd.NewRegistry(registry.Addrs(appconfig.MustGet().ETCDAddress))),
// 	micro.RegisterTTL(time.Second * 30),
// 	micro.RegisterInterval(time.Second * 15),
// 	micro.Broker(nats.NewBroker(broker.Addrs(appconfig.MustGet().NatsAddress))),
// }

// if options.IsTrace {
// 	t, closer, err := trace.NewTracer(name, appconfig.MustGet().JaegerAddress)

// 	if err != nil {
// 		return fmt.Errorf("trace error: %w", err)
// 	}
// 	defer closer.Close()
// 	microOpts = append(microOpts, micro.WrapHandler(ocplugin.NewHandlerWrapper(t)))
// 	logrus.Infoln("open trace")
// }

// 	if options.IsMonitor {
// 		microOpts = append(microOpts, micro.WrapHandler(prometheus.NewHandlerWrapper()))
// 		logrus.Infoln("open monitor")
// 	}

// 	// New Service
// 	service := micro.NewService(microOpts...)

// 	// Initialise service
// 	service.Init()

// 	starter.Start(service)

// 	// Run service
// 	if err := service.Run(); err != nil {
// 		return fmt.Errorf("micro service run error: %w", err)
// 	}
// 	return nil
// }
