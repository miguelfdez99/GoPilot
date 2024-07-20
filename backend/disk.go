package backend

import (
	"github.com/shirou/gopsutil/disk"
)

func (b *Backend) GetDiskUsage() (DiskUsage, error) {
	usageStat, err := disk.Usage("/")
	if err != nil {
		return DiskUsage{}, err
	}

	return DiskUsage{Total: usageStat.Total, Used: usageStat.Used}, nil
}
