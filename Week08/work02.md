作业：

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

Redis字符串占用的内存，远比实际字符串的长度要大。

Redis 字符串实际的数据结构包括了以下几个内容

- buf：字节数组，保存实际数据。为了表示字节数组的结束，Redis 会自动在数组最后加一个“\0”，这就会额外占用 1 个字节的开销。
- len：占 4 个字节，表示 buf 的已用长度。
- alloc：也占个 4 字节，表示 buf 的实际分配长度，一般大于 len。

在 SDS 中，buf 保存实际数据，而 len 和 alloc 本身其实是 SDS 结构体的额外开销。

另外，对于 String 类型来说，除了 SDS 的额外开销，还有一个来自于 RedisObject 结构体的开销。

可以根据情况选择其他的数据类型来存储。

```
开始时间: 2021-06-18 15:50:11.7333446 +0800 CST m=+0.022999701 ,key的前缀为 testGoCamp:20:; Value的长度：20， 内存统计开始:
# Memory
used_memory:276779232
used_memory_human:263.96M
used_memory_rss:309612544
used_memory_rss_human:295.27M
used_memory_peak:355365320
used_memory_peak_human:338.90M
total_system_memory:8201560064
total_system_memory_human:7.64G
used_memory_lua:49152
used_memory_lua_human:48.00K
maxmemory:4000000000
maxmemory_human:3.73G
maxmemory_policy:allkeys-lru
mem_fragmentation_ratio:1.12
mem_allocator:jemalloc-3.6.0

数据的个数 100000.000000 ；内存统计结束.结束时间：2021-06-18 15:50:28.6026471 +0800 CST m=+16.893533501# Memory
used_memory:290346024
used_memory_human:276.90M
used_memory_rss:315568128
used_memory_rss_human:300.95M
used_memory_peak:355365320
used_memory_peak_human:338.90M
total_system_memory:8201560064
total_system_memory_human:7.64G
used_memory_lua:51200
used_memory_lua_human:50.00K
maxmemory:4000000000
maxmemory_human:3.73G
maxmemory_policy:allkeys-lru
mem_fragmentation_ratio:1.09
mem_allocator:jemalloc-3.6.0


计算：（290346024-276779232)/100000  =  133.67
```


```
开始时间: 2021-06-18 15:54:04.790872 +0800 CST m=+0.027024901 ,key的前缀为 testGoCamp:50:; Value的长度：50， 内存统计开始:
# Memory
used_memory:276963520
used_memory_human:264.13M
used_memory_rss:284516352
used_memory_rss_human:271.34M
used_memory_peak:355365320
used_memory_peak_human:338.90M
total_system_memory:8201560064
total_system_memory_human:7.64G
used_memory_lua:46080
used_memory_lua_human:45.00K
maxmemory:4000000000
maxmemory_human:3.73G
maxmemory_policy:allkeys-lru
mem_fragmentation_ratio:1.03
mem_allocator:jemalloc-3.6.0

数据的个数 100000.000000 ；内存统计结束.结束时间：2021-06-18 15:54:24.1798718 +0800 CST m=+19.416695601# Memory
used_memory:293567008
used_memory_human:279.97M
used_memory_rss:300339200
used_memory_rss_human:286.43M
used_memory_peak:355365320
used_memory_peak_human:338.90M
total_system_memory:8201560064
total_system_memory_human:7.64G
used_memory_lua:49152
used_memory_lua_human:48.00K
maxmemory:4000000000
maxmemory_human:3.73G
maxmemory_policy:allkeys-lru
mem_fragmentation_ratio:1.02
mem_allocator:jemalloc-3.6.0

计算：（293567008-276963520)/100000  =  166.03
```


```

开始时间: 2021-06-18 15:56:42.5002118 +0800 CST m=+0.014999001 ,key的前缀为 testGoCamp:50:; Value的长度：50， 内存统计开始:
# Memory
used_memory:276874048
used_memory_human:264.05M
used_memory_rss:300490752
used_memory_rss_human:286.57M
used_memory_peak:355365320
used_memory_peak_human:338.90M
total_system_memory:8201560064
total_system_memory_human:7.64G
used_memory_lua:51200
used_memory_lua_human:50.00K
maxmemory:4000000000
maxmemory_human:3.73G
maxmemory_policy:allkeys-lru
mem_fragmentation_ratio:1.09
mem_allocator:jemalloc-3.6.0

数据的个数 500000.000000 ；内存统计结束.结束时间：2021-06-18 15:58:36.345897 +0800 CST m=+113.884446101# Memory
used_memory:354366408
used_memory_human:337.95M
used_memory_rss:361537536
used_memory_rss_human:344.79M
used_memory_peak:355365320
used_memory_peak_human:338.90M
total_system_memory:8201560064
total_system_memory_human:7.64G
used_memory_lua:49152
used_memory_lua_human:48.00K
maxmemory:4000000000
maxmemory_human:3.73G
maxmemory_policy:allkeys-lru
mem_fragmentation_ratio:1.02
mem_allocator:jemalloc-3.6.0


计算：（354366408-276874048)/500000  =  154.98 
```
