package utils

func ChunkString(s string, chunkSize int) []string {
	sLen := len(s)
	resChunks := make([]string, 0, sLen-1/chunkSize+1)

	if sLen == 0 {
		return resChunks
	}

	if chunkSize >= sLen {
		return []string{s}
	}

	runes := []rune(s)
	currLen := 0
	currStart := 0
	for i := range runes {
		if currLen == chunkSize {
			resChunks = append(resChunks, s[currStart:i])
			currLen = 0
			currStart = i
		}
		currLen++
	}
	resChunks = append(resChunks, s[currStart:])
	return resChunks
}
