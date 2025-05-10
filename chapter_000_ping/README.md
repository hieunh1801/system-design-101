# Chapter 000

## Architecture
- 1 nginx
- 2 go

## Go 
```bash
# to run
go run main.go

# to build
docker build -t myapp .
docker run myapp
docker compose up
docker compose up app1
docker compose up app1 app2 nginx
```

## Test benchmark
```bash
wrk -t12 -c400 -d30s http://localhost:8080/ping # app1
wrk -t12 -c400 -d30s http://localhost:8081/ping # app2
wrk -t12 -c400 -d30s http://localhost/ping # nginx

# Running 30s test @ http://localhost:8080/ping
#   12 threads and 400 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency    13.03ms   13.27ms 174.00ms   87.22%
#     Req/Sec     3.18k   813.33     5.11k    65.86%
#   1139811 requests in 30.06s, 163.05MB read
#   Socket errors: connect 0, read 359, write 0, timeout 0
# Requests/sec:  37912.96
# Transfer/sec:      5.42MB
```
