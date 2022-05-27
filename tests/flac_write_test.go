package tests

import (
	"bytes"
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFLACWrite(t *testing.T) {
	asrt := assert.New(t)

	flac, err := tag.ReadFile("BeeMoved.flac")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	out, err := ioutil.TempFile("", "flacTst.flac")
	asrt.NoError(err)
	defer os.Remove(out.Name())
	err = flac.SaveFile(out.Name())
	asrt.NoError(err)
	flac2, err := tag.ReadFile(out.Name())
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	//Can't binary compare files as tag order is not deterministic

	//Check file main data is untouched
	result := bytes.Compare(flac.GetFileData(), flac2.GetFileData())
	asrt.Equal(result, 0)
	//Compare metadata is 1:1
	if value, err := flac.GetTitle(); err == nil {
		value2, err := flac2.GetTitle()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetArtist(); err == nil {
		value2, err := flac2.GetArtist()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetAlbum(); err == nil {
		value2, err := flac2.GetAlbum()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetYear(); err == nil {
		value2, err := flac2.GetYear()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetComment(); err == nil {
		value2, err := flac2.GetComment()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetGenre(); err == nil {
		value2, err := flac2.GetGenre()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetAlbumArtist(); err == nil {
		value2, err := flac2.GetAlbumArtist()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetDate(); err == nil {
		value2, err := flac2.GetDate()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetArranger(); err == nil {
		value2, err := flac2.GetArranger()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetAuthor(); err == nil {
		value2, err := flac2.GetAuthor()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetBPM(); err == nil {
		value2, err := flac2.GetBPM()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetCatalogNumber(); err == nil {
		value2, err := flac2.GetCatalogNumber()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetCompilation(); err == nil {
		value2, err := flac2.GetCompilation()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetComposer(); err == nil {
		value2, err := flac2.GetComposer()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetConductor(); err == nil {
		value2, err := flac2.GetConductor()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetCopyright(); err == nil {
		value2, err := flac2.GetCopyright()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, err := flac.GetDescription(); err == nil {
		value2, err := flac2.GetDescription()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, valuea, err := flac.GetDiscNumber(); err == nil {
		value2, valueb, err := flac2.GetDiscNumber()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
		asrt.Equal(valuea, valueb)
	}
	if value, err := flac.GetEncodedBy(); err == nil {
		value2, err := flac2.GetEncodedBy()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
	if value, valuea, err := flac.GetTrackNumber(); err == nil {
		value2, valueb, err := flac2.GetTrackNumber()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
		asrt.Equal(valuea, valueb)
	}
	if value, err := flac.GetPicture(); err == nil {
		value2, err := flac2.GetPicture()
		asrt.NoError(err)
		if err != nil {
			return
		}
		asrt.Equal(value, value2)
	}
}
