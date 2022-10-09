package pc_monitoring

import (
	"github.com/mackerelio/go-osstat/cpu"
)

type CpuMonitor interface {
	GetCpusNumber() (int, error)
	GetCpuGuest() (float64, error)
	GetCpuGuestNice() (float64, error)
	GetCpuIdle() (float64, error)
	GetCpuIowait() (float64, error)
	GetCpuIrq() (float64, error)
	GetCpuNice() (float64, error)
	GetCpuSoftirq() (float64, error)
	GetCpuStatCont() (float64, error)
	GetCpuSteal() (float64, error)
	GetCpuSystem() (float64, error)
	GetCpuUserUsage() (float64, error)
	GetCpuTotalUsage() (float64, error)
}

type cpuMonitor struct{}

func NewCpuMonitor() CpuMonitor {
	return &cpuMonitor{}
}

func (c *cpuMonitor) GetCpusNumber() (int, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return int(stat.Guest), nil
}

func (c *cpuMonitor) GetCpuGuest() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Guest), nil
}

func (c *cpuMonitor) GetCpuGuestNice() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.GuestNice), nil
}

func (c *cpuMonitor) GetCpuIdle() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Idle), nil
}

func (c *cpuMonitor) GetCpuIowait() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Iowait), nil
}

func (c *cpuMonitor) GetCpuIrq() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Irq), nil
}

func (c *cpuMonitor) GetCpuNice() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Nice), nil
}

func (c *cpuMonitor) GetCpuSoftirq() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Softirq), nil
}

func (c *cpuMonitor) GetCpuStatCont() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.StatCount), nil
}

func (c *cpuMonitor) GetCpuSteal() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Steal), nil
}

func (c *cpuMonitor) GetCpuSystem() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.System), nil
}

func (c *cpuMonitor) GetCpuUserUsage() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.User), nil
}

func (c *cpuMonitor) GetCpuTotalUsage() (float64, error) {
	stat, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	return float64(stat.Total), nil
}
