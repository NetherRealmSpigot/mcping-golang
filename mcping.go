package main

import (
	"context"
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

func doConnect(ips []net.IP, port uint16, timeout int, verbose bool) (net.Conn, error) {
	for _, ip := range ips {
		remoteAddress := fmt.Sprintf("%s:%d", ip.String(), port)
		if verbose {
			logToStderr(fmt.Sprintf("Connecting to %s", remoteAddress))
		}
		conn, err := connect(remoteAddress, timeout)
		if err != nil {
			continue
		}
		if verbose {
			logToStderr(fmt.Sprintf("Connected to %s", remoteAddress))
		}
		return conn, nil
	}
	return nil, errors.New("all ip addresses are not reachable, giving up")
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

func getIP(host string) ([]net.IP, error) {
	return net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
}

func getIPFromSRV(host string) ([]net.IP, uint16) {
	_, srvs, _ := net.LookupSRV("minecraft", "tcp", host)
	if len(srvs) > 0 {
		for _, srv := range srvs {
			srvips, err := getIP(srv.Target)
			if err != nil {
				continue
			}
			return srvips, srv.Port
		}
	}
	return nil, 0
}

func composePacket(host string, port uint16, protocolNumVarint []byte) []byte {
	hostLen := len(host)
	hostLenVarint := createVarint(hostLen)
	hostLenVarintByteLen := len(hostLenVarint)

	protocolNumVarintByteLen := len(protocolNumVarint)

	handshakeLen := 1 /* Packet ID */ + protocolNumVarintByteLen + hostLenVarintByteLen + hostLen + 2 + 1 /* Port. State. */
	handshakeLenVarint := createVarint(handshakeLen)
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

func parseVarint(conn *net.Conn) (int, error) {
	numread := 0
	result := 0
	buf := make([]byte, 1)
	for {
		if err := readByte(conn, buf, 1); err != nil {
			return -1, err
		}
		result |= (int(buf[0]) & 0x7F) << (7 * numread)
		numread++
		if numread > 5 {
			return -1, errors.New("not a varint. read too many bytes")
		}
		if (buf[0] & 0x80) == 0 {
			break
		}
	}
	return result, nil
}

func Ping(host string, port uint16, fakeHost string, protocol int, timeout int, verbose bool) ([]byte, error) {
	if host == "" {
		return nil, errors.New("invalid host string")
	}

	var ips []net.IP

	if ip := net.ParseIP(host); ip != nil {
		if ip4 := ip.To4(); ip4 == nil {
			return nil, errors.New("invalid host string")
		} else {
			ip = ip4
		}
		ips = append(ips, ip)
	}
	if port < 0 || port > 65535 {
		return nil, errors.New("invalid port")
	}
	if fakeHost == "" {
		fakeHost = host
	}
	if !IsKnownProtocolNumber(protocol) {
		return nil, errors.New("unknown protocol number")
	}
	if timeout <= 0 {
		if verbose {
			logToStderr(fmt.Sprintf("Ping timeout set to %d seconds", defaultTimeout))
		}
		timeout = defaultTimeout
	}

	if len(ips) == 0 {
		if port <= 0 {
			if verbose {
				logToStderr(fmt.Sprintf("Resolving _minecraft._tcp.%s", host))
			}
			srvips, srvport := getIPFromSRV(host)
			if srvips != nil {
				ips = append(ips, srvips...)
				port = srvport
			}
		}

		if len(ips) == 0 {
			if verbose {
				logToStderr(fmt.Sprintf("Resolving %s", host))
			}
			hostips, err := getIP(host)
			if err != nil {
				return nil, err
			}
			ips = append(ips, hostips...)
		}
	}

	if port <= 0 {
		if verbose {
			logToStderr(fmt.Sprintf("Port set to 25565"))
		}
		port = 25565
	}

	conn, err := doConnect(ips, port, timeout, verbose)
	if err != nil {
		return nil, err
	}

	if verbose {
		logToStderr(fmt.Sprintf("Protocol number %d", protocol))
	}
	protocolNumVarint := createVarint(protocol)
	packet := composePacket(fakeHost, port, protocolNumVarint)
	if verbose {
		logToStderr("Sending packet")
	}
	if err := writeByte(&conn, packet); err != nil {
		return nil, err
	}

	// Read response. Read packet length.
	if verbose {
		logToStderr(fmt.Sprintf("Receiving packet"))
	}
	packetLen, err := parseVarint(&conn)
	if err != nil {
		return nil, err
	}
	if verbose {
		logToStderr(fmt.Sprintf("Read packet length: %d", packetLen))
	}
	// Read packet ID
	buf := make([]byte, 1)
	if err := readByte(&conn, buf, 1); err != nil {
		return nil, err
	}
	if buf[0] != byte(0) {
		return nil, errors.New(fmt.Sprintf("invalid packet. unknown packet id %d", buf[0]))
	}
	jsonPayloadLen, err := parseVarint(&conn)
	if verbose {
		logToStderr(fmt.Sprintf("Read json payload length: %d", jsonPayloadLen))
	}
	buf = make([]byte, jsonPayloadLen)
	if err := readByte(&conn, buf, jsonPayloadLen); err != nil {
		return nil, err
	}
	return buf, nil
}
