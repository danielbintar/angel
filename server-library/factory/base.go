package factory

func MockAsyncPublisher(options ...string) DummyAsyncPublisher {
	return DummyAsyncPublisher{}
}

type DummyAsyncPublisher struct {
	Options          []string
	PublishCallCount uint
}

func (self *DummyAsyncPublisher) Publish(_ string, _ string) {
	self.PublishCallCount++
	return
}

func (self *DummyAsyncPublisher) Close() {
	return
}

