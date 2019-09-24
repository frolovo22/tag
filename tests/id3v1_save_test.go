package tests

import (
	"bytes"
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestId3v1Save(t *testing.T) {
	asrt := assert.New(t)

	id3v1, err := tag.ReadFile("tests/id3v1.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	err = id3v1.SaveFile("tests/id3v1.save.mp3")
	asrt.NoError(err, "save")
	if err != nil {
		return
	}

	cmp := compareFiles("tests/id3v1.mp3", "tests/id3v1.save.mp3")
	asrt.True(cmp)

	err = os.Remove("tests/id3v1.save.mp3")
	asrt.NoError(err, "remove")
	if err != nil {
		return
	}
}

func compareFiles(path1, path2 string) bool {
	data1, err := ioutil.ReadFile(path1)
	if err != nil {
		return false
	}

	data2, err := ioutil.ReadFile(path2)
	if err != nil {
		return false
	}

	result := bytes.Compare(data1, data2)
	return result == 0
}
