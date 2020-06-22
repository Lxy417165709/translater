package syntex

import (
	"fmt"
	"lex/terminator"
	"strconv"
	"syntex/table"
)


// 存储变量
var ideToValue  = make(map[string]float64)


type TreeNode struct {
	symbol string
	pair *terminator.Pair
	son []*TreeNode
}

func NewTreeNode(symbol string) *TreeNode{
	return &TreeNode{symbol:symbol}
}

func (tn *TreeNode)FormSon(sentence *table.Sentence) {
	tn.son = nil
	symbols := sentence.GetSymbols()
	for i:=0;i<=len(symbols)-1;i++ {
		tn.son = append(tn.son,NewTreeNode(symbols[i]))
	}
}

func (tn *TreeNode) Show() {
	if tn==nil{
		return
	}
	for i:=0;i<len(tn.son);i++{
		if tn.son[i]==nil{
			continue
		}
		if tn.son[i].pair!=nil{
			fmt.Printf("id:%p(%s) -- .%s. --> id:%p(%s-%v)\n",
				tn,tn.symbol,"",
				tn.son[i],tn.son[i].symbol,tn.son[i].pair.GetValue(),
			)
		}else{
			fmt.Printf("id:%p(%s) -- .%s. --> id:%p(%s)\n",
				tn,tn.symbol,"",
				tn.son[i],tn.son[i].symbol,
			)
		}

		tn.son[i].Show()
	}
}


func (tn *TreeNode) Parse() {
	switch tn.symbol {
	case "BEGIN":
		tn.son[0].Parse()
	case "FUNC":


		if tn.son[0].symbol=="print" {
			fmt.Println(tn.son[1].ParseValue())
			return
		}
		if tn.son[0].symbol=="cube" {
			value := tn.son[1].ParseValue()
			fmt.Println(value *value *value)
			return
		}
		if tn.son[0].symbol=="PD" {
			if tn.son[0].son[1].ParseBool() {
				tn.son[0].son[3].Parse()
				return
			}
			if tn.son[0].son[5].symbol=="BLA"{
				return
			}

			tn.son[0].son[5].son[2].Parse()

			return
		}


		panic("未知函数")
	case "BLA","EXP":
	case "FZ":
		ide := tn.son[0].pair.GetValue().(string)
		value := tn.son[2].ParseValue()

		ideToValue[ide] = value
		fmt.Printf(
			"创建了变量 %s, 其值为 %v\n",
			tn.son[0].pair.GetValue().(string),
			value,
		)
	default:
		panic(fmt.Sprintf("%s %s","存在没有考虑的情况",tn.symbol))
	}

}
func (tn *TreeNode) ParseBool() bool{
	value1 := tn.son[0].ParseValue()
	value2 := tn.son[2].ParseValue()
	return judge(value1,value2,tn.son[1].son[0].symbol)
}


// 这里其实可以输出汇编代码
// 为了简单起见，这里输出执行结果吧
// 这里其实还需要前面解析出来的 变量表的支持，否则就无法知道变量的值是多少
// 为了测试，我先假定全是整数,没有括号啥的
func (tn *TreeNode) ParseValue() float64{
	switch tn.symbol {
	case "EXP":
		return tn.son[1].ParseOptAndValue(tn.son[0].ParseValue())
	case "EXPt":
		return tn.son[1].ParseOptAndValue(tn.son[0].ParseValue())
	case "EXPf":
		// 这个表示的是有括号的情况
		if len(tn.son)==3{
			return tn.son[1].ParseValue()
		}else{

			// 表示是变量
			if tn.son[0].pair.GetSymbol()=="IDE"{
				return ideToValue[tn.son[0].pair.GetValue().(string)]
			}

			value,err := strconv.ParseFloat(tn.son[0].pair.GetValue().(string),64)
			if err !=nil{
				panic(err)
			}
			// 表示字面量
			return value
		}
	}
	panic("存在没有考虑的东西")
}

func (tn *TreeNode) ParseOptAndValue(fatherValue float64) float64{
	// BLA情况
	if len(tn.son)==1{
		return fatherValue
	}
	leftValue := Do(fatherValue,tn.son[1].ParseValue(),tn.son[0].pair.GetValue().(string))
	return tn.son[2].ParseOptAndValue(leftValue)
}


func judge(leftValue float64,rightValue float64, tjOpt string) bool{
	switch tjOpt {
	case"<":
		return leftValue<rightValue
	case"<=":
		return leftValue<=rightValue
	case">=":
		return leftValue>=rightValue
	case"==":
		return leftValue==rightValue
	case">":
		return leftValue>rightValue
	default:
		return false
	}
}


func Do(leftValue float64,rightValue float64,opt string) float64{
	switch opt {
	case"+":
		return leftValue+rightValue
	case"-":
		return leftValue-rightValue
	case"*":
		return leftValue*rightValue
	case"/":	// 这里可以进行除0处理
		return leftValue/rightValue
	default:
		return leftValue
	}
}



