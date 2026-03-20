package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"codeberg.org/jakobdev/asar-go"
)

func main() {
	desktopFilename := flag.String("desktop-filename", os.Getenv("FLATPAK_ID"), "Set a different desktop filename")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: patch-desktop-filename: <path>")
		os.Exit(1)
	}

	archivePath := flag.Arg(0)
	archive, err := asar.OpenFile(archivePath, &asar.ReadOptions{UnpackedDir: fmt.Sprintf("%s.unpacked", archivePath)})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer archive.Close()

	entry, err := archive.GetEntry("/package.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader, err := entry.GetReader()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	packageMap := make(map[string]any)
	err = json.Unmarshal(content, &packageMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	packageMap["desktopName"] = fmt.Sprintf("%s.desktop", *desktopFilename)

	packageJson, err := json.Marshal(packageMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = archive.CreateFileEntry("/package.json", packageJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = archive.Save(&asar.WriteOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
