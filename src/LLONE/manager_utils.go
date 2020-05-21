package LLONE

func (one *LLOne) flushDelimitersBuffer() {
	one.delimitersBuffer = map[string][]string{}
}
