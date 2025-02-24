package main

func createVarint(num int) []byte {
	if num < 0 {
		return nil
	}

	v := uint(num)

	var buf = make([]byte, 0, 5)

	for count := 0; count < 5; {
		if (v & ^uint(0x7f)) == 0 {
			buf = append(buf, byte(v))
			break
		}
		buf = append(buf, byte((v&0x7f)|0x80))
		count++
		v >>= 7
	}

	return buf
}
