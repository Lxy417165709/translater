package syntex

import (
	"fmt"
	"lex/terminator"
	"syntex/table"
)

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

// 这里其实可以输出汇编代码
// 为了简单起见，这里输出执行结果吧
// 这里其实还需要前面解析出来的 变量表的支持，否则就无法知道变量的值是多少
// 为了测试，我先假定全是整数,没有括号啥的
func (tn *TreeNode) ParseValue() int{
	switch tn.symbol {
	case "EXP":
		return tn.son[1].ParseOptAndValue(tn.son[0].ParseValue())
	case "EXPt":
		return tn.son[1].ParseOptAndValue(tn.son[0].ParseValue())
	case "EXPf":
		value,_ := tn.son[0].pair.GetValue().(int)
		return value
	}
	panic("存在没有考虑的东西")
}

func (tn *TreeNode) ParseOptAndValue(fatherValue int) int{
	// BLA情况
	if len(tn.son)==1{
		return fatherValue
	}
	leftValue := Do(fatherValue,tn.son[1].ParseValue(),tn.son[0].pair.GetValue().(string))
	return tn.son[2].ParseOptAndValue(leftValue)
}


func Do(leftValue int,rightValue int,opt string) int{
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



