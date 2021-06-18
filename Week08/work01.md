### 作业
1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。


### 系统配置 2核 8G

### 测试的场景，十万次，并发50 


#### 结论
 
 


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