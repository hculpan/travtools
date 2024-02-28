package embed

import "embed"

//go:embed data/*
var embeddedFiles embed.FS

func ReadDataFile(filename string) ([]byte, error) {
	return embeddedFiles.ReadFile("data/" + filename)
}
