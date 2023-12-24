package internal

type Service interface {
	DoBusiness()
}

type Builder interface {
	reset()
	setName(name string)
	buildLogger(logger Logger)
	buildNotifier(notifier Notifier)
	buildDataLayer(dataLayer DataLayer)
	buildUploader(uploader Uploader)
	result() Service
}

/*
In normal way every builder has this function ->
There for can have many builder by implement builder interface
*/

type serviceBuilder struct {
	service *complexService
}

func NewBuilder() *serviceBuilder {
	return &serviceBuilder{}
}

func (s *serviceBuilder) reset() {
	s.service = &complexService{}
}

func (s *serviceBuilder) setName(name string) {
	s.service.setName(name)
}

func (s *serviceBuilder) buildLogger(logger Logger) {
	s.service.setLogger(logger)
}

func (s *serviceBuilder) buildNotifier(notifier Notifier) {
	s.service.setNotifier(notifier)
}

func (s *serviceBuilder) buildDataLayer(dataLayer DataLayer) {
	s.service.setDataLayer(dataLayer)
}

func (s *serviceBuilder) buildUploader(uploader Uploader) {
	s.service.setUploader(uploader)
}

func (s *serviceBuilder) result() Service {
	return s.service
}
