package fileProcessor

import (
	"encoding/json"
	"fmt"
	"os"
	events "user_events/Model"
)

func ReadJSON(filePath string) ([]events.Event, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading input file: %w", err)
	}

	var events []events.Event
	err = json.Unmarshal(file, &events)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return events, nil
}

func WriteJSON(filePath string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}

	return nil
}
