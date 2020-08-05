package crypto_tasks

func DecryptCeaser(message string, offset int) string {
	result := make([]byte, len(message))

	for i, c := range message {
		r := int(c) - 'a'
		r = (r + offset) % 26  + 'a'
		result[i] = byte(r)
	}
	return string(result)
}

func DecryptVigenere(message, keyword string) string {
	result := make([]byte, len(message))

	for i, c := range message {
		k := int32(keyword[i % len(keyword)])
		d := (c - k + 26) % 26 + 'a'
		result[i] = byte(d)
	}
	return string(result)
}
