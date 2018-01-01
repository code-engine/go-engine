package logger

func NewFakeOutput() *FakeOutput {
	return &FakeOutput{}
}

type FakeOutput struct {
	Data interface{}
}

func (f *FakeOutput) Print(message ...interface{}) {
	f.Data = message[0]
}
