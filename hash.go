package hash

func DJBHash(str string) uint {
	var hash uint
	length := len(str)

	for i := 0; i < length; i++ {
		hash = ((hash << 5) + hash) + uint(str[i])
	}

	return hash
}
