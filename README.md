# GC_metrics_in_Prometheus
Monitoring  the Garbage collection  of a basic http server collected via Prometheus 


This repository contains a simple Go web server application.<br>
It demonstrates basic HTTP request handling in Go and includes features to monitor system metrics like garbage collection.

#Features
  - Basic HTTP server implementation in Go.
  - Endpoints for testing load and monitoring server responses.
  - Integration examples for system metrics monitoring (garbage collection, memory usage, disk I/O).
&nsbsp;

    
Getting Started
Prerequisites<br>
Go (version 1.x or later)<br>
curl (for testing endpoints)
Monitoring tools :  Prometheus and Grafana
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/go-web-server-example.git
cd go-web-server-example
Running the Server
To run the server, use the following command:

bash
Copy code
go run main.go
The server will start and listen on a predefined port (e.g., 8084).

Testing the Endpoints
You can test the server's endpoints using curl. For example:

curl http://localhost:8084
Monitoring
If you have set up Prometheus and Grafana, you can monitor the server's performance and system metrics. The server exposes several endpoints for metrics collection.


**Prometheus can natively collect various memory-related statistics, including those provided by the Go runtime.**<br>

The `go_memstats_alloc_bytes`  is one such metric. It represents the current  number of bytes allocated  in the heap , which is updated at each GC cycle.<br>
which is similar to `mem.Alloc' present in our Go codebase .






