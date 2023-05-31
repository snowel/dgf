package dgf

/*---  Utills  ---*/
// Infamous contains
func Contains[E comparable](val E, slice []E) bool {
  for _, v := range slice {
		if v == val {return true}
  }
  return false
}

func intPow(x int) int {
	return x * x
}

func ExclusiveAppend[E comparable](slice []E, val ...E) []E {
	for _, v := range val {
		if !Contains(v, slice) {slice = append(slice, v)}
	}

	return slice
}

func BoardCopy[V any](target []V, source []V) int {
	if len(target) != len(source) {return -1}
	
	for i, _ := range target {
		target[i] = source[i]
	}
	return 0
}
