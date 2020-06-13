package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

var (
	inputName  = "input.csv"
	outputName = "output.csv"
)

type Map struct {
	Title    int
	Username int
	Email    int
	Password int
	Website  int
	Note     int
}

func main() {
	// Open input CSV file for reading
	fr, err := os.Open(inputName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fr)

	// Open output CSV file for writing
	fw, err := os.Create(outputName)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(fw)
	defer w.Flush()

	// Actual conversion
	var m = &Map{}
	for i := 0; ; i++ {
		// Read row by row until EOF
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Init data from header
		if i == 0 {
			initMap(m, row)
			if err := writeHeader(w); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// Convert current row
		if err := writeRow(w, m, row); err != nil {
			log.Fatal(err)
		}
	}
}

func initMap(m *Map, row []string) {
	for k, v := range row {
		switch v {
		case "TITLE":
			m.Title = k
		case "USERNAME":
			m.Username = k
		case "PASSWORD":
			m.Password = k
		case "URL":
			m.Website = k
		case "EMAIL":
			m.Email = k
		case "NOTES":
			m.Note = k
		}
	}
}

func writeHeader(w *csv.Writer) error {
	return w.Write([]string{
		"Title",
		"Username",
		"Email",
		"Password",
		"Website",
		"Note",
	})
}

func writeRow(w *csv.Writer, m *Map, row []string) error {
	return w.Write([]string{
		row[m.Title],
		row[m.Username],
		row[m.Email],
		row[m.Password],
		row[m.Website],
		row[m.Note],
	})
}
