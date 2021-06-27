package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	// based/length 调用开始
	fmt.Println("按固定长度发包和拆包开始,发送数据为")
	// 按固定长度发包和拆包 5个字节为一个包，发送数据和接受数据长度固定，缺点是每次发送的数据长度必须是5的倍数
	arr := basedLengthSend("abcd,efg1,2345,6789,", 5)
	basedLengthRecv(arr,5)
	fmt.Println("按固定长度发包和拆包结束")


	// fix length/delimiter 调用开始
	fmt.Println("特殊分割符拆包开始 ")
	delimite := []byte("\r")[0]
	recv := fixLengthRecv([]byte("abcd\refg12345\r6789,\r"), delimite)
	for k, v := range recv {
		fmt.Println("第"+strconv.Itoa(k+1)+"个包数据为", string(v))
	}

	// 解析简单版的 自定义长度包的口EXO
	// testFixFrame()

	// 测试解析go im
	testGoImConn()

}

// goIm 连接解析
func testGoImConn() {
	fmt.Println("尝试解析goIm连接尝试解析开始")
	fmt.Println("构造三个数据包")
	head := make(map[string]string)
	head["method"] = "post"
	pack1 := SendGoImConn("第一个包，认证", head, 1)
	pack2 := SendGoImConn("第二个包，心跳检测", head, 2)
	pack3 := SendGoImConn("第三个包，Hello word", head, 3)
	allByte := make([]byte, 0)
	allByte = append(allByte, pack1.ByteStr...)
	allByte = append(allByte, pack2.ByteStr...)
	allByte = append(allByte, pack3.ByteStr...)

	// 尝试拆包
	SplitMsg(allByte)

	fmt.Println("尝试解析goIm连接尝试解析结束")

}

// 简单尝试解析自定义长度的包
func testFixFrame() {
	headLen := 2
	// 拼接第1个包
	sendPack1 := lenFixFrameSend("1", headLen)
	// 拼接第2个包
	sendPack2 := lenFixFrameSend("5，3", headLen)
	// 拼接第3个包
	sendPack3 := lenFixFrameSend("aaa,bbb,ccc", headLen)

	allBytes := make([]byte, 0)
	allBytes = append(allBytes, sendPack1...)
	allBytes = append(allBytes, sendPack2...)
	allBytes = append(allBytes, sendPack3...)

	lenFixFrameRecv(allBytes, 2)
}

// 固定长度封包 按5个字节为一个包
func basedLengthSend(content string, num int64) [][]byte {
	arr := []byte(content)
	max := int64(len(arr))
	var segmens = make([][]byte, 0)
	if max < num {
		segmens = append(segmens, arr)
		return segmens
	}

	quantity := max / num
	for i := int64(1); i <= quantity; i++ {
		end := i * num
		segmens = append(segmens, arr[(i-1)*num:end])

	}
	return segmens
}

// 固定长度拆包 按5个字节为一个包
func basedLengthRecv(content [][]byte, num int64) {
	for k, v := range content {

		fmt.Printf("第%d个包数据为%v \n", k, string(v))
	}
}

// 分隔符字符的手包
func fixLengthRecv(content []byte, delimite byte) [][]byte {
	var segmens = make([][]byte, 0)
	pos := 0
	for k, v := range content {
		// 如果定位到了分隔符,就切割
		if v == delimite {
			segmens = append(segmens, content[pos:k-1])
			pos = k + 1
			continue
		}
	}
	// 如果一个也没找到，就把整体发送处理
	if pos == 0 {
		segmens = append(segmens, content)
	}

	return segmens
}

// 头部2个字节，后面为数据长度
func lenFixFrameSend(content string, num int) []byte {
	arr := []byte(content)
	var headLen uint16
	headLen = uint16(len(arr))

	// 生命一个新数组
	newLen := 2 + int(headLen)
	newArr := make([]byte, newLen, newLen)

	//将头部的长度转换为字节
	head := make([]byte, num)
	binary.LittleEndian.PutUint16(head, headLen)

	// 先把head 的长度写进去
	for i := 0; i < len(head); i++ {
		newArr[i] = head[i]
	}

	for k, v := range arr {
		newArr[k+num] = v
	}
	return newArr
}

func lenFixFrameRecv(content []byte, num int) {
	fmt.Println(content, len(content))

	currentPos := 0
	index := 1
	for currentPos < len(content) {
		head := content[currentPos : num+currentPos]
		dataLen := binary.LittleEndian.Uint16(head)
		fmt.Printf("第%d次数据包长度为:%d", index, dataLen)
		// 读取数据
		nextPos := int(dataLen) + num
		data := content[num+currentPos : nextPos+currentPos]
		fmt.Println(" 内容数据为", string(data))

		currentPos += nextPos
	}

}
