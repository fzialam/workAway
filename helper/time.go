package helper

import "time"

func ConvertSQLTimeToHTML(convert string) string {
	layout := "2006-01-02"
	inputLayout := "2006-01-02T15:04:05-07:00"
	parsedTime, _ := time.Parse(inputLayout, convert)
	convert = parsedTime.Format(layout)

	return convert
}

func TimeNowToString() string {
	layout := "2006/01/02 15:04:00"
	inputLayout := time.Now()

	return inputLayout.Format(layout)
}

func ConvertSQLTimeStamp(convert string) string {
	layout := "2006-01-02 15:04:00"
	inputLayout := "2006-01-02T15:04:05-07:00"
	parsedTime, _ := time.Parse(inputLayout, convert)
	convert = parsedTime.Format(layout)

	return convert
}
