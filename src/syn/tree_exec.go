package syn

import (
	"fmt"
	"strconv"
)


func (tn *TreeNode) Exec() error {
	switch tn.token.Symbol {
	case "BEGIN":
		return tn.son[0].Exec()
	case "FUNC":
		if tn.son[0].token.Symbol == "print" {
			fmt.Printf("%0.6f\n", tn.son[1].GetNumber())
			return nil
		}
	case "PD":
		if tn.son[1].GetBool() {
			return tn.son[3].Exec()
		}
		return tn.son[5].Exec()
	case "PD'":
		if tn.son[0].token.Symbol == "BLA" {
			return nil
		}
		return tn.son[2].Exec()
	case "FZ":
		ideName := tn.son[0].token.Value.(string)
		value := tn.son[2].GetNumber()
		ideStorage[ideName] = value
		fmt.Printf(
			"创建了变量 %s, 其值为 %0.6f\n",
			ideName,
			value,
		)
	case "BLA":
	}

	return nil
}

func (tn *TreeNode) GetBool() bool {
	value1 := tn.son[0].GetNumber()
	value2 := tn.son[2].GetNumber()
	return judge(value1, value2, tn.son[1].son[0].token.Value.(string))
}
func (tn *TreeNode) GetNumber() float64 {
	switch tn.token.Symbol {
	case "E":
		return tn.son[1].GetNumberWithLeftNumber(tn.son[0].GetNumber())
	case "T":
		return tn.son[1].GetNumberWithLeftNumber(tn.son[0].GetNumber())
	case "F":
		switch {
		case len(tn.son) == 3:
			// 这个表示的是有括号的情况
			return tn.son[1].GetNumber()
		case tn.son[0].token.Symbol == "sz":
			number, _ := strconv.ParseFloat(tn.son[0].token.Value.(string), 64)
			return number
		case tn.son[0].token.Symbol == "ide":
			return ideStorage[tn.son[0].token.Symbol]
		}
	}
	panic("[panic] 存在没考虑的情况 " + tn.token.Symbol)
}
func (tn *TreeNode) GetNumberWithLeftNumber(leftNumber float64) float64 {
	switch tn.token.Symbol {
	case "E'", "T'", "F'":
		if tn.son[0].token.Symbol == "BLA" {
			return leftNumber
		}
		leftValue := calculate(leftNumber, tn.son[1].GetNumber(), tn.son[0].token.Symbol)
		return tn.son[2].GetNumberWithLeftNumber(leftValue)
	}
	panic("[panic] 存在没考虑的情况")
}



func judge(leftValue float64, rightValue float64, opt string) bool {
	switch opt {
	case "<":
		return leftValue < rightValue
	case "<=":
		return leftValue <= rightValue
	case ">=":
		return leftValue >= rightValue
	case "==":
		return leftValue == rightValue
	case ">":
		return leftValue > rightValue
	default:
		return false
	}
}
func calculate(leftValue float64, rightValue float64, opt string) float64 {
	switch opt {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	case "/":
		return leftValue / rightValue
	default:
		return leftValue
	}
}
