package main

import "strconv"

func Encode(text string) string {
	result := ""

	for _, b := range []byte(text) {
		binary := toBinary(b)
		repeated := repeatEachThreeTimes(binary)
		binaryToAscii(repeated)

		result += string(repeated)
	}

	return result
}

func Decode(bits string) string {
	chunks := chunkEvery(bits, 3)
	corrected := correct(chunks)
	correctedBytes := chunkEvery(string(corrected), 8)

	return binariesToAscii(correctedBytes)
}

func toBinary(x byte) []byte {
	result := make([]byte, 8, 8)
	index := 7

	for x > 0 {
		if x%2 > 0 {
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
	new := make([]byte, 0, len(binary)*3)

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

func chunkEvery(str string, n int) [][]byte {
	bytes := []byte(str)
	result := make([][]byte, len(str)/n+len(str)%n)
	for i := range result {
		result[i] = make([]byte, n)
	}

	for i, byte := range bytes {
		remainder := i % n
		switch remainder {
		case 0:
			result[i/n][0] = byte
		default:
			result[i/n][remainder] = byte
		}
	}

	return result
}

func binariesToAscii(binary [][]byte) string {
	result := ""

	for _, bytes := range binary {
		ascii, _ := strconv.ParseInt(bytesToString(bytes), 2, 64)
		result += string(byte(ascii))
	}

	return result
}

func bytesToString(bytes []byte) string {
	result := ""
	for _, b := range bytes {
		result += strconv.Itoa(int(b))
	}

	return result
}

func correct(bytes [][]byte) []byte{
	result := make([]byte, len(bytes))

	for i, byte := range bytes {
		if sumBinary(byte) > 1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}

func sumBinary(binary []byte) int {
	result := 0

	for _, b := range binary {
		result += int(b) - 48
	}

	return result
}
