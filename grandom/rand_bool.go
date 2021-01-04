package grandom

func RandBool() bool {
	randI := Random() * float64(11)

	if randI > 5 {
		return true
	}

	return false
}
