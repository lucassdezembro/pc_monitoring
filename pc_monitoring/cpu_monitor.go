package pc_monitoring

import (
	"github.com/mackerelio/go-osstat/cpu"
)

type CpuMonitor interface {
	GetCpuUsage() (float64, error)
}

type cpuMonitor struct{}

func NewCpuMonitor() CpuMonitor {
	return &cpuMonitor{}
}

func (c *cpuMonitor) GetCpuUsage() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.User), nil
}
