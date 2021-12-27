package gnet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type pingICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func pingIcmpCheckSum(data []byte) (rt uint16) {
	var (
		sum    uint32
		length int = len(data)
		idex   int
	)
	for length > 1 {
		sum += uint32(data[idex])<<8 + uint32(data[idex+1])
		idex += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[idex]) << 8
	}
	rt = uint16(sum) + uint16(sum>>16)
	return ^rt
}
func Ping(domain string, duration time.Duration) (int64, error) {
	sendData := pingIcmpBytes()
	con, e := pingBuildIcmpCon(domain)
	if e != nil {
		fmt.Printf("链接错误\n")
		return 0, e
	}
	if _, e := con.Write(sendData); e != nil {
		fmt.Printf("发送错误\n")
		return 0, e
	}
	t_st := time.Now()
	con.SetReadDeadline(time.Now().Add(duration))
	recv := make([]byte, 1024)
	_, e = con.Read(recv)
	con.Close()
	if e != nil {
		fmt.Printf("读取错误 %v %s\n", con.LocalAddr().String(), con.RemoteAddr().String())
		//return 0, e
	}
	return time.Now().Sub(t_st).Nanoseconds(), nil
}
func pingIcmpBytes() []byte {
	icmp := pingICMP{8, 0, 0, 0, 0}
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = pingIcmpCheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)
	return buffer.Bytes()
}
func pingBuildIcmpCon(domain string) (*net.IPConn, error) {
	laddr := net.IPAddr{IP: net.ParseIP("0.0.0.0")}
	raddr, _ := net.ResolveIPAddr("ip", domain)
	return net.DialIP("ip4:icmp", &laddr, raddr)
}
