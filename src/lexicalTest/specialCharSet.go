package lexicalTest

const eps = ' '

func init() {
	for i:=byte('0');i<='9';i++{
		digits=append(digits,i)
	}
	for i:=byte('a');i<='z';i++{
		letters=append(letters,i)
	}
	for i:=byte('A');i<='Z';i++{
		letters=append(letters,i)
	}
}
// 代表字符 'D'
var digits []byte

// 代表字符 'L'
var letters []byte

func isDigit(ch byte) bool{
	for i:=0;i<len(digits);i++{
		if digits[i]==ch{
			return true
		}
	}
	return false
}
func isLetter(ch byte) bool{
	for i:=0;i<len(letters);i++{
		if letters[i]==ch{
			return true
		}
	}
	return false
}
