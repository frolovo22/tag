package tag

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"time"
)

type Id3v2Flags byte

func (flags Id3v2Flags) String() string {
	return strconv.Itoa(int(flags))
}

func (flags Id3v2Flags) IsUnsynchronisation() bool {
	return GetBit(byte(flags), 7) == 1
}

func (flags Id3v2Flags) SetUnsynchronisation(data bool) {
	SetBit((*byte)(&flags), data, 7)
}

func (flags Id3v2Flags) HasExtendedHeader() bool {
	return GetBit(byte(flags), 6) == 1
}

func (flags Id3v2Flags) SetExtendedHeader(data bool) {
	SetBit((*byte)(&flags), data, 7)
}

func (flags Id3v2Flags) IsExperimentalIndicator() bool {
	return GetBit(byte(flags), 5) == 1
}

func (flags Id3v2Flags) SetExperimentalIndicator(data bool) {
	SetBit((*byte)(&flags), data, 7)
}

type ID3v2Frame struct {
	Key   string
	Value []byte
}

type ID3v2 struct {
	Marker     string // Always 'ID3'
	Version    TagVersion
	SubVersion int
	Flags      Id3v2Flags
	Length     int
	Frames     []ID3v2Frame

	Data []byte
}

func (id3v2 *ID3v2) GetAllTagNames() []string {
	var result []string
	for _, value := range id3v2.Frames {
		result = append(result, value.Key)
	}
	return result
}

func (id3v2 *ID3v2) GetVersion() TagVersion {
	return id3v2.Version
}

func (id3v2 *ID3v2) GetFileData() []byte {
	return id3v2.Data
}

func (id3v2 *ID3v2) GetTitle() (string, error) {
	return id3v2.getString("TIT2")
}

func (id3v2 *ID3v2) GetArtist() (string, error) {
	return id3v2.getString("TPE1")
}

func (id3v2 *ID3v2) GetAlbum() (string, error) {
	return id3v2.getString("TALB")
}

func (id3v2 *ID3v2) GetYear() (int, error) {
	date, err := id3v2.getTimestamp("TDOR")
	return date.Year(), err
}

func (id3v2 *ID3v2) GetComment() (string, error) {
	// id3v2
	// Comment struct must be greater than 4
	// [lang \x00 text] - comment format
	// lang - 3 symbols
	// \x00 - const, delimeter
	// text - all after
	commentStr, err := id3v2.getString("COMM")
	if err != nil {
		return "", err
	}

	if len(commentStr) < 4 {
		return "", ErrorIncorrectLength
	}

	return commentStr[4:], nil
}

func (id3v2 *ID3v2) GetGenre() (string, error) {
	return id3v2.getString("TCON")
}

func (id3v2 *ID3v2) GetAlbumArtist() (string, error) {
	return id3v2.getString("TPE2")
}

func (id3v2 *ID3v2) GetDate() (time.Time, error) {
	return id3v2.getTimestamp("TDRC")
}

func (id3v2 *ID3v2) GetArranger() (string, error) {
	return id3v2.getString("TIPL")
}

func (id3v2 *ID3v2) GetAuthor() (string, error) {
	return id3v2.getString("TOLY")
}

func (id3v2 *ID3v2) GetBMP() (int, error) {
	return id3v2.getInt("TBPM")
}

func (id3v2 *ID3v2) GetCatalogNumber() (int, error) {
	return id3v2.getInt("TXXX")
}

func (id3v2 *ID3v2) GetCompilation() (string, error) {
	return id3v2.getString("TCMP")
}

func (id3v2 *ID3v2) GetComposer() (string, error) {
	return id3v2.getString("TCOM")
}

func (id3v2 *ID3v2) GetConductor() (string, error) {
	return id3v2.getString("TPE3")
}

func (id3v2 *ID3v2) GetCopyright() (string, error) {
	return id3v2.getString("TCOP")
}

func (id3v2 *ID3v2) GetDescription() (string, error) {
	return id3v2.getString("TIT3")
}

func (id3v2 *ID3v2) GetDiscNumber() (int, int, error) {
	dickNumber, err := id3v2.getString("TPOS")
	if err != nil {
		return 0, 0, err
	}
	numbers := strings.Split(dickNumber, "/")
	if len(numbers) != 2 {
		return 0, 0, ErrorIncorrectLength
	}
	number, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, 0, err
	}
	total, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, 0, err
	}
	return number, total, nil
}

func (id3v2 *ID3v2) GetEncodedBy() (string, error) {
	return id3v2.getString("TENC")
}

func (id3v2 *ID3v2) GetTrackNumber() (int, int, error) {
	track, err := id3v2.getInt("TRCK")
	return track, track, err
}

func (id3v2 *ID3v2) GetPicture() (Picture, error) {
	panic("implement me")
}

func (id3v2 *ID3v2) SetTitle(title string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetArtist(artist string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetAlbum(album string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetYear(year int) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetComment(comment string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetGenre(genre string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetAlbumArtist(albumArtist string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetDate(date time.Time) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetArranger(arranger string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetAuthor(author string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetBMP(bmp int) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetCatalogNumber(catalogNumber int) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetCompilation(compilation string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetComposer(composer string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetConductor(conductor string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetCopyright(copyright string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetDescription(description string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetDiscNumber(number int, total int) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetEncodedBy(encodedBy string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetTrackNumber(number int, total int) error {
	panic("implement me")
}

func (id3v2 *ID3v2) SetPicture(picture Picture) error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAll() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteTitle() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteArtist() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAlbum() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteYear() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteComment() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteGenre() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAlbumArtist() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteDate() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteArranger() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAuthor() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteBMP() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteCatalogNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteCompilation() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteComposer() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteConductor() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteCopyright() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteDescription() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteDiscNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteEncodedBy() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteTrackNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeletePicture() error {
	panic("implement me")
}

func (id3v2 *ID3v2) SaveFile(path string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) Save(input io.WriteSeeker) error {
	panic("implement me")
}

func (id3v2 *ID3v2) String() string {
	result := "Marker: " + id3v2.Marker + "\n" +
		"Version: " + id3v2.Version.String() + "\n" +
		"Subversion: " + strconv.Itoa(id3v2.SubVersion) + "\n" +
		"Flags: " + id3v2.Flags.String() + "\n" +
		"Length: " + strconv.Itoa(id3v2.Length) + "\n"

	for _, frame := range id3v2.Frames {
		result += frame.Key + ": " + string(frame.Value) + "\n"
	}

	return result
}

func checkID3v2(input io.ReadSeeker) TagVersion {
	if input == nil {
		return TagVersionUndefined
	}

	// read marker (3 bytes) and version (1 byte) for ID3v2
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return TagVersionUndefined
	}
	marker := string(data[0:3])

	// id3v2
	if marker == "ID3" {
		versionByte := data[3]
		switch versionByte {
		case 2:
			return TagVersionID3v22
		case 3:
			return TagVersionID3v23
		case 4:
			return TagVersionID3v24
		}
	}

	return TagVersionUndefined
}

func ReadID3v2(input io.ReadSeeker) (*ID3v2, error) {
	header := ID3v2{}
	if input == nil {
		return nil, errors.New("empty file")
	}

	// Seek to file start
	startIndex, err := input.Seek(0, io.SeekStart)
	if startIndex != 0 {
		return nil, errors.New("error seek file")
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
	switch versionByte {
	case 2:
		header.Version = TagVersionID3v22
	case 3:
		header.Version = TagVersionID3v23
	case 4:
		header.Version = TagVersionID3v24
	default:
		return nil, errors.New("error file version")
	}

	// Sub version
	subVersionByte := headerByte[4]
	header.SubVersion = int(subVersionByte)

	// Flags
	header.Flags = Id3v2Flags(headerByte[5])

	// Length
	length := ByteToIntSynchsafe(headerByte[6:10])
	header.Length = length

	// Extended headers
	header.Frames = []ID3v2Frame{}
	curRead := 0
	for curRead < length {
		bytesExtendedHeader := make([]byte, 10)
		nReaded, err = input.Read(bytesExtendedHeader)
		if err != nil {
			return nil, err
		}
		if nReaded != 10 {
			return nil, errors.New("error extended header length")
		}
		// Frame identifier
		key := string(bytesExtendedHeader[0:4])

		/*if bytesExtendedHeader[0] == 0 && bytesExtendedHeader[1] == 0 && bytesExtendedHeader[2] == 0 && bytesExtendedHeader[3] == 0 {
			break
		}*/

		// Frame data size
		size := ByteToInt(bytesExtendedHeader[4:8])

		bytesExtendedValue := make([]byte, size)
		nReaded, err = input.Read(bytesExtendedValue)
		if err != nil {
			return nil, err
		}
		if nReaded != size {
			return nil, errors.New("error extended value length")
		}

		header.Frames = append(header.Frames, ID3v2Frame{
			key,
			bytesExtendedValue,
		})

		curRead += 10 + size
	}

	// TODO
	if curRead != length {
		return nil, errors.New("error extended frames")
	}
	return &header, nil
}

func (id3v2 *ID3v2) getString(tagName string) (string, error) {
	for _, val := range id3v2.Frames {
		if val.Key == tagName {
			return GetString(val.Value)
		}
	}
	return "", ErrorUnsupportedTag
}

func (id3v2 *ID3v2) getTimestamp(tagName string) (time.Time, error) {
	str, err := id3v2.getString(tagName)
	if err != nil {
		return time.Now(), err
	}
	result, err := time.Parse("2006-01-02T15:04:05", str)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}

func (id3v2 *ID3v2) getInt(tagName string) (int, error) {
	for _, val := range id3v2.Frames {
		if val.Key == tagName {
			intStr, err := GetString(val.Value)
			if err != nil {
				return 0, err
			}
			return strconv.Atoi(intStr)
		}
	}
	return 0, ErrorUnsupportedTag
}