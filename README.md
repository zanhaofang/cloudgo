# cloudgo
## 框架：Martini
原因：使用简单；类似于js，比较容易入手。
## 测试：
curl：

在终端使用curl：
```
curl -v http://localhost:3000
```
```
* Rebuilt URL to: http://localhost:3000/
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.38.0
> Host: localhost:3000
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Mon, 13 Nov 2017 11:40:11 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
```
ApacheBench:

在终端上用ApacheBench:
```
ab -n 1000 -c 100 http://localhost:3000
```
```
This is ApacheBench, Version 2.3 <$Revision: 1604373 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            3000

Document Path:          /
Document Length:        12 bytes

Concurrency Level:      100
Time taken for tests:   1.633 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      129000 bytes
HTML transferred:       12000 bytes
Requests per second:    1912.05 [#/sec] (mean)
Time per request:       53.562 [ms] (mean)
Time per request:       0.526 [ms] (mean, across all concurrent requests)
Transfer rate:          240.12 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   4.6      0      20
Processing:     1   48  13.7     47     104
Waiting:        1   48  13.7     47     104
Total:          2   54  13.0     49     104

Percentage of the requests served within a certain time (ms)
  50%    51
  66%    54
  75%    57
  80%    65
  90%    72
  95%    80
  98%    89
  99%    95
 100%    104 (longest request)
```
平均请求响应时间为53.562ms,50％的用户响应时间小于51毫秒，最大的响应时间小于104毫秒
