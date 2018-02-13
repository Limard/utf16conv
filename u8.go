package utf16conv

func U8ByteToString(s []byte) string {
	for i, v := range s {
		if v == 0 {
			s = s[0:i]
			break
		}
	}
	return string(s)
}
