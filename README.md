Lab - Speedtest

Brief Description: Build a Docker container that can run a speedtest to Cloudflare. 

Prerequisites: GoLang, Docker installed

Installation: 
1. Download and build the Dockerfile.
1. Run the container.

Functionality/Features:  Runs speedtest to check Internet connectivity.  
Uses net/http/httptrace to determine DNS and connection times.  
Calculates average latency, tcp jitter, and dns times in milliseconds.
Prints to stdout in Prometheus format so your exporter can pick it up.

Outputs:
cf_start_timestamp 
cf_latency_ms 
cf_latency_ms 
cf_dnslookup_ms 

Technologies Used:
GoLang
Docker
