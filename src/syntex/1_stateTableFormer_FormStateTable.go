package syntex


func (stf *StateTableFormer)FormStateTable() {
	stf.stateTable = make(map[string]map[string]*sentence)
	for sentenc,terminators := range stf._select{
		for _,terminator := range terminators{
			if stf.stateTable[stf.sentenceToNonTerminator[sentenc]]==nil{
				stf.stateTable[stf.sentenceToNonTerminator[sentenc]] = make(map[string]*sentence)
			}
			stf.stateTable[stf.sentenceToNonTerminator[sentenc]][terminator]=sentenc
		}
	}
}

