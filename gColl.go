/*
The underlying objective here is to observe and analyze the Garbage collection & 
memory metrics of a simple app that is a basic operation for a http server writing "Hello World" .
It is mainly used as a trigger to update & observe the GC metrics.
Each time the HTTP endpoint is accessed, it not only responds with 'Hello World' but also invokes 
the `recordMetrics` function.
This function updates the Prometheus metrics with the latest memory and GC statistics from the Go runtime.
*/


package main

import (
	"fmt"
	"runtime"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"          // providers Go client library for defining and collecting custom metrics.
	"github.com/prometheus/client_golang/prometheus/promauto" // provides a way to create metrics that are automatically registered by prometheus.
	"github.com/prometheus/client_golang/prometheus/promhttp" // provides HTTP handlers to expose Prometheus metrics.
)




var (
	memAlloc = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_mem_alloc_bytes", 
		Help: "Number of bytes allocated and still in use.",
	})
	
	totalAlloc = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_total_alloc_bytes", 
		Help: "Total number of bytes allocated, even if freed.", 
	})
		
	heapAlloc = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_heap_alloc_bytes", 
		Help: "Number of heap bytes allocated & still in use . ",
	})

	numGC = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_gc_cycles_total", 
		Help: "Total number of completed GC cycles . " ,
	}) 

)



func recordMetrics() { 
	 var mem runtime.MemStats
	 runtime.ReadMemStats(&mem) 

	memAlloc.Set(float64(mem.Alloc))
	totalAlloc.Set(float64(mem.TotalAlloc))
	heapAlloc.Set(float64(mem.HeapAlloc))
	numGC.Set(float64(mem.NumGC)) 
}




func main() {
		recordMetrics()

		http.Handle("/metrics", promhttp.Handler()) 
		http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
			recordMetrics() // Update metrics on each http request.
			w.Write([]byte("Hello World\n"))
		})
// Start the HTTP server on port 8084 and check for errors
fmt.Println("Starting server on port 8084...")

err := http.ListenAndServe(":8084", nil)
    if err != nil {
        // If there is an error, print it and exit the program
        fmt.Printf("Failed to start server: %v\n", err)
        os.Exit(1) // Exit the program with a non-zero exit code
    }
}