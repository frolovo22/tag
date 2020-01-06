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
		return ReadID3v22(input)
	case TagVersionID3v23:
		return ReadID3v23(input)
	case TagVersionID3v24:
		return ReadID3v24(input)
	case TagVersionMP4:
		return ReadMp4(input)
	case TagVersionFLAC:
		return ReadFLAC(input)
	default:
		return nil, ErrorUnsupportedFormat
	}
	return nil, ErrorUnsupportedFormat
}

func CheckVersion(input io.ReadSeeker) TagVersion {
	if checkID3v24(input) {
		return TagVersionID3v24
	}

	if checkID3v23(input) {
		return TagVersionID3v23
	}

	if checkID3v22(input) {
		return TagVersionID3v22
	}

	if checkID3v1(input) {
		return TagVersionID3v1
	}

	if checkMp4(input) {
		return TagVersionMP4
	}

	if checkFLAC(input) {
		return TagVersionFLAC
	}

	return TagVersionUndefined
}
