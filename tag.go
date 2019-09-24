package tag

import (
	"io"
	"os"
)

func ReadFile(path string) (Metadata, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Read(file)
}

func Read(input io.ReadSeeker) (Metadata, error) {
	version := CheckVersion(input)
	switch version {
	case TagVersionID3v1:
		return ReadID3v1Tags(input)
	default:
		return nil, ErrorUnsupportedFormat
	}
}

func CheckVersion(input io.ReadSeeker) TagVersion {
	if IsID3v1(input) {
		return TagVersionID3v1
	}
	return TagVersionUndefined
}
