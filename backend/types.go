package backend

type BackupOptions struct {
	SourceDir string
	DestDir   string
}

type CronJob struct {
	Schedule string
	Command  string
}

type DiskUsage struct {
	Total uint64
	Used  uint64
}

type MemoryUsage struct {
	RAM  float64 `json:"ram"`
	Swap float64 `json:"swap"`
}

type NetworkUsage struct {
	BytesSent     float64 `json:"bytes_sent"`
	BytesReceived float64 `json:"bytes_received"`
}

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

type RunningService struct {
	Name        string
	ActiveState string
	Description string
}

type Service struct {
	Name          string
	StartupStatus string
}

type SystemInfo struct {
	OsName          string `json:"osName"`
	KernelVer       string `json:"kernelVer"`
	Uptime          string `json:"uptime"`
	Hostname        string `json:"hostname"`
	DesktopEnv      string `json:"desktopEnv"`
	CurrentUsername string `json:"currentUsername"`
	Memory          string `json:"memory"`
}

type SystemStatThresholds struct {
	CPU    float64 `json:"CPU"`
	Memory float64 `json:"Memory"`
	Disk   float64 `json:"Disk"`
}

type User struct {
	Username      string
	Password      string
	UID           int
	GID           int
	HomeDirectory string
	Shell         string
	Groups        []string
}
