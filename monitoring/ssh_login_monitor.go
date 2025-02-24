package monitoring

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var lastOffset int64 = 0

func MonitorSshLogins() bool {
	logPath := "/var/log/auth.log"
	file, err := os.Open(logPath)
	if err != nil {
		log.Errorf("cannot open ssh log file: %v", err)
		return false
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Errorf("cannot get file info: %v", err)
		return false
	}

	currentOffset := fileInfo.Size()

	if lastOffset == 0 {
		lastOffset = currentOffset
		return false
	}

	if currentOffset > lastOffset {
		_, err = file.Seek(lastOffset, 0)
		if err != nil {
			log.Errorf("cannot seek file: %v", err)
			return false
		}

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			line = strings.TrimSpace(line)
			lastOffset += int64(len(line)) + 1

			if strings.Contains(line, "session opened") || strings.Contains(line, "Accepted publickey") {
				log.Info("New login detected!")
				return true
			}
		}
	}

	return false
}
