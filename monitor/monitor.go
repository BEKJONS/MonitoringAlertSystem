package monitor

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"runtime" // Provides functions to interact with the Go runtime environment
)

// Metrics struct holds the collected data for CPU, memory, and disk I/O usage
type Metrics struct {
	CPUUsage    float64 // Percentage of CPU used
	MemoryUsage float64 // Percentage of memory used
	DiskIO      float64 // Disk I/O in MB/s
}

// CollectMetrics function gathers system metrics by calling helper functions
func CollectMetrics() Metrics {
	// Collect CPU usage, memory usage, and disk I/O using helper functions
	cpuUsage := getCPUUsage()
	memoryUsage := getMemoryUsage()
	diskIO := getDiskIO()

	// Return a Metrics struct with the collected values
	return Metrics{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		DiskIO:      diskIO,
	}
}

// getCPUUsage function calculates CPU usage (dummy implementation)
func getCPUUsage() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		panic(err)
	}
	return percent[0]
}

// getMemoryUsage function calculates memory usage using Go runtime's memory statistics
func getMemoryUsage() float64 {
	var m runtime.MemStats   // MemStats is a struct that holds runtime memory statistics
	runtime.ReadMemStats(&m) // Read current memory stats into m
	// Example: convert bytes to percentage of system memory (this is a simplified approach)
	return float64(m.Alloc) / float64(m.Sys) * 100
}

// getDiskIO function calculates disk I/O usage (dummy implementation)
func getDiskIO() float64 {
	// Get disk I/O stats
	io, err := disk.IOCounters()
	if err != nil {
		fmt.Println(err)
		return 0.0
	}

	// Calculate disk I/O usage
	var totalReadBytes uint64
	var totalWriteBytes uint64
	for _, v := range io {
		totalReadBytes += v.ReadBytes
		totalWriteBytes += v.WriteBytes
	}

	// Calculate disk I/O usage in percentage
	totalBytes := totalReadBytes + totalWriteBytes
	if totalBytes == 0 {
		return 0.0
	}
	diskIOUsage := float64(totalBytes) / (1024 * 1024) // Convert to MB

	return diskIOUsage
}
