// simple_csv is a a wrapper around the built-in csv package which adds some features such as reading from
// a start and end point. It makes some options such as changing the delimiter more clear e.g. it is called
// 'delimiter' instead of 'comma'

package simple_csv

import (
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"io"
	"fmt"
)

const (
	MAXUINT = ^uint(0)
)

// Holds csv reader info used when calling the Read() method. Only filePath is required. Other fields are optional.
type csvReader struct {
	filePath	string
	Start		int
	End			int
	MaxLines	int
	delimiter 	rune
	commentChar	rune
	Data		[][]string
}

// Constructor for creating a csvReader instance. Initialised only with the csv filePath to read. Other fields
// are set with setter methods.
func NewCSVReader(filePath string) (csvReader, error) {
	if _, err := os.Open(filePath); err != nil {
		return csvReader{}, fmt.Errorf("invalid file path: %v", err)
	}
	return csvReader{
		filePath: filePath,
		Start: -1,
		End: int(MAXUINT >> 1),
		MaxLines: int(MAXUINT >> 1),
		delimiter: ',',
		commentChar: 0}, nil
}

func (csvReader *csvReader) Delimiter(delim rune) error {
	if delim == rune('\n') || delim == rune('\r') {
		return fmt.Errorf("delimiter cannot be '\\n' or '\\r'")
	}
	csvReader.delimiter = delim
	return nil
}

func (csvReader *csvReader) CommentChar(commChar rune) error {
	if commChar == rune('\n') || commChar == rune('\r') || commChar == rune(',') {
		return fmt.Errorf("delimiter cannot be '\\n', '\\r' or ','")
	}
	csvReader.commentChar = commChar
	return nil
}

func (csvReader *csvReader) Read() error {

	f, err := os.Open(csvReader.filePath)
	if err != nil {
		return fmt.Errorf("unable to open csv file: %v", err)
	}

	reader := csv.NewReader(bufio.NewReader(f))
	reader.Comma = csvReader.delimiter
	reader.Comment = csvReader.commentChar

	if csvReader.Start > 0 {
		for i := 0; i < csvReader.Start; i++ {
			_, err := reader.Read()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	reader.FieldsPerRecord = 0
	for i := csvReader.Start; i < csvReader.End; i++ {
		line, err := reader.Read()
		if err == io.EOF || i == csvReader.MaxLines {
			break
		}
		if err != nil {
			return fmt.Errorf("failed reading csv lines at line %d. returning lines read up to line %d: %v", i, i, err)
		}
		csvReader.Data = append(csvReader.Data, line)
	}
	return nil
}