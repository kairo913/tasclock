package utility

import "crypto/rand"

func MakeRandomStr(length int) string {
	b := make([]byte, length)
	for i := 0; i < 10; i++ {
		if _, err := rand.Read(b); err == nil {
			break
		}
		if i == 9 {
			return ""
		}
	}

	var result string
	for _, v := range b {
		result += string(v%byte(94) + 33)
	}
	return result
}