package backend

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	User       string
	Pid        int32
	CpuPercent float64
	MemPercent float32
	VMS        uint64
	RSS        uint64
	TTY        string
	Status     string
	StartTime  int64
	Cmdline    string
}

func (b *Backend) GetProcessInfo() ([]ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	procInfos := make([]ProcessInfo, 0, len(procs))

	for _, proc := range procs {
		name, _ := proc.Username()
		cmdline, _ := proc.Cmdline()
		cpuPercent, _ := proc.CPUPercent()
		memPercent, _ := proc.MemoryPercent()
		vms, _ := proc.MemoryInfo()
		status, _ := proc.Status()
		startTime, _ := proc.CreateTime()
		tty, _ := proc.Terminal()

		procInfos = append(procInfos, ProcessInfo{
			User:       name,
			Pid:        proc.Pid,
			CpuPercent: cpuPercent,
			MemPercent: memPercent,
			VMS:        vms.VMS,
			RSS:        vms.RSS,
			TTY:        tty,
			Status:     status,
			StartTime:  startTime,
			Cmdline:    cmdline,
		})
	}

	return procInfos, nil
}
