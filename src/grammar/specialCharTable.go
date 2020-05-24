package grammar



func (sct *SpecialCharTable) Show() {
	for _, item := range sct.specialCharItems {
		item.Show()
	}
}





