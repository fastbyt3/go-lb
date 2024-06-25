# Build your own load balancer (Layer 7 HTTP Load Balancer)

> https://codingchallenges.fyi/challenges/challenge-load-balancer/#share-your-solutions

## Success criteria

Load balancer is for ensuring requests are distributed across all server nodes that are capable of handling that request

1. Minimize response time and maximise utilization whilst ensuring no server is overloaded
2. if a server goes offline, LBs should redirect traffic to nodes apart from the one that went down
3. Health check the servers
4. handle servers coming back online (passing a health check)

## Steps

### Step 1

- create a basic http server that listens for connections and forwards them to a single server

```bash
./lb
Received request from 127.0.0.1
GET / HTTP/1.1
Host: localhost
User-Agent: curl/7.85.0
Accept: */*
```
