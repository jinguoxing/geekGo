package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
)

type GoImPackage struct {
	PackageLen uint32            // 包长度
	HeadLen    uint16            // 头长度
	Head       map[string]string // 头的具体内容
	Protocol   uint16            // 版本
	Operation  uint32            // 操作码
	SequenceId uint32            // 自增id
	Body       string            //  消息的内容
	ByteStr    []byte            // 源数据
}

// 解包
func (l *GoImPackage) GoImConnDecode() {
	content := l.ByteStr
	// 解析出包的长度 4个字节
	l.PackageLen = binary.LittleEndian.Uint32(content[0:4])

	// 解析出头的长度 2个字节
	l.HeadLen = binary.LittleEndian.Uint16(content[4:6])

	//解析出版本 2个字节
	l.Protocol = binary.LittleEndian.Uint16(content[6:8])
	//解析出版本 4个字节
	l.Operation = binary.LittleEndian.Uint32(content[8:12])
	// 解析出自增id 4个字节
	l.SequenceId = binary.LittleEndian.Uint32(content[12:16])

	// 解析head
	headContent := content[16 : l.HeadLen+16]
	json.Unmarshal(headContent, &l.Head)

	// 剩下的是解析内容
	l.Body = string(content[int(l.HeadLen)+16:l.PackageLen])
}

// 拆分goIm 消息
func SplitMsg(content []byte) {
	currentPos := 0
	packIndexLen := 4

	// 解析出所有的消息，分解个一个个单个消息体
	for currentPos < len(content) {
		packageLen := int(binary.LittleEndian.Uint32(content[currentPos:packIndexLen+currentPos]))
		data := content[currentPos : packIndexLen+packageLen +currentPos]
		hand := &GoImPackage{
			ByteStr: data,
		}

		// 解析到具体消息
		hand.GoImConnDecode()

		// 记录下一个位置
		currentPos += int(packageLen)
		fmt.Println("currentPos",currentPos,packageLen)

		// 打印解析结果
		fmt.Println("--------------------------------------------")
		fmt.Println("解析结果:版本", hand.Protocol)
		fmt.Println("解析结果:操作码", hand.Operation)
		fmt.Println("解析结果:唯一id", hand.SequenceId)
		fmt.Println("解析结果:操作码", hand.Operation)
		fmt.Println("解析结果:头", hand.Head)
		fmt.Println("解析结果:消息内容", hand.Body)
		fmt.Println("--------------------------------------------")
	}
}

// body 消息的内容
// head 自定义头部
// 操作码 1 是心跳 2 是认证 3 是发送信息
func SendGoImConn(body string, head map[string]string, op uint32) *GoImPackage {
	msg := &GoImPackage{
		Body:       body,
		Head:       head,
		Operation:  op,
		SequenceId: rand.Uint32(),
		Protocol:   uint16(1),
	}

	// 计算头部的长度
	headByte, _ := json.Marshal(head)
	msg.HeadLen = uint16(len(headByte))

	// 计算消息内容
	bodyByte := []byte(body)

	// 构造详细数据
	byteStr := make([]byte, 0)

	// 写入包长度 16 = 4个长度的包长度+ 2个长度的头长度 +2个长度的版本 + 4个长度的操作符 + 4个字符的唯一id
	msg.PackageLen = uint32(len(headByte)) + uint32(len(bodyByte)) + 16
	bytePackAge := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytePackAge, msg.PackageLen)
	byteStr = append(byteStr, bytePackAge...)

	// 写入head长度 2个字节
	byteHeadLen := make([]byte, 2)
	binary.LittleEndian.PutUint16(byteHeadLen, msg.HeadLen)
	byteStr = append(byteStr, byteHeadLen...)

	// 写版本 长度为2
	byteProtocol := make([]byte, 2)
	binary.LittleEndian.PutUint16(byteProtocol, msg.Protocol)
	byteStr = append(byteStr, byteProtocol...)

	// 写入操作符
	byteOperation := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteOperation, msg.Operation)
	byteStr = append(byteStr, byteOperation...)

	// 写入唯一id,目前是随机数
	byteSequenceId := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteSequenceId, msg.SequenceId)
	byteStr = append(byteStr, byteSequenceId...)

	// 写入head
	byteStr = append(byteStr, headByte...)
	fmt.Println("head_check",byteStr,headByte)

	// 写入消息详细内容
	byteStr = append(byteStr, bodyByte...)
	msg.ByteStr = byteStr

	return msg
}

