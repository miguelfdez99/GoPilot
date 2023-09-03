package backend

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gen2brain/beeep"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type SystemStatThresholds struct {
	CPU    float64 `json:"CPU"`
	Memory float64 `json:"Memory"`
	Disk   float64 `json:"Disk"`
}

func (b *Backend) MonitorFile(filepath string) {
	b.watchList.Store(filepath, "file")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				b.logger.Info(fmt.Sprint("event:", event))
				if event.Op&fsnotify.Write == fsnotify.Write {
					msg := fmt.Sprint("modified file:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					msg := fmt.Sprint("created file:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					msg := fmt.Sprint("removed file:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					msg := fmt.Sprint("renamed file:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				b.logger.Error(fmt.Sprint("error:", err))
			}
		}
	}()

	err = watcher.Add(filepath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func (b *Backend) MonitorDir(dirpath string) {
	b.watchList.Store(dirpath, "dir")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				b.logger.Info(fmt.Sprint("event:", event))
				if event.Op&fsnotify.Write == fsnotify.Write {
					msg := fmt.Sprint("modified file in directory:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					msg := fmt.Sprint("created file in directory:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					msg := fmt.Sprint("removed file in directory:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					msg := fmt.Sprint("renamed file in directory:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					msg := fmt.Sprint("changed file mode in directory:", event.Name)
					b.logger.Info(msg)
					err := beeep.Notify("File Monitor", msg, "assets/information.png")
					if err != nil {
						log.Fatal(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				b.logger.Error(fmt.Sprint("error:", err))
			}
		}
	}()

	err = watcher.Add(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func (b *Backend) WatchList() map[string]string {
	watchList := make(map[string]string)
	b.watchList.Range(func(key, value interface{}) bool {
		k, kOk := key.(string)
		v, vOk := value.(string)
		if kOk && vOk {
			watchList[k] = v
		}
		return true
	})
	return watchList
}

func (b *Backend) MonitorSystemStats(thresholds SystemStatThresholds) {

	memoryThresholdInBytes := uint64(thresholds.Memory * float64(1024*1024*1024))
	diskThresholdInBytes := uint64(thresholds.Disk * float64(1024*1024*1024))

	done := make(chan bool)
	go func() {
		for {
			v, _ := mem.VirtualMemory()
			d, _ := disk.Usage("/")

			if v.Used > memoryThresholdInBytes {
				err := beeep.Notify("System Monitor", "RAM usage is high!", "assets/warning.png")
				if err != nil {
					log.Fatal(err)
				}
			}
			if float64(d.Used) > float64(diskThresholdInBytes) {
				err := beeep.Notify("System Monitor", "Disk usage is high!", "assets/warning.png")
				if err != nil {
					log.Fatal(err)
				}
			}
			cpuUsage, _ := cpu.Percent(0, false)
			if cpuUsage[0] > thresholds.CPU {
				err := beeep.Notify("System Monitor", "CPU usage is high!", "assets/warning.png")
				if err != nil {
					log.Fatal(err)
				}
			}

			time.Sleep(60 * time.Second)
		}
	}()
	<-done
}
