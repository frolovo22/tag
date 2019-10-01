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
	case TagVersionID3v22, TagVersionID3v23, TagVersionID3v24:
		return ReadID3v2(input)
	default:
		return nil, ErrorUnsupportedFormat
	}
}

func CheckVersion(input io.ReadSeeker) TagVersion {
	if checkID3v1(input) != TagVersionUndefined {
		return TagVersionID3v1
	}

	if version := checkID3v2(input); version != TagVersionUndefined {
		return version
	}

	return TagVersionUndefined
}
