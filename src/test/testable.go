package test

type Testable interface {
	Test() bool
	GetErrMsg() string
}
