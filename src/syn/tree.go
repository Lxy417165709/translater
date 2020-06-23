package syn

import (
	"dto"
	"fmt"
	"syn/tb"
)

type TreeNode struct {
	token *dto.Token
	son   []*TreeNode
}

func NewTreeNode(token *dto.Token) *TreeNode {
	return &TreeNode{token: token}
}

func (tn *TreeNode) FormSon(sentence *tb.Sentence) {
	tn.son = nil
	symbols := sentence.Symbols
	for i := 0; i <= len(symbols)-1; i++ {
		tn.son = append(tn.son, NewTreeNode(&dto.Token{
			Symbol: symbols[i],
		}))
	}
}

// 存在特殊字符，导致mermaid无法正常显示
func (tn *TreeNode) Show() {
	if tn == nil {
		return
	}
	for i := 0; i < len(tn.son); i++ {
		if tn.son[i] == nil {
			continue
		}
		if tn.son[i].token != nil {
			fmt.Printf("id:%p( %s ) -- .%s. --> id:%p(( %s ))\n",
				tn, handleToSuitMermaid(tn.token.Symbol), "",
				tn.son[i], handleToSuitMermaid(stringfy(tn.son[i].token.Value)),
			)
		} else {
			fmt.Printf("id:%p( %s ) -- .%s. --> id:%p( %s )\n",
				tn, handleToSuitMermaid(tn.token.Symbol), "",
				tn.son[i], handleToSuitMermaid(tn.son[i].token.Symbol),
			)
		}

		tn.son[i].Show()
	}
}


func stringfy(itf interface{}) string{
	if itf==nil{
		return ""
	}
	return itf.(string)
}


// 冗余
func handleToSuitMermaid(str string) string {
	strToSuitMermaid := make(map[string]string)
	strToSuitMermaid["-"] = "减号"
	strToSuitMermaid[","] = "逗号"
	strToSuitMermaid["("] = "左括号"
	strToSuitMermaid[")"] = "右括号"
	strToSuitMermaid["["] = "左中括号"
	strToSuitMermaid["]"] = "右中括号"
	strToSuitMermaid["{"] = "左大括号"
	strToSuitMermaid["}"] = "右大括号"
	strToSuitMermaid[";"] = "分号"
	strToSuitMermaid[`"`] = "双引号"
	strToSuitMermaid[`.`] = "小点"
	if strToSuitMermaid[str] == "" {
		return str
	}
	return strToSuitMermaid[str]
}


