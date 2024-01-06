package aggregate

import (
	"log"
	"time"
	fileprocessor "user_events/FileProcessor"
)

func AggregateEvents(inputFile, outputFile string) {
	events, err := fileprocessor.ReadJSON(inputFile)
	if err != nil {
		log.Fatal("Error reading input file:", err)
	}

	dailySummary := make(map[int]map[string]map[string]int)

	for _, event := range events {
		userID := event.UserID
		eventType := event.EventType
		timestamp := event.Timestamp
		date := TimestampToDate(timestamp)

		if dailySummary[userID] == nil {
			dailySummary[userID] = make(map[string]map[string]int)
		}

		if dailySummary[userID][date] == nil {
			dailySummary[userID][date] = make(map[string]int)
		}

		dailySummary[userID][date][eventType]++
	}

	err = fileprocessor.WriteJSON(outputFile, FormatOutput(dailySummary))
	if err != nil {
		log.Fatal("Error writing to output file:", err)
	}
}

func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).UTC().Format("2006-01-02")
}

func FormatOutput(dailySummary map[int]map[string]map[string]int) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	for userID, userData := range dailySummary {
		for date, eventData := range userData {
			summary := map[string]interface{}{"userId": userID, "date": date}
			for eventType, count := range eventData {
				summary[eventType] = count
			}
			result = append(result, summary)
		}
	}

	return result
}
