package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

// ScanPort checks if a given port on a hostname is open or closed.
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "closed"
		return result
	}

	defer conn.Close()

	result.State = "open"
	return result
}

// IniteleScan performs a basic scan of ports from 1 to 1024 on the given hostname.
func IniteleScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
		results = append(results, ScanPort("http", hostname, i))
	}

	return results
}
