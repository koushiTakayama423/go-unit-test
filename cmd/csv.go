package cmd

import (
	"encoding/csv"
	"log"
	"os"
)

func Input_csv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Input_csv err→%v", err)
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Printf("Input_csv err→%v", err)
		return nil, err
	}
	return rows, nil
}

func Output_csv(records [][]string) error {
	file, err := os.Create("output/output.csv")
	if err != nil {
		log.Printf("Output_csv err→%v", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		var container []string
		for i, r := range record {
			if i == 1 {
				txt, err := Encrypt(r)
				if err != nil {
					log.Printf("Output_csv err→%v", err)
					return err
				}
				container = append(container, txt)
			} else {
				container = append(container, r)
			}
		}
		writer.Write(container)
	}

	return nil
}
