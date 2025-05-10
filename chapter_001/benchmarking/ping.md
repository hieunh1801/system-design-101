```bash
wrk -t12 -c400 -d30s http://localhost:8088/ping

# Running 30s test @ http://127.0.0.1:8080/ping
#   12 threads and 400 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency    11.56ms   13.72ms 207.74ms   88.97%
#     Req/Sec     3.90k     1.46k    6.34k    62.49%
#   1396209 requests in 30.05s, 159.78MB read
#   Socket errors: connect 0, read 382, write 0, timeout 0
# Requests/sec:  46456.17
# Transfer/sec:      5.32MB


# Running 30s test @ http://localhost:8080/ping
#   12 threads and 400 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency    11.97ms    9.53ms 129.43ms   90.68%
#     Req/Sec     2.24k     1.21k    5.15k    62.77%
#   209973 requests in 30.10s, 27.83MB read
#   Socket errors: connect 0, read 256, write 0, timeout 2772
# Requests/sec:   6975.73
# Transfer/sec:      0.92MB
```