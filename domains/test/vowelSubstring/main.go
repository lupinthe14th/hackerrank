package vowelsubstring

func findSubstring(s string, k int32) string {
	out := "Not found!"
	var t int32
	for i := 0; i < len(s)-int(k)+1; i++ {
		c := count(s[i : i+int(k)])
		if count(s[i:i+int(k)]) > t {
			out = s[i : i+int(k)]
			t = c
		}
	}
	return out
}

func count(s string) int32 {
	var out int32
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			out++
		}
	}
	return out
}
