package main

func Encode(text string) string {
	result := ""

	for _, b := range []byte(text){
		binary := toBinary(b)
		repeated := repeatEachThreeTimes(binary)
		binaryToAscii(repeated)

		result += string(repeated)
	}

	return result
}

func Decode(bits string) string {
  return ""
}

func toBinary(x byte) []byte {
	result := make([]byte, 8, 8)
	index := 7

	for x > 0 {
		if x % 2 > 0 {
			result[index] = 1
		} else {
			result[index] = 0
		}

		index -= 1
		x = x / 2
	}

	return result
}

func repeatEachThreeTimes(binary []byte) []byte {
	new := make([]byte, 0, len(binary) * 3)

	for _, b := range binary {
		new = append(new, b, b, b)
	}

	return new
}

func binaryToAscii(binary []byte) {
	for i, b := range binary {
		binary[i] = b + 48
	}
}
