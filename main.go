package main

import (
	"fmt"
	"os"
	"time"
)

const
(
	errorLogType string = "ERROR"
	warningLogType string = "WARNING"
	infoLogType string = "INFO"
)

func getLogCurrentDayFileName() string {
	year, month, day := time.Now().Date()
	yearString, monthString, dayString := formatNumberStringWithLeadingZero(year), formatNumberStringWithLeadingZero(int(month)), formatNumberStringWithLeadingZero(day)
	return fmt.Sprintf("%v-%v-%v.logGo", yearString, monthString, dayString)
}

func createLogFileAndReturnWriter() (*os.File, error) {
	w, err := os.Create(fmt.Sprintf("%v", getLogCurrentDayFileName()))

	if err != nil {
		return nil, err
	}

	return w, nil
}

func getFormattedTimeString(tm *time.Time) string {
	hourString := formatNumberStringWithLeadingZero(tm.Hour())
	minuteString := formatNumberStringWithLeadingZero(tm.Minute())
	secondString := formatNumberStringWithLeadingZero(tm.Second())

	return fmt.Sprintf("%s:%s:%s", hourString, minuteString, secondString)
}

func formatNumberStringWithLeadingZero(number int) string {
	if number < 10 {
		return fmt.Sprintf("%02d", number)
	}

	return fmt.Sprint(number)
}

func LogError(logMessage *string) {
	writeLogMessage(errorLogType, logMessage)
}

func LogWarning(logMessage *string) {
	writeLogMessage(warningLogType, logMessage)
}

func LogInfo(logMessage *string) {
	writeLogMessage(infoLogType, logMessage)
}

func writeLogMessage(logType string, logMessage *string) {
	fileName := getLogCurrentDayFileName()
	w, err := os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
	if err != nil {
		w, err = createLogFileAndReturnWriter()
		if err != nil {
			fmt.Printf("Program failed to create a new log file! Error: %s\n", err)
		}
	}

	currentTime := time.Now()
	w.WriteString(fmt.Sprintf("[%v] - %v: %v\n", logType, getFormattedTimeString(&currentTime), *logMessage))
}