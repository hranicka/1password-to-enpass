package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
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
	fr, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fr)

	fw, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(fw)
	defer w.Flush()

	var m = &Map{}
	for i := 0; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i == 0 {
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

			err = w.Write([]string{
				"Title",
				"Username",
				"Email",
				"Password",
				"Website",
				"Note",
			})
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		err = w.Write([]string{
			row[m.Title],
			row[m.Username],
			row[m.Email],
			row[m.Password],
			row[m.Website],
			row[m.Note],
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
