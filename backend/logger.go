package backend

import (
	"fmt"
	"os"
	"os/user"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type CustomLogger struct {
	logger   logger.Logger
	logFile  *os.File
	logMutex sync.Mutex
}

func NewCustomLogger(l logger.Logger, filePath string) (*CustomLogger, error) {
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &CustomLogger{
		logger:  l,
		logFile: logFile,
	}, nil
}

func (c *CustomLogger) logWithDetails(level string, msg string) {
	c.logMutex.Lock()
	defer c.logMutex.Unlock()

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	currentUser, err := user.Current()
	var username string
	if err != nil {
		username = "unknown"
	} else {
		username = currentUser.Username
	}

	logMessage := fmt.Sprintf("[%s] [%s] [User: %s] %s", currentTime, level, username, msg)

	switch level {
	case "INFO":
		c.logger.Info(logMessage)
	case "ERROR":
		c.logger.Error(logMessage)
	default:
		c.logger.Info(logMessage)
	}

	fmt.Fprintln(c.logFile, logMessage)
}

func (c *CustomLogger) Info(msg string) {
	c.logWithDetails("INFO", msg)
}

func (c *CustomLogger) Error(msg string) {
	c.logWithDetails("ERROR", msg)
}

func (c *CustomLogger) Close() error {
	return c.logFile.Close()
}
