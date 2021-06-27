### 1、总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用。
### 2、实现一个从 socket connection 中解码出 goim 协议的解码器。


### 解决tcp 粘包大约有三种方式：
1：LengthFieldBasedFrameDecoder作用
在消息头中定义长度字段，来标识消息的总长度 固定值的包头包尾设定。缺点是：实现起来比较负责，需要计算好头的长度，消息内容的长度。优点是：非常灵活，方便定制。
2：fix length/delimiter
特殊字符结尾。优点：实现非常简单。缺点是，如果内容中含有自定义字符，需要转义。另外实现复杂结构，比如包括head 和 body 两种数据，传输起来可能比较复杂。
3： based/length
固定长度。服务端每次定期读取固定的字节。缺点：局限性比较大
