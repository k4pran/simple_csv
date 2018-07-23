package csv_simple

import "testing"

func TestDelimiter(t *testing.T) {
	reader, err := NewCSVReader("seeds.csv")
	if err != nil {
		t.Errorf("unable to create csvReader: %v", err)
	}

	reader.SetDelimiter('\t')
	if err = reader.Read(); err != nil {
		t.Errorf("unable to read csv file: %v", err)
	}
	expectedLines := 210
	if len(reader.Data) != expectedLines {
		t.Errorf("expected lines: %d, found: %d", expectedLines, len(reader.Data))
	}
}
