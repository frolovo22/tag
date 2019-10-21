package tag

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strconv"
	"time"
)

const MP4_MARKER = "ftyp"
const MP4_MOOV_ATOM = "moov"
const MP4_META_ATOM = "meta"
const MP4_META_UPTA = "udta"
const MP4_META_ILST = "ilst"

const MP4_TAG_ALBUM = "album"
const MP4_TAG_ARTIST = "artist"
const MP4_TAG_ALBUM_ARTIST = "album_artist"
const MP4_TAG_YEAR = "year"
const MP4_TAG_TITLE = "title"
const MP4_TAG_GENRE = "genre"
const MP4_TAG_TRACK = "track"
const MP4_TAG_COMPOSER = "composer"
const MP4_TAG_ENCODER = "encoder"
const MP4_TAG_COPYRIGHT = "copyright"
const MP4_TAG_PICTURE = "picture"
const MP4_TAG_GROUPING = "grouping"
const MP4_TAG_KEYWORD = "keyword"
const MP4_TAG_LYRICS = "lyrics"
const MP4_TAG_COMMENT = "comment"
const MP4_TAG_TEMPO = "tempo"
const MP4_TAG_COMPILATION = "compilation"
const MP4_TAG_DISC = "disk"

var MP4_TYPES = [...]string{
	"mp41",
	"mp42",
	"isom",
	"iso2",
	"M4A ",
	"M4B ",
}

var atoms = map[string]string{
	"\xa9alb": MP4_TAG_ALBUM,
	"\xa9art": MP4_TAG_ARTIST,
	"\xa9ART": MP4_TAG_ARTIST,
	"aART":    MP4_TAG_ALBUM_ARTIST,
	"\xa9day": MP4_TAG_YEAR,
	"\xa9nam": MP4_TAG_TITLE,
	"\xa9gen": MP4_TAG_GENRE,
	"trkn":    MP4_TAG_TRACK,
	"\xa9wrt": MP4_TAG_COMPOSER,
	"\xa9too": MP4_TAG_ENCODER,
	"cprt":    MP4_TAG_COPYRIGHT,
	"covr":    MP4_TAG_PICTURE,
	"\xa9grp": MP4_TAG_GROUPING,
	"keyw":    MP4_TAG_KEYWORD,
	"\xa9lyr": MP4_TAG_LYRICS,
	"\xa9cmt": MP4_TAG_COMMENT,
	"tmpo":    MP4_TAG_TEMPO,
	"cpil":    MP4_TAG_COMPILATION,
	"disk":    MP4_TAG_DISC,
}

type MP4 struct {
	data map[string]interface{}
}

func (MP4) GetAllTagNames() []string {
	panic("implement me")
}

func (mp4 *MP4) GetVersion() TagVersion {
	return TagVersionMP4
}

func (MP4) GetFileData() []byte {
	panic("implement me")
}

func (mp4 *MP4) GetTitle() (string, error) {
	return mp4.getString(MP4_TAG_TITLE)
}

func (mp4 *MP4) GetArtist() (string, error) {
	return mp4.getString(MP4_TAG_ARTIST)
}

func (mp4 *MP4) GetAlbum() (string, error) {
	return mp4.getString(MP4_TAG_ALBUM)
}

func (mp4 *MP4) GetYear() (int, error) {
	year, err := mp4.getString(MP4_TAG_YEAR)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(year)
}

func (MP4) GetComment() (string, error) {
	panic("implement me")
}

func (mp4 *MP4) GetGenre() (string, error) {
	return mp4.getString(MP4_TAG_GENRE)
}

func (mp4 *MP4) GetAlbumArtist() (string, error) {
	return mp4.getString(MP4_TAG_ALBUM_ARTIST)
}

func (MP4) GetDate() (time.Time, error) {
	panic("implement me")
}

func (MP4) GetArranger() (string, error) {
	panic("implement me")
}

func (MP4) GetAuthor() (string, error) {
	panic("implement me")
}

func (MP4) GetBMP() (int, error) {
	panic("implement me")
}

func (MP4) GetCatalogNumber() (string, error) {
	panic("implement me")
}

func (MP4) GetCompilation() (string, error) {
	panic("implement me")
}

func (mp4 *MP4) GetComposer() (string, error) {
	return mp4.getString(MP4_TAG_COMPOSER)
}

func (MP4) GetConductor() (string, error) {
	panic("implement me")
}

func (mp4 *MP4) GetCopyright() (string, error) {
	return mp4.getString(MP4_TAG_COPYRIGHT)
}

func (MP4) GetDescription() (string, error) {
	panic("implement me")
}

func (MP4) GetDiscNumber() (int, int, error) {
	panic("implement me")
}

func (mp4 *MP4) GetEncodedBy() (string, error) {
	return mp4.getString(MP4_TAG_ENCODER)
}

func (mp4 *MP4) GetTrackNumber() (int, int, error) {
	track, err := mp4.getInt(MP4_TAG_TRACK)
	if err != nil {
		return 0, 0, err
	}
	total, err2 := mp4.getInt(MP4_TAG_TRACK + "_TOTAL")
	if err2 != nil {
		return 0, 0, err2
	}
	return track, total, nil
}

func (mp4 *MP4) GetPicture() (image.Image, error) {
	pictureBlock, ok := mp4.data[MP4_TAG_PICTURE]
	if !ok {
		return nil, ErrorTagNotFound
	}
	picture := pictureBlock.(AttachedPicture)
	switch picture.MIME {
	case "image/jpeg":
		return jpeg.Decode(bytes.NewReader(picture.Data))
	case "image/png":
		return png.Decode(bytes.NewReader(picture.Data))
	}

	return nil, ErrorIncorrectTag
}

func (MP4) SetTitle(title string) error {
	panic("implement me")
}

func (MP4) SetArtist(artist string) error {
	panic("implement me")
}

func (MP4) SetAlbum(album string) error {
	panic("implement me")
}

func (MP4) SetYear(year int) error {
	panic("implement me")
}

func (MP4) SetComment(comment string) error {
	panic("implement me")
}

func (MP4) SetGenre(genre string) error {
	panic("implement me")
}

func (MP4) SetAlbumArtist(albumArtist string) error {
	panic("implement me")
}

func (MP4) SetDate(date time.Time) error {
	panic("implement me")
}

func (MP4) SetArranger(arranger string) error {
	panic("implement me")
}

func (MP4) SetAuthor(author string) error {
	panic("implement me")
}

func (MP4) SetBMP(bmp int) error {
	panic("implement me")
}

func (MP4) SetCatalogNumber(catalogNumber string) error {
	panic("implement me")
}

func (MP4) SetCompilation(compilation string) error {
	panic("implement me")
}

func (MP4) SetComposer(composer string) error {
	panic("implement me")
}

func (MP4) SetConductor(conductor string) error {
	panic("implement me")
}

func (MP4) SetCopyright(copyright string) error {
	panic("implement me")
}

func (MP4) SetDescription(description string) error {
	panic("implement me")
}

func (MP4) SetDiscNumber(number int, total int) error {
	panic("implement me")
}

func (MP4) SetEncodedBy(encodedBy string) error {
	panic("implement me")
}

func (MP4) SetTrackNumber(number int, total int) error {
	panic("implement me")
}

func (MP4) SetPicture(picture image.Image) error {
	panic("implement me")
}

func (MP4) DeleteAll() error {
	panic("implement me")
}

func (MP4) DeleteTitle() error {
	panic("implement me")
}

func (MP4) DeleteArtist() error {
	panic("implement me")
}

func (MP4) DeleteAlbum() error {
	panic("implement me")
}

func (MP4) DeleteYear() error {
	panic("implement me")
}

func (MP4) DeleteComment() error {
	panic("implement me")
}

func (MP4) DeleteGenre() error {
	panic("implement me")
}

func (MP4) DeleteAlbumArtist() error {
	panic("implement me")
}

func (MP4) DeleteDate() error {
	panic("implement me")
}

func (MP4) DeleteArranger() error {
	panic("implement me")
}

func (MP4) DeleteAuthor() error {
	panic("implement me")
}

func (MP4) DeleteBMP() error {
	panic("implement me")
}

func (MP4) DeleteCatalogNumber() error {
	panic("implement me")
}

func (MP4) DeleteCompilation() error {
	panic("implement me")
}

func (MP4) DeleteComposer() error {
	panic("implement me")
}

func (MP4) DeleteConductor() error {
	panic("implement me")
}

func (MP4) DeleteCopyright() error {
	panic("implement me")
}

func (MP4) DeleteDescription() error {
	panic("implement me")
}

func (MP4) DeleteDiscNumber() error {
	panic("implement me")
}

func (MP4) DeleteEncodedBy() error {
	panic("implement me")
}

func (MP4) DeleteTrackNumber() error {
	panic("implement me")
}

func (MP4) DeletePicture() error {
	panic("implement me")
}

func (MP4) SaveFile(path string) error {
	panic("implement me")
}

func (MP4) Save(input io.WriteSeeker) error {
	panic("implement me")
}

func (mp4 *MP4) getString(tag string) (string, error) {
	val, ok := mp4.data[tag]
	if !ok {
		return "", ErrorTagNotFound
	}
	return val.(string), nil
}

func (mp4 *MP4) getInt(tag string) (int, error) {
	val, ok := mp4.data[tag]
	if !ok {
		return 0, ErrorTagNotFound
	}
	return val.(int), nil
}

func checkMp4(input io.ReadSeeker) bool {
	if input == nil {
		return false
	}

	data, err := seekAndRead(input, 0, io.SeekStart, 12)
	if err != nil {
		return false
	}
	marker := string(data[4:8])

	if marker == MP4_MARKER {
		mp4type := string(data[8:12])
		for _, t := range MP4_TYPES {
			if mp4type == t {
				return true
			}
		}
	}

	return false
}

func ReadMp4(input io.ReadSeeker) (*MP4, error) {
	header := MP4{}
	header.data = map[string]interface{}{}

	// Seek to file start
	startIndex, err := input.Seek(0, io.SeekStart)
	if startIndex != 0 {
		return nil, ErrorSeekFile
	}

	if err != nil {
		return nil, err
	}

	for {
		var size uint32 = 0
		err = binary.Read(input, binary.BigEndian, &size)
		if err != nil {
			break
		}

		nameBytes := make([]byte, 4)
		_, err = input.Read(nameBytes)
		if err != nil {
			break
		}
		name := string(nameBytes)

		bytes := make([]byte, size-8)
		_, err = input.Read(bytes)
		if err != nil {
			break
		}

		if name == MP4_MOOV_ATOM {
			parseMoovAtom(bytes, &header)
		}
	}

	return &header, nil
}

func parseMoovAtom(bytes []byte, mp4 *MP4) {
	for {
		size := binary.BigEndian.Uint32(bytes[0:4])
		name := string(bytes[4:8])

		if name == MP4_META_ATOM {
			bytes = bytes[4:]
			size = size - 4
			parseMoovAtom(bytes[8:], mp4)

		} else if name == MP4_META_UPTA || name == MP4_META_ILST {
			parseMoovAtom(bytes[8:], mp4)
		} else {
			atomName, ok := atoms[name]
			if ok {
				parseAtomData(bytes[8:size], atomName, mp4)
			}
		}

		bytes = bytes[size:]

		if len(bytes) == 0 {
			break
		}
	}
}

func parseAtomData(bytes []byte, atomName string, mp4 *MP4) {
	// TODO : different types
	value := string(bytes[16:])
	mp4.data[atomName] = value
	//println(atomName)

	datatype := binary.BigEndian.Uint32(bytes[8:12])
	if datatype == 13 {
		mp4.data[atomName] = AttachedPicture{
			MIME: "image/jpeg",
			Data: bytes[16:],
		}

	}

	if atomName == MP4_TAG_TRACK || atomName == MP4_TAG_DISC {
		mp4.data[atomName] = int(bytes[19:20][0])
		mp4.data[atomName+"_TOTAL"] = int(bytes[21:22][0])
	}
}
