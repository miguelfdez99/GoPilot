package backend

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gen2brain/beeep"
)

func (b *Backend) MonitorFile(filepath string) {
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
					b.logger.Info(fmt.Sprint("modified file in directory:", event.Name))
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					b.logger.Info(fmt.Sprint("created file in directory:", event.Name))
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					b.logger.Info(fmt.Sprint("removed file in directory:", event.Name))
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					b.logger.Info(fmt.Sprint("renamed file in directory:", event.Name))
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					b.logger.Info(fmt.Sprint("changed file permissions in directory:", event.Name))
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
