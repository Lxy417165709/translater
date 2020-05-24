package machine


const (
	endChar = byte('#')
	Eps = byte(0)
	additionalCharOfMatchingOnceAtLess = byte('@')
	additionalCharOfMatchingZeroTimesAtLess = byte('$')
)

var additionalChars = []byte{
	additionalCharOfMatchingOnceAtLess,
	additionalCharOfMatchingZeroTimesAtLess,
}




