### 作业
1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。


### 系统配置 2核 8G

### 测试的场景，十万次，并发50 


#### 结论
 
 10 20 50 100 200 1k 5k 字节 value 大小，性能差别不大；GET的性能比SET的性能稍微差点；当超过10K，随着字节数越来越大，
 性能的表现越来越差。

### redis-benchmark使用方法


```
Usage: redis-benchmark [-h <host>] [-p <port>] [-c <clients>] [-n <requests]> [-k <boolean>]

 -h <hostname>      Server hostname (default 127.0.0.1)
 -p <port>          Server port (default 6379)
 -s <socket>        Server socket (overrides host and port)
 -a <password>      Password for Redis Auth
 -c <clients>       Number of parallel connections (default 50)
 -n <requests>      Total number of requests (default 100000)
 -d <size>          Data size of SET/GET value in bytes (default 2)
 --dbnum <db>        SELECT the specified db number (default 0)
 -k <boolean>       1=keep alive 0=reconnect (default 1)
 -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD
  Using this option the benchmark will expand the string __rand_int__
  inside an argument with a 12 digits number in the specified range
  from 0 to keyspacelen-1. The substitution changes every time a command
  is executed. Default tests use this to hit random keys in the
  specified range.
 -P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
 -e                 If server replies with errors, show them on stdout.
                    (no more than 1 error per second is displayed)
 -q                 Quiet. Just show query/sec values
 --csv              Output in CSV format
 -l                 Loop. Run the tests forever
 -t <tests>         Only run the comma separated list of tests. The test
                    names are the same as the ones produced as output.
 -I                 Idle mode. Just open N idle connections and wait.

Examples:

 Run the benchmark with the default configuration against 127.0.0.1:6379:
   $ redis-benchmark

 Use 20 parallel clients, for a total of 100k requests, against 192.168.1.1:
   $ redis-benchmark -h 192.168.1.1 -p 6379 -n 100000 -c 20

 Fill 127.0.0.1:6379 with about 1 million keys only using the SET test:
   $ redis-benchmark -t set -n 1000000 -r 100000000

 Benchmark 127.0.0.1:6379 for a few commands producing CSV output:
   $ redis-benchmark -t ping,set,get -n 100000 --csv

 Benchmark a specific command line:
   $ redis-benchmark -r 10000 -n 10000 eval 'return redis.call("ping")' 0

 Fill a list with 10000 random elements:
   $ redis-benchmark -r 10000 -n 10000 lpush mylist __rand_int__

 On user specified command lines __rand_int__ is replaced with a random integer
 with a range of values selected by the -r option.
```

### 参数的作用


| 作用分类 | 参数及作用 |
| ---- | ---- |
|连接 Redis 服务相关参数|-h ：Redis 服务主机地址，默认为 127.0.0.1 。|
|_ |-p ：Redis 服务端口，默认为 6379 。|
|_	|-s ：指定连接的 Redis 服务地址，用于覆盖 -h 和 -p 参数。一般情况下，我们并不会使用。|
|_	|-a ：Redis 认证密码。|
|	_|–dbnum ：选择 Redis 数据库编号。|
|	_|k ：是否保持连接。默认会持续保持连接。|
请求相关参数 |	-c ：并发的客户端数|
|	_|-n ：总共发起的操作（请求）数|
|	_|-d ：指定 SET/GET 操作的数据大小，单位：字节。|
|	_|-r ：SET/GET/INCR 使用随机 KEY ，SADD 使用随机值。通过设置-r参数，可以设置KEY的随机范围，比如-r 10生成的KEY范围为[0，9）|
|	_|-P ：默认情况下，Redis 客户端一次请求只发起一个命令。通过 -P 参数，可以设置使用 _|pipeline功能，一次发起指定个请求，从而提升 QPS 。|
|_|	-l ：循环，一直执行基准测试。|
|_|	-t ：指定需要测试的 Redis 命令，多个命令通过逗号分隔。默认情况下，测试 PING_INLINE/PING_BULK/SET/GET 等等命令。若只想测试 SET/GET 命令，则可以 -t SET,GET 来指定。|
|	_|-I ：Idle 模式。仅仅打开 N 个 Redis Idle 个连接，然后等待，啥也不做。不是很理解这个参数的目的，目前猜测，仅仅用于占用 Redis 连接。|
结果输出相关参数|
|_	|-e ：如果 Redis Server 返回错误，是否将错误打印出来。默认情况下不打印，通过该参数开启。|
|_|	-q ：精简输出结果。即只展示每个命令的 QPS 测试结果|
|_|	-csv ：按照 CSV 的格式，输出结果|


### 具体测试数据

```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 10
====== SET ======
  100000 requests completed in 1.07 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
93457.94 requests per second

====== GET ======
  100000 requests completed in 1.05 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.99% <= 1 milliseconds
100.00% <= 1 milliseconds
95602.30 requests per second
```

```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 20
SET: -nan
====== SET ======
  100000 requests completed in 1.48 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.73% <= 1 milliseconds
99.85% <= 2 milliseconds
99.88% <= 3 milliseconds
99.90% <= 6 milliseconds
99.95% <= 11 milliseconds
100.00% <= 11 milliseconds
67430.88 requests per second

====== GET ======
  100000 requests completed in 1.19 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.78% <= 1 milliseconds
99.92% <= 4 milliseconds
99.97% <= 5 milliseconds
100.00% <= 5 milliseconds
84317.03 requests per second
```

```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 50
====== SET ======
  100000 requests completed in 1.00 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.90% <= 1 milliseconds
100.00% <= 1 milliseconds
100200.40 requests per second

====== GET ======
  100000 requests completed in 1.06 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.96% <= 1 milliseconds
100.00% <= 1 milliseconds
94696.97 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 100
====== SET ======
  100000 requests completed in 1.02 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.96% <= 1 milliseconds
100.00% <= 1 milliseconds
98425.20 requests per second

====== GET ======
  100000 requests completed in 1.01 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.93% <= 1 milliseconds
100.00% <= 1 milliseconds
98522.17 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 200
====== SET ======
  100000 requests completed in 1.04 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.98% <= 1 milliseconds
100.00% <= 1 milliseconds
95877.28 requests per second

====== GET ======
  100000 requests completed in 1.02 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.96% <= 1 milliseconds
100.00% <= 1 milliseconds
98328.42 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 1000
====== SET ======
  100000 requests completed in 1.05 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.89% <= 1 milliseconds
99.95% <= 3 milliseconds
100.00% <= 3 milliseconds
95602.30 requests per second

====== GET ======
  100000 requests completed in 1.11 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.89% <= 1 milliseconds
100.00% <= 1 milliseconds
90497.73 requests per second
```

```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 2000
====== SET ======
  100000 requests completed in 1.01 seconds
  50 parallel clients
  2000 bytes payload
  keep alive: 1

99.93% <= 1 milliseconds
100.00% <= 1 milliseconds
99206.34 requests per second

====== GET ======
  100000 requests completed in 1.10 seconds
  50 parallel clients
  2000 bytes payload
  keep alive: 1

99.89% <= 1 milliseconds
100.00% <= 1 milliseconds
90579.71 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 4000
====== SET ======
  100000 requests completed in 1.08 seconds
  50 parallel clients
  4000 bytes payload
  keep alive: 1

99.95% <= 1 milliseconds
100.00% <= 1 milliseconds
92592.59 requests per second

====== GET ======
  100000 requests completed in 1.17 seconds
  50 parallel clients
  4000 bytes payload
  keep alive: 1

99.88% <= 1 milliseconds
100.00% <= 1 milliseconds
85397.09 requests per second

```

```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 5000
====== SET ======
  100000 requests completed in 1.15 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.85% <= 1 milliseconds
100.00% <= 1 milliseconds
86805.56 requests per second

====== GET ======
  100000 requests completed in 1.23 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.82% <= 1 milliseconds
100.00% <= 1 milliseconds
81499.59 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 10000
====== SET ======
  100000 requests completed in 1.19 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1

99.89% <= 1 milliseconds
100.00% <= 1 milliseconds
84033.61 requests per second

====== GET ======
  100000 requests completed in 1.32 seconds
  50 parallel clients
  10000 bytes payload
  keep alive: 1

99.75% <= 1 milliseconds
99.99% <= 2 milliseconds
100.00% <= 2 milliseconds
75757.57 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 15000
====== SET ======
  100000 requests completed in 1.23 seconds
  50 parallel clients
  15000 bytes payload
  keep alive: 1

99.89% <= 1 milliseconds
100.00% <= 1 milliseconds
81037.28 requests per second

====== GET ======
  100000 requests completed in 1.44 seconds
  50 parallel clients
  15000 bytes payload
  keep alive: 1

99.43% <= 1 milliseconds
99.98% <= 2 milliseconds
100.00% <= 2 milliseconds
69252.08 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 20000
====== SET ======
  100000 requests completed in 1.30 seconds
  50 parallel clients
  20000 bytes payload
  keep alive: 1

98.80% <= 1 milliseconds
100.00% <= 2 milliseconds
100.00% <= 2 milliseconds
76745.97 requests per second

====== GET ======
  100000 requests completed in 1.88 seconds
  50 parallel clients
  20000 bytes payload
  keep alive: 1

98.58% <= 1 milliseconds
99.94% <= 2 milliseconds
99.98% <= 5 milliseconds
100.00% <= 6 milliseconds
53219.80 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 40000
====== SET ======
  100000 requests completed in 1.85 seconds
  50 parallel clients
  40000 bytes payload
  keep alive: 1

92.57% <= 1 milliseconds
99.90% <= 2 milliseconds
99.99% <= 3 milliseconds
100.00% <= 3 milliseconds
54171.18 requests per second

====== GET ======
  100000 requests completed in 2.56 seconds
  50 parallel clients
  40000 bytes payload
  keep alive: 1

95.31% <= 1 milliseconds
99.97% <= 2 milliseconds
100.00% <= 2 milliseconds
39108.33 requests per second
```
```
[root@test-redis]# redis-benchmark -h nosql.redis.01.ickey.cn -a ickeyredis20150301 -p 6379 -t set,get -d 50000
====== SET ======
  100000 requests completed in 2.18 seconds
  50 parallel clients
  50000 bytes payload
  keep alive: 1

45.04% <= 1 milliseconds
99.86% <= 2 milliseconds
99.95% <= 3 milliseconds
99.96% <= 4 milliseconds
100.00% <= 4 milliseconds
45955.88 requests per second

====== GET ======
  100000 requests completed in 3.68 seconds
  50 parallel clients
  50000 bytes payload
  keep alive: 1

59.49% <= 1 milliseconds
99.57% <= 2 milliseconds
99.82% <= 3 milliseconds
99.83% <= 5 milliseconds
99.87% <= 6 milliseconds
99.93% <= 8 milliseconds
99.95% <= 9 milliseconds
99.95% <= 11 milliseconds
99.99% <= 12 milliseconds
100.00% <= 12 milliseconds
27173.91 requests per second

```