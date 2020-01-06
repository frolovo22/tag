package tag

import (
	"errors"
	"image"
	"io"
	"time"
)

type ID3v22Frame struct {
	Key   string
	Value []byte
}

type ID3v22 struct {
	Marker string // Always
	Length int
	Frames []ID3v22Frame
}

func (id3v2 *ID3v22) GetAllTagNames() []string {
	panic("implement me")
}

func (id3v2 *ID3v22) GetVersion() TagVersion {
	panic("implement me")
}

func (id3v2 *ID3v22) GetFileData() []byte {
	panic("implement me")
}

func (id3v2 *ID3v22) GetTitle() (string, error) {
	return id3v2.GetString("TT2")
}

func (id3v2 *ID3v22) GetArtist() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetAlbum() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetYear() (int, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetComment() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetGenre() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetAlbumArtist() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetDate() (time.Time, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetArranger() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetAuthor() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetBMP() (int, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetCatalogNumber() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetCompilation() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetComposer() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetConductor() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetCopyright() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetDescription() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetDiscNumber() (int, int, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetEncodedBy() (string, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetTrackNumber() (int, int, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) GetPicture() (image.Image, error) {
	panic("implement me")
}

func (id3v2 *ID3v22) SetTitle(title string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetArtist(artist string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetAlbum(album string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetYear(year int) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetComment(comment string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetGenre(genre string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetAlbumArtist(albumArtist string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetDate(date time.Time) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetArranger(arranger string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetAuthor(author string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetBMP(bmp int) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetCatalogNumber(catalogNumber string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetCompilation(compilation string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetComposer(composer string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetConductor(conductor string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetCopyright(copyright string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetDescription(description string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetDiscNumber(number int, total int) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetEncodedBy(encodedBy string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetTrackNumber(number int, total int) error {
	panic("implement me")
}

func (id3v2 *ID3v22) SetPicture(picture image.Image) error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteAll() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteTitle() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteArtist() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteAlbum() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteYear() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteComment() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteGenre() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteAlbumArtist() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteDate() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteArranger() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteAuthor() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteBMP() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteCatalogNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteCompilation() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteComposer() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteConductor() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteCopyright() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteDescription() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteDiscNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteEncodedBy() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeleteTrackNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v22) DeletePicture() error {
	panic("implement me")
}

func (id3v2 *ID3v22) SaveFile(path string) error {
	panic("implement me")
}

func (id3v2 *ID3v22) Save(input io.WriteSeeker) error {
	panic("implement me")
}

func ReadID3v22(input io.ReadSeeker) (*ID3v22, error) {
	header := ID3v22{}

	if input == nil {
		return nil, ErrorEmptyFile
	}

	// Seek to file start
	startIndex, err := input.Seek(0, io.SeekStart)
	if startIndex != 0 {
		return nil, ErrorSeekFile
	}

	if err != nil {
		return nil, err
	}

	// Header size
	headerByte := make([]byte, 10)
	nReaded, err := input.Read(headerByte)
	if err != nil {
		return nil, err
	}
	if nReaded != 10 {
		return nil, errors.New("error header length")
	}

	// Marker
	marker := string(headerByte[0:3])
	if marker != "ID3" {
		return nil, errors.New("error file marker")
	}

	header.Marker = marker

	// Version
	versionByte := headerByte[3]
	if versionByte != 2 {
		return nil, ErrorUnsupportedFormat
	}

	// Length
	length := ByteToIntSynchsafe(headerByte[6:10])
	header.Length = length

	curRead := 0
	for curRead < length {
		bytesExtendedHeader := make([]byte, 6)
		nReaded, err = input.Read(bytesExtendedHeader)
		if err != nil {
			return nil, err
		}
		if nReaded != 6 {
			return nil, errors.New("error extended header length")
		}
		// Frame identifier
		key := string(bytesExtendedHeader[0:3])

		// Frame data size
		size := ByteToInt(bytesExtendedHeader[3:6])

		bytesExtendedValue := make([]byte, size)
		nReaded, err = input.Read(bytesExtendedValue)
		if err != nil {
			return nil, err
		}
		if nReaded != size {
			return nil, errors.New("error extended value length")
		}

		if key[0:1] == "T" {
			pos := -1
			for i, v := range bytesExtendedValue {
				if v == 0 && i > 0 {
					pos = i
				}
			}
			if pos != -1 {
				bytesExtendedValue = bytesExtendedValue[0:pos]
			}

		}

		header.Frames = append(header.Frames, ID3v22Frame{
			key,
			bytesExtendedValue,
		})

		curRead += 6 + size

	}
	return &header, nil
}

func checkID3v22(input io.ReadSeeker) bool {
	if input == nil {
		return false
	}

	// read marker (3 bytes) and version (1 byte) for ID3v2
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return false
	}
	marker := string(data[0:3])

	// id3v2
	if marker != "ID3" {
		return false
	}

	versionByte := data[3]

	if versionByte != 2 {
		return false
	}

	return true
}

func (id3v2 *ID3v22) GetString(name string) (string, error) {
	for _, val := range id3v2.Frames {
		if val.Key == name {
			return GetString(val.Value)
		}
	}
	return "", ErrorTagNotFound
}
