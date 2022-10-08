package main

import (
	"fmt"

	"github.com/lucassdezembro/pc_monitoring/pc_monitoring"
)

func main() {
	pcMonitor := pc_monitoring.NewCpuMonitor()

	fmt.Print(pcMonitor.GetCpuUsage())
}
