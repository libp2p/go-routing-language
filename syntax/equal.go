package syntax

func IsEqual(x, y Node) bool {
	switch x1 := x.(type) {
	case Bool:
		switch y1 := y.(type) {
		case Bool:
			return IsEqualBool(x1, y1)
		}
	case String:
		switch y1 := y.(type) {
		case String:
			return IsEqualString(x1, y1)
		}
	case Number:
		switch y1 := y.(type) {
		case Number:
			return IsEqualNumber(x1, y1)
		}
	case Bytes:
		switch y1 := y.(type) {
		case Bytes:
			return IsEqualBytes(x1, y1)
		}
	case Dict:
		switch y1 := y.(type) {
		case Dict:
			return IsEqualDict(x1, y1)
		}
	case List:
		switch y1 := y.(type) {
		case List:
			return IsEqualList(x1, y1)
		}
	case Predicate:
		switch y1 := y.(type) {
		case Predicate:
			return IsEqualPredicate(x1, y1)
		}
	}
	return false
}
