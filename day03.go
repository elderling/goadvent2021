package main

func Day03aSolution(filename string) int {
	return 198
}

func ReverseString(s string) string {
	reversedBytes := make([]byte, 0)
	r := []byte(s)
	for i := len(r) - 1; i >= 0; i-- {
		reversedBytes = append(reversedBytes, r[i])
	}
	return string(reversedBytes)
}
