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
		return ReadID3v1(input)
	case TagVersionID3v22:
	case TagVersionID3v23:
	case TagVersionID3v24:
		return ReadID3v24(input)
	default:
		return nil, ErrorUnsupportedFormat
	}
	return nil, ErrorUnsupportedFormat
}

func CheckVersion(input io.ReadSeeker) TagVersion {
	if checkID3v1(input) != TagVersionUndefined {
		return TagVersionID3v1
	}

	if checkID3v24(input) != TagVersionUndefined {
		return TagVersionID3v24
	}

	return TagVersionUndefined
}
