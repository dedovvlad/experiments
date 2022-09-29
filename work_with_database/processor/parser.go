package processor

import (
	"encoding/csv"
	"fmt"
	"os"
)

type (
	ID struct {
		Series string
		Number string
	}
)

const (
	series = iota
	number
)

func Parser(filePath string, from, to int) ([]*ID, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, len(reader), err
	}

	switch {
	case from < 0:
		from = 0
	case to <= 0:
		to = len(reader)
	case from > len(reader):
		return []*ID{}, len(reader) - 1, fmt.Errorf("'from' isn't correct")
	}

	result := make([]*ID, 0, len(reader))

	for _, row := range reader[1+from : to] {
		result = append(result, &ID{
			Series: row[series],
			Number: row[number],
		})
	}
	return result, len(reader) - 1, nil
}
