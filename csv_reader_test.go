package simple_csv

import "testing"

func TestDelimiter(t *testing.T) {
	reader, err := NewCSVReader("seeds.csv")
	if err != nil {
		t.Errorf("unable to create csvReader: %v", err)
	}

	reader.Delimiter('\t')
	if err = reader.Read(); err != nil {
		t.Errorf("unable to read csv file: %v", err)
	}
	expectedLines := 210
	if len(reader.Data) != expectedLines {
		t.Errorf("expected lines: %d, found: %d", expectedLines, len(reader.Data))
	}
}

func TestReadStartEnd(t *testing.T) {
	reader, _ := NewCSVReader("simple.csv")
	reader.Start = 2
	reader.End = 7

	reader.Read()
	expectedLines := reader.End - reader.Start
	if len(reader.Data) != expectedLines {
		t.Errorf("expected lines: %d, found: %d", expectedLines, len(reader.Data))
	}
}