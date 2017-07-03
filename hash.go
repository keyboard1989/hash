package hash

func DJBHash(str string) uint {
	var hash uint
	hash = 5381
	length := len(str)

	for i := 0; i < length; i++ {
		hash = ((hash << 5) + hash) + uint(str[i])
	}

	return hash
}
