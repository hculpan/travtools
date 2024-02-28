package embed

import "embed"

// content holds our static web server content.
//
//go:embed assets/* templates/*
var embeddedFiles embed.FS

func ReadFile(filename string) ([]byte, error) {
	return embeddedFiles.ReadFile(filename)
}

func GetEmbeddedFS() *embed.FS {
	return &embeddedFiles
}
