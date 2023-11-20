# GC_metrics_in_Prometheus
Monitoring  the Garbage collection  of a basic http server collected via Prometheus 


This repository contains a simple Go web server application.<br>
It demonstrates basic HTTP request handling in Go and includes features to monitor system metrics like garbage collection.




# Features
  - Basic HTTP server implementation in Go.<br>
  - Endpoints for testing load and monitoring server responses.<br>
  - Integration examples for system metrics monitoring (garbage collection, memory usage, disk I/O).<br>



    
## Getting Started / Prerequisites <br>

Go (version 1.x or later)<br>
curl (for testing endpoints)<br>
Monitoring tools :  Prometheus and Grafana<br>

## Installation<br>


Clone the repository.<br>
Run  the server (`go run gColl.go`).<br>
The server will start and listen on a predefined port (e.g., 8084).<br>

## Test the endpoints.<br>
You can test the server's endpoints using curl. For example:  `curl http://localhost:8084`<br>  

## Monitor
If you have set up Prometheus and Grafana, you can monitor the server's performance and system metrics.<br>
The server exposes several endpoints for metrics collection.<br>



**Prometheus can natively collect various memory-related statistics, including those provided by the Go runtime.**  
 

The `go_memstats_alloc_bytes`  is one such native metric. It represents the current  number of bytes allocated  in the heap, <br>
which is updated at each GC cycle &  is similar to `mem.Alloc' present in our Go codebase .






