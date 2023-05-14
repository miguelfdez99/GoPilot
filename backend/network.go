package backend

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

type NetworkUsage struct {
	BytesSent     float64 `json:"bytes_sent"`
	BytesReceived float64 `json:"bytes_received"`
}

var lastNetIOCounters *net.IOCountersStat

func (b *Backend) GetNetworkUsage() (*NetworkUsage, error) {
	netIOCounters, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	if len(netIOCounters) > 0 {
		if lastNetIOCounters != nil {
			bytesSent := float64(netIOCounters[0].BytesSent-lastNetIOCounters.BytesSent) / (1024 * 1024) * 10
			bytesReceived := float64(netIOCounters[0].BytesRecv-lastNetIOCounters.BytesRecv) / (1024 * 1024) * 10

			lastNetIOCounters = &netIOCounters[0]
			return &NetworkUsage{
				BytesSent:     bytesSent,
				BytesReceived: bytesReceived,
			}, nil
		}
		lastNetIOCounters = &netIOCounters[0]
	}

	return nil, fmt.Errorf("no network interface found")
}
