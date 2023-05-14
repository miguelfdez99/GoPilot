package backend

import (
	"github.com/shirou/gopsutil/disk"
)

type DiskUsage struct {
	Total uint64
	Used  uint64
}

func (b *Backend) GetDiskUsage() (DiskUsage, error) {
	usageStat, err := disk.Usage("/")
	if err != nil {
		return DiskUsage{}, err
	}

	return DiskUsage{Total: usageStat.Total, Used: usageStat.Used}, nil
}
