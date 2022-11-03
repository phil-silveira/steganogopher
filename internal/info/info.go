package info

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
)

type ImageInfo struct {
	Name                    string
	DocumentType            string
	Width                   int
	Height                  int
	AvailableSpaceInBytes   int
	AvailableSpaceFormatted string
}

func GetImageInfo(imagePath string) (ImageInfo, error) {
	fd, err := os.Open(imagePath)
	defer fd.Close()

	if err != nil {
		log.Fatal(err)
		return ImageInfo{}, err
	}

	data, err := png.Decode(fd)

	if err != nil {
		log.Fatal(err)
		return ImageInfo{}, err
	}
	s := strings.Split(fd.Name(), "/")
	filename := s[len(s)-1]
	width := data.Bounds().Max.X
	height := data.Bounds().Max.Y

	availableSpaceInBytes := width * height * 3 / 8

	return ImageInfo{
		Name:                    filename,
		DocumentType:            strings.ToUpper(strings.Split(filename, ".")[1]),
		Width:                   width,
		Height:                  height,
		AvailableSpaceInBytes:   availableSpaceInBytes,
		AvailableSpaceFormatted: formatSpace(availableSpaceInBytes),
	}, nil
}

func formatSpace(spaceInBytes int) string {
	if spaceInBytes > 1024 {
		return fmt.Sprintf("%d KB", spaceInBytes/1024)
	} else {
		return fmt.Sprintf("%d B", spaceInBytes)

	}
}
