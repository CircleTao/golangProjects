/*
实现ping操作与ICMP协议
参考视频：https://www.bilibili.com/video/BV1CV4y1D7A4 P1
代码复现：CircleTAO
日期：2024年2月29日16:01:39
*/

// 注：按照视频中ping www.baidu.com会出现“wsasend: An attempt was made to access a socket in a way forbidden by its access permissions.”
// 的错误，暂时使用www.icourse163.org代替则正常，推测是ping百度时是经ipv6的而不是ipv4（我这里用终端ping得到的也是ipv6地址的回复），而ping中国慕课
// 得到的是ipv4的回复，此时代码可以正常运行

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"
)

var (
	timeout      int64
	size         int
	count        int
	typ          uint8 = 8
	code         uint8 = 0
	sendCount    int
	successCount int
	failCount    int
	minTs        int64 = math.MaxInt32
	maxTs        int64
	totalTs      int64
)

type ICMP struct { // 定义ICMP结构体
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func main() {
	getCommandArgs()
	desIp := os.Args[len(os.Args)-1] // 获取最后一个参数，即目标IP地址
	// func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
	// DialTimeout类似Dial,在网络network上连接地址address，并返回一个Conn接口,但采用了超时。timeout参数如果必要可包含名称解析。
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil { // 错误处理，若连接失败则返回
		log.Fatal(err)
		return
	}
	// Close方法关闭该连接，并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
	defer conn.Close()
	fmt.Printf("正在 Ping %s [%s] 具有 %d 字节的数据：\n", desIp, conn.RemoteAddr(), size)
	for i := 0; i < count; i++ {
		sendCount++
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			CheckSum:    0,
			ID:          1,
			SequenceNum: 1,
		}

		data := make([]byte, size)

		// Buffer是一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲。
		var buffer bytes.Buffer
		// func Write(w io.Writer, order ByteOrder, data interface{}) error
		// 将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针。order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。
		binary.Write(&buffer, binary.BigEndian, icmp) // BigEndian意为大端写入
		// func (b *Buffer) Write(p []byte) (n int, err error)
		// Write将p的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
		buffer.Write(data)
		// func (b *Buffer) Bytes() []byte
		// 返回未读取部分字节数据的切片，len(b.Bytes()) == b.Len()。如果中间没有调用其他方法，修改返回的切片的内容会直接改变Buffer的内容。
		data = buffer.Bytes()
		checkSum := checkSum(data)
		data[2] = byte(checkSum >> 8) // checkSum右移8位，并转byte，获取到checkSum中的高8位的值，若不右移则会取到低8位
		data[3] = byte(checkSum)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond)) // 设置超时时间
		t1 := time.Now()
		n, err := conn.Write(data)
		if err != nil {
			failCount++
			log.Println(err)
			continue
		}

		buf := make([]byte, 65535)
		n, err = conn.Read(buf)
		if err != nil {
			failCount++
			log.Println(err)
			continue
		}
		successCount++
		ts := time.Since(t1).Milliseconds() // 计算从t1到此刻的时间
		if minTs > ts {
			minTs = ts
		}
		if maxTs < ts {
			maxTs = ts
		}
		totalTs += ts
		fmt.Printf("来自 %d.%d.%d.%d 的回复：字节=%d 时间=%d ms TTL=%d \n", buf[12], buf[13], buf[14], buf[15], n-28, ts, buf[8])
	}
	fmt.Printf("%s 的 Ping 统计信息: \n 数据包: 已发送 = %d, 已接收 = %d, 丢失 = %d (%.2f%% 丢失), \n 往返行程的估计时间(以毫秒为单位):\n 最短 = %dms,最长 = %dms,平均 = %dms",
		conn.RemoteAddr(), sendCount, successCount, failCount, float64(failCount)/float64(sendCount)*100, minTs, maxTs, totalTs/int64(sendCount))
}

func getCommandArgs() {
	// func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
	// Int64Var用指定的名称、默认值、使用信息注册一个int64类型flag，并将flag的值保存到p指向的变量。
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长，单位：毫秒")

	// func IntVar(p *int, name string, value int, usage string)
	// IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量。
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，单位：字节")
	flag.IntVar(&count, "n", 4, "发送请求数")

	// 从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
}

func checkSum(data []byte) uint16 { // 校验功能的实现
	// 将校验位两两拼接并求和
	length := len(data)
	index := 0
	var sum uint32 = 0
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	// 若为奇数，将剩余的一位累加求和
	if length != 0 {
		sum += uint32(data[index])
	}
	// 重复操作使高16位为0
	high16 := sum >> 16
	for high16 != 0 {
		sum = high16 + uint32(uint16(sum))
		high16 = sum >> 16
	}
	// 将计算结果取反，返回16位的校验位
	return uint16(^sum)
}
