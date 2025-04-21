package protocols

import (
	"errors"
	"net"
)

func CreateVarint(num int) []byte {
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

func ParseVarint(conn *net.Conn) (int, error) {
	numRead := 0
	result := 0
	buf := make([]byte, 1)
	for {
		if err := readByte(conn, buf, 1); err != nil {
			return -1, err
		}
		result |= (int(buf[0]) & 0x7F) << (7 * numRead)
		numRead++
		if numRead > 5 {
			return -1, errors.New("not a varint. read too many bytes")
		}
		if (buf[0] & 0x80) == 0 {
			break
		}
	}
	return result, nil
}
