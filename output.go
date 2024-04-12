package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"os"
)

func GenerateTXT(files []File) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, f := range files {
		_, err := writer.WriteString("- " + f.Path + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateCSV(files []File) error {
	file, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, f := range files {
		err := writer.Write([]string{f.Path})
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateJSON(files []File) error {
	file, err := os.Create("output.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(files)
	if err != nil {
		return err
	}

	return nil
}

func GenerateXML(files []File) error {
	file, err := os.Create("output.xml")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	err = encoder.Encode(files)
	if err != nil {
		return err
	}

	return nil
}
