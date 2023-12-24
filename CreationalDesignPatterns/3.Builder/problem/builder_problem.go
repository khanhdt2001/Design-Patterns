package problem

// set of service in struct
type Logger interface {
	Log(...any)
}

type Notifier interface {
	Send(msg string)
}

type DataLayer interface {
	Save()
}

type Uploader interface {
	Upload()
}

type StdLogger struct{}
type FileLogger struct{}

func (l StdLogger) Log(...any)  {}
func (l FileLogger) Log(...any) {}

type SMSNotifier struct{}
type EmailNotifier struct{}

func (s SMSNotifier) Send(msg string)   {}
func (s EmailNotifier) Send(msg string) {}

type NoSqlDataLayer struct{}
type SqlDataLayer struct{}

func (s SqlDataLayer) Save()   {}
func (s NoSqlDataLayer) Save() {}

type AWSUploader struct{}
type DriveUploader struct{}

func (u AWSUploader) Upload()   {}
func (u DriveUploader) Upload() {}

// struct using all complex service
type Services struct {
	name      string
	log       Logger
	notifier  Notifier
	dataLayer DataLayer
	uploader  Uploader
}

func (s Services) setLogger(l Logger) {
	s.log = l
}
func (s Services) setNotifier(n Notifier) {
	s.notifier = n
}
func (s Services) setDataLayer(dl DataLayer) {
	s.dataLayer = dl
}
func (s Services) setUploader(u Uploader) {
	s.uploader = u
}

func (s Services) doSomeThing() {
	s.notifier.Send("msg")
	s.log.Log("builder")
	s.dataLayer.Save()
	s.uploader.Upload()
}

// way 1 to create service
func NewService(name string, l Logger, n Notifier, dl DataLayer, u Uploader) Services {
	return Services{
		name:      name,
		log:       l,
		notifier:  n,
		dataLayer: dl,
		uploader:  u,
	}
}

// some other way to create service
func NewServiceWithName(name string) Services {
	return Services{
		name:      name,
		log:       StdLogger{},
		notifier:  EmailNotifier{},
		dataLayer: SqlDataLayer{},
		uploader:  AWSUploader{},
	}
}

/*
This way of coding is hard to maintain and difficult to reuse
this code is not solid || O: open for extend close for modify
*/