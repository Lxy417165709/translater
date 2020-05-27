package test

type Testable interface {
	Test() bool
	GetErrMsg() string
}
type TestableDir interface {
	Testable
	GetTestType() TestName
}
