package parser

import (
	"path/filepath"
	"github.com/thimunri/logtest/handlers"
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strings"
)

type LogParser struct {
	logFileUser string
	LogPath string
}

func (l *LogParser) Init() error {

	logFiles, err := l.getLogFiles()

	for _, logFile := range logFiles {
		  l.normalizeLogByUser(logFile)
	}

	return err
}

func (l *LogParser) normalizeLogByUser(logPath string){

	reg, _ := regexp.Compile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)
	regDate, _ := regexp.Compile(`[0-9]{2}\/[a-zA-z]{3}\/[0-9]{4}`)
	file, err := os.Open(logPath)
	if err != nil {

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if reg.MatchString(scanner.Text()) {
			userId := reg.FindAllString(scanner.Text(), 1)[0]
			dateLog := regDate.FindAllString(scanner.Text(), 1)[0]
			dirLog := fmt.Sprintf("%s/%s", l.LogPath, strings.Replace(dateLog, "/", "", -1))
			handlers.CheckDir(dirLog)
			userFileLog := fmt.Sprintf("%s/%s.log", dirLog, userId)

			file, _ := os.OpenFile(userFileLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 666)

			fmt.Println(userFileLog)
			file.WriteString(scanner.Text())
		}
	}
}

// Scan directory for log files
func (l *LogParser) getLogFiles() ([]string, error){
	var logFiles []string
	err := filepath.Walk(l.LogPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			foundLogs,_ := filepath.Glob(fmt.Sprintf("%s/*.log", path))
			for _,value := range foundLogs {
				logFiles = append(logFiles, value)
			}
		}
		return nil
	})
	return logFiles, err
}