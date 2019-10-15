package tag

import (
	"bytes"
	"encoding/binary"
	"image"
	"io"
	"strings"
	"time"
)

type FLAC struct {
	Blocks []*FlacMeatadataBlock
}

func (F *FLAC) GetAllTagNames() []string {
	panic("implement me")
}

func (flac *FLAC) GetVersion() TagVersion {
	return TagVersionFLAC
}

func (F *FLAC) GetFileData() []byte {
	panic("implement me")
}

func (flac *FLAC) GetTitle() (string, error) {
	return flac.GetVorbisComment("TITLE")
}

func (flac *FLAC) GetArtist() (string, error) {
	return flac.GetVorbisComment("ARTIST")
}

func (flac *FLAC) GetAlbum() (string, error) {
	return flac.GetVorbisComment("ALBUM")
}

func (F *FLAC) GetYear() (int, error) {
	panic("implement me")
}

func (F *FLAC) GetComment() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetGenre() (string, error) {
	panic("implement me")
}

func (flac *FLAC) GetAlbumArtist() (string, error) {
	return flac.GetVorbisComment("ALBUMARTIST")
}

func (F *FLAC) GetDate() (time.Time, error) {
	panic("implement me")
}

func (F *FLAC) GetArranger() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetAuthor() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetBMP() (int, error) {
	panic("implement me")
}

func (F *FLAC) GetCatalogNumber() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetCompilation() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetComposer() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetConductor() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetCopyright() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetDescription() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetDiscNumber() (int, int, error) {
	panic("implement me")
}

func (F *FLAC) GetEncodedBy() (string, error) {
	panic("implement me")
}

func (F *FLAC) GetTrackNumber() (int, int, error) {
	panic("implement me")
}

func (F *FLAC) GetPicture() (image.Image, error) {
	panic("implement me")
}

func (F *FLAC) SetTitle(title string) error {
	panic("implement me")
}

func (F *FLAC) SetArtist(artist string) error {
	panic("implement me")
}

func (F *FLAC) SetAlbum(album string) error {
	panic("implement me")
}

func (F *FLAC) SetYear(year int) error {
	panic("implement me")
}

func (F *FLAC) SetComment(comment string) error {
	panic("implement me")
}

func (F *FLAC) SetGenre(genre string) error {
	panic("implement me")
}

func (F *FLAC) SetAlbumArtist(albumArtist string) error {
	panic("implement me")
}

func (F *FLAC) SetDate(date time.Time) error {
	panic("implement me")
}

func (F *FLAC) SetArranger(arranger string) error {
	panic("implement me")
}

func (F *FLAC) SetAuthor(author string) error {
	panic("implement me")
}

func (F *FLAC) SetBMP(bmp int) error {
	panic("implement me")
}

func (F *FLAC) SetCatalogNumber(catalogNumber string) error {
	panic("implement me")
}

func (F *FLAC) SetCompilation(compilation string) error {
	panic("implement me")
}

func (F *FLAC) SetComposer(composer string) error {
	panic("implement me")
}

func (F *FLAC) SetConductor(conductor string) error {
	panic("implement me")
}

func (F *FLAC) SetCopyright(copyright string) error {
	panic("implement me")
}

func (F *FLAC) SetDescription(description string) error {
	panic("implement me")
}

func (F *FLAC) SetDiscNumber(number int, total int) error {
	panic("implement me")
}

func (F *FLAC) SetEncodedBy(encodedBy string) error {
	panic("implement me")
}

func (F *FLAC) SetTrackNumber(number int, total int) error {
	panic("implement me")
}

func (F *FLAC) SetPicture(picture image.Image) error {
	panic("implement me")
}

func (F *FLAC) DeleteAll() error {
	panic("implement me")
}

func (F *FLAC) DeleteTitle() error {
	panic("implement me")
}

func (F *FLAC) DeleteArtist() error {
	panic("implement me")
}

func (F *FLAC) DeleteAlbum() error {
	panic("implement me")
}

func (F *FLAC) DeleteYear() error {
	panic("implement me")
}

func (F *FLAC) DeleteComment() error {
	panic("implement me")
}

func (F *FLAC) DeleteGenre() error {
	panic("implement me")
}

func (F *FLAC) DeleteAlbumArtist() error {
	panic("implement me")
}

func (F *FLAC) DeleteDate() error {
	panic("implement me")
}

func (F *FLAC) DeleteArranger() error {
	panic("implement me")
}

func (F *FLAC) DeleteAuthor() error {
	panic("implement me")
}

func (F *FLAC) DeleteBMP() error {
	panic("implement me")
}

func (F *FLAC) DeleteCatalogNumber() error {
	panic("implement me")
}

func (F *FLAC) DeleteCompilation() error {
	panic("implement me")
}

func (F *FLAC) DeleteComposer() error {
	panic("implement me")
}

func (F *FLAC) DeleteConductor() error {
	panic("implement me")
}

func (F *FLAC) DeleteCopyright() error {
	panic("implement me")
}

func (F *FLAC) DeleteDescription() error {
	panic("implement me")
}

func (F *FLAC) DeleteDiscNumber() error {
	panic("implement me")
}

func (F *FLAC) DeleteEncodedBy() error {
	panic("implement me")
}

func (F *FLAC) DeleteTrackNumber() error {
	panic("implement me")
}

func (F *FLAC) DeletePicture() error {
	panic("implement me")
}

func (F *FLAC) SaveFile(path string) error {
	panic("implement me")
}

func (F *FLAC) Save(input io.WriteSeeker) error {
	panic("implement me")
}

func checkFLAC(input io.ReadSeeker) bool {
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return false
	}

	if string(data) != "fLaC" {
		return false
	}

	return true
}

type FlacMetadataBlockType byte

/*
BLOCK_TYPE:
	0 : STREAMINFO
	1 : PADDING
	2 : APPLICATION
	3 : SEEKTABLE
	4 : VORBIS_COMMENT
	5 : CUESHEET
	6 : PICTURE
	7-126 : reserved
	127 : invalid, to avoid confusion with a frame sync code
*/
const (
	FlacStreamInfo    FlacMetadataBlockType = 0
	FlacPadding       FlacMetadataBlockType = 1
	FlacApplication   FlacMetadataBlockType = 2
	FlacSeekTable     FlacMetadataBlockType = 3
	FlacVorbisComment FlacMetadataBlockType = 4
	FlacCueSheet      FlacMetadataBlockType = 5
	FlacPicture       FlacMetadataBlockType = 6
)

type FlacMeatadataBlock struct {
	IsLast bool // Last-metadata-block flag: '1' if this block is the last metadata block before the audio blocks, '0' otherwise.
	Type   FlacMetadataBlockType
	Size   int
	Data   []byte
}

func ReadFLAC(input io.ReadSeeker) (*FLAC, error) {
	flac := FLAC{}

	// FLAC identifier
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return nil, err
	}

	if string(data) != "fLaC" {
		return nil, ErrorFileMarker
	}

	// read blocks
	for {
		block, err := readMeatadataBlock(input)
		if err != nil {
			return nil, err
		}

		// last block before audio frame
		if block.IsLast {
			break
		}

		flac.Blocks = append(flac.Blocks, block)
	}

	return &flac, nil
}

func readMeatadataBlock(input io.Reader) (*FlacMeatadataBlock, error) {
	header := FlacMeatadataBlock{}

	// 4 - header size
	headerBytes, err := readBytes(input, 4)
	if err != nil {
		return nil, err
	}

	// first bit
	isLastBit := GetBit(headerBytes[0], 7)
	if isLastBit == 1 {
		header.IsLast = true
	}

	// Only 0-6 types
	// 1-7 bits
	blockType := headerBytes[0] & 0x7F
	if blockType > 6 {
		return nil, ErrorReadFile
	}
	header.Type = FlacMetadataBlockType(blockType)

	// 3 bytes size
	header.Size = ByteToInt(headerBytes[1:])

	// block data
	header.Data, err = readBytes(input, header.Size)
	if err != nil {
		return nil, err
	}
	return &header, nil
}

type VorbisComment struct {
	Name  string
	Value string
}

func (flac *FLAC) GetVorbisComment(key string) (string, error) {
	for _, block := range flac.Blocks {
		if block.Type == FlacVorbisComment {
			comments, err := readVorbisComments(bytes.NewReader(block.Data))
			if err != nil {
				return "", err
			}
			for _, comment := range comments {
				if comment.Name == key {
					return comment.Value, nil
				}
			}
		}
	}
	return "", ErrorTagNotFound
}

//The comment header is decoded as follows:
//
//	1) [vendor_length] = read an unsigned integer of 32 bits
//	2) [vendor_string] = read a UTF-8 vector as [vendor_length] octets
//	3) [user_comment_list_length] = read an unsigned integer of 32 bits
//	4) iterate [user_comment_list_length] times {
//
//		5) [length] = read an unsigned integer of 32 bits
//		6) this iteration's user comment = read a UTF-8 vector as [length] octets
//
//	}
//
//	7) [framing_bit] = read a single bit as boolean
func readVorbisComments(input io.Reader) ([]VorbisComment, error) {
	result := []VorbisComment{}

	// vendor
	_, err := readLengthData(input)
	if err != nil {
		return nil, err
	}

	// user_comment_list_length
	var length uint32
	err = binary.Read(input, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}

	// iterate
	for i := 0; i < int(length); i++ {
		data, err := readLengthData(input)
		if err != nil {
			return nil, err
		}

		// Parse data
		vorbis := strings.SplitN(string(data), "=", 2)
		if len(vorbis) != 2 {
			return nil, ErrorIncorrectTag
		}

		comment := VorbisComment{
			Name:  vorbis[0],
			Value: vorbis[1],
		}
		result = append(result, comment)
	}
	return result, nil
}
