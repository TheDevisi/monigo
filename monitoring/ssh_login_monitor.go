package monitoring

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func monitorSshLogins() {
	/*
		Basically we reading a log file, and storing it's original version.
		When someone is trying to connect or already connect we send a message and updating this file.
		In Ubuntu or Debian based distros everything is storing in /var/log/auth.log
	*/

	logPath := "/var/log/auth" // We're using this var because RedHat-based distros using another path
	// But for now, only ubuntu support
	authentificationLog, err := os.Open(logPath)

	// The "coolest" go's error handling xD
	if err != nil {
		log.Error("cannot start ssh monitoring, error: ", err)
	}

	// Create scanner

	scanner := bufio.NewScanner(authentificationLog)

	for {
		scanner.Scan()
		line := scanner.Text()

		if strings.Contains(line, "session oppened") || strings.Contains(line, "accepted publickey") {
			fmt.Println("pizda!!!")
			continue

		}

	}

}
