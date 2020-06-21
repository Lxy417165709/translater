package test

type TypeOfTest string
type TypeOfCreateFunction = func(string) Testable

const (
	NFATest    TypeOfTest = "NFA测试"
	SyntaxTest TypeOfTest = "语法分析测试"
)

var globalFactory *Factory

func init() {
	globalFactory = &Factory{
		make(map[TypeOfTest]TypeOfCreateFunction),
	}
	globalFactory.add(NFATest, NewNFATestItem)
	globalFactory.add(SyntaxTest, NewSyntaxAnalyzerTestItem)
}

type Factory struct {
	testTypeToCreateFunction map[TypeOfTest]TypeOfCreateFunction
}

func (f *Factory) NewFunctionIsExist(testType TypeOfTest) bool {
	return f.testTypeToCreateFunction[testType] != nil
}

func (f *Factory) GetCreateFunction(testType TypeOfTest) TypeOfCreateFunction {
	return f.testTypeToCreateFunction[testType]
}

func (f *Factory) add(testType TypeOfTest, createFunction TypeOfCreateFunction) {
	f.testTypeToCreateFunction[testType] = createFunction
}
