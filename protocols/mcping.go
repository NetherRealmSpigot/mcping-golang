package protocols

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

func connect(address string, timeout int) (net.Conn, error) {
	return net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
}

func doConnect(ips []net.IP, port uint16, timeout int) (net.Conn, error) {
	var conn net.Conn
	for i, ip := range ips {
		remoteAddress := fmt.Sprintf("%s:%d", ip.String(), port)
		conn0, err := connect(remoteAddress, timeout)
		conn = conn0
		if err != nil {
			if i == len(ips)-1 {
				return nil, fmt.Errorf("all ip addresses are not reachable, giving up: %w", err)
			}
			continue
		} else {
			break
		}
	}
	return conn, nil
}

func readByte(conn *net.Conn, buf []byte, n int) error {
	if n <= 0 {
		return errors.New("read length must be greater than zero")
	}

	_, err := io.ReadFull(*conn, buf[:n])
	if err != nil {
		return errors.Join(errors.New("failed to read bytes"), err)
	}
	return nil
}

func writeByte(conn *net.Conn, buf []byte) error {
	_, err := (*conn).Write(buf)
	return err
}

func composePacket(host string, port uint16, protocolNumVarint []byte) []byte {
	hostLen := len(host)
	hostLenVarint := CreateVarint(hostLen)
	hostLenVarintByteLen := len(hostLenVarint)

	protocolNumVarintByteLen := len(protocolNumVarint)

	handshakeLen := 1 /* Packet ID */ + protocolNumVarintByteLen + hostLenVarintByteLen + hostLen + 2 + 1 /* Port. State. */
	handshakeLenVarint := CreateVarint(handshakeLen)
	handshakeLenVarintByteLen := len(handshakeLenVarint)

	rawContent := make([]byte, handshakeLenVarintByteLen+handshakeLen+2) /* Handshake + Ping */

	i := 0

	copy(rawContent[i:], handshakeLenVarint)
	i += handshakeLenVarintByteLen

	rawContent[i] = 0x00 // Packet ID
	i++

	copy(rawContent[i:], protocolNumVarint)
	i += protocolNumVarintByteLen

	copy(rawContent[i:], hostLenVarint)
	i += hostLenVarintByteLen

	copy(rawContent[i:], host)
	i += hostLen

	binary.BigEndian.PutUint16(rawContent[i:], port)
	i += 2

	rawContent[i] = 0x01 // Next state
	i++
	// Ping packet
	rawContent[i] = 0x01
	i++
	rawContent[i] = 0x00 // Packet ID

	return rawContent
}

func Ping(host string, port uint16, fakeHost string, protocol int, timeout int) (string, uint16, string, int, []byte, error) {
	if host == "" {
		return "", 0, "", 0, nil, errors.New("invalid host string")
	}

	var ips []net.IP

	if ip := net.ParseIP(host); ip != nil {
		if ip4 := ip.To4(); ip4 == nil {
			return "", 0, "", 0, nil, errors.New("invalid host IP")
		} else {
			ip = ip4
		}
		ips = append(ips, ip)
	}
	if !IsKnownProtocolNumber(protocol) {
		return "", 0, "", 0, nil, fmt.Errorf("unknown protocol number %d", protocol)
	}
	if timeout <= 0 {
		return "", 0, "", 0, nil, errors.New("timeout in seconds must be bigger than 0")
	}

	if len(ips) == 0 {
		if port <= 0 {
			srvips, srvport := ResolveFromSRV(host)
			if srvips != nil {
				ips = append(ips, srvips...)
				port = srvport
			}
		}
		if len(ips) == 0 {
			hostips, err := Resolve(host)
			if err != nil {
				return "", 0, "", 0, nil, err
			}
			ips = append(ips, hostips...)
		}
	}
	if port <= 0 {
		port = 25565
	}
	if fakeHost == "" {
		fakeHost = host
	}

	conn, err := doConnect(ips, port, timeout)
	if err != nil {
		return host, port, fakeHost, 0, nil, err
	}

	protocolNumVarint := CreateVarint(protocol)
	packet := composePacket(fakeHost, port, protocolNumVarint)
	// Send packet
	if err := writeByte(&conn, packet); err != nil {
		return host, port, fakeHost, protocol, nil, err
	}

	// Read response. Read packet length.
	if _, err := ParseVarint(&conn); err != nil {
		return host, port, fakeHost, protocol, nil, err
	}

	// Read packet ID
	buf := make([]byte, 1)
	if err := readByte(&conn, buf, 1); err != nil {
		return host, port, fakeHost, protocol, nil, err
	}
	if buf[0] != byte(0) {
		return host, port, fakeHost, protocol, nil, errors.New(fmt.Sprintf("invalid packet. unknown packet id %d", buf[0]))
	}
	jsonPayloadLen, err := ParseVarint(&conn)
	buf = make([]byte, jsonPayloadLen)
	if err := readByte(&conn, buf, jsonPayloadLen); err != nil {
		return host, port, fakeHost, protocol, nil, err
	}
	return host, port, fakeHost, protocol, buf, nil
}
