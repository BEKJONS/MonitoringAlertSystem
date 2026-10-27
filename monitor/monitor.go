package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
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

// getCPUUsage function returns overall CPU usage as a percentage
func getCPUUsage() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		panic(err)
	}
	return percent[0]
}

// getMemoryUsage function returns the percentage of system memory in use
func getMemoryUsage() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
		return 0.0
	}
	return v.UsedPercent
}

// Disk counters from the previous sample, used to turn cumulative bytes into a rate
var (
	lastDiskBytes uint64
	lastDiskTime  time.Time
)

// getDiskIO function returns disk read+write throughput in MB/s since the previous sample
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

	// The counters are cumulative since boot, so compare with the previous sample
	totalBytes := totalReadBytes + totalWriteBytes
	now := time.Now()
	prevBytes, prevTime := lastDiskBytes, lastDiskTime
	lastDiskBytes, lastDiskTime = totalBytes, now
	if prevTime.IsZero() || totalBytes < prevBytes {
		return 0.0 // first sample, or the counters were reset
	}

	return float64(totalBytes-prevBytes) / (1024 * 1024) / now.Sub(prevTime).Seconds()
}
