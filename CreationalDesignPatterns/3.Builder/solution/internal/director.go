package internal

type Director interface {
	BuildService(builder Builder) Service
}

type serviceBuilderDirector struct {
}

func (s serviceBuilderDirector) BuildService(builder Builder) Service {
	builder.reset()
	builder.setName("Complex Service")
	builder.buildLogger(StdLogger{})
	builder.buildNotifier(EmailNotifier{})
	builder.buildDataLayer(MySQLDataLayer{})
	builder.buildUploader(AWSS3Uploader{})

	return builder.result()
}

func NewDirector() serviceBuilderDirector {
	return serviceBuilderDirector{}
}
