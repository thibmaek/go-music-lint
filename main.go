package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path string `json:"path" xml:"path"`
}

type Flags struct {
	outputFormat string
}

func getCLIArgs() (string, Flags) {
	mainDir := os.Args[1]

	flags := Flags{}
	fs := flag.NewFlagSet("flags", flag.ExitOnError)
	fs.StringVar(&flags.outputFormat, "outputFormat", "", "Output format: csv, json, xml, txt")
	err := fs.Parse(os.Args[2:])

	if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	return mainDir, flags
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <directory>")
		return
	}

	dir, f := getCLIArgs()

	fmt.Println(dir, os.Args)

	var files []File

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".flac") {
			fmt.Println(path)
			files = append(files, File{Path: path})
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	switch f.outputFormat {
	case "csv":
		err = GenerateCSV(files)
	case "json":
		err = GenerateJSON(files)
	case "xml":
		err = GenerateXML(files)
	case "txt":
		err = GenerateTXT(files)
	default:
		return
	}

	if err != nil {
		fmt.Println("Error generating report:", err)
	}
}
