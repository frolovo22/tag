# Tag

Its pure golang library

# Install

```go 
go get github.com/frolovo22/tag
```

# Supported tags

| Name              | ID3v1       | ID3v2.2 | ID3v2.3 | ID3v2.4 |
|-------------------|-------------|---------|---------|---------|
| Title             | Title       | TT2     | TIT2    | TIT2    |
| Artist            | Artist      | TP1     | TPE1    | TPE1    |
| Album             | Album       | TOT     | TALB    | TALB    |
| Year              | Year        | TYE     | TYER    | TDOR    |
| Comment           | Comment     | COM     | COMM    | COMM    |
| Genre             | Genre       | -       | TCON    | TCON    |
| Album Artist      | -           | TOA     | TPE2    | TPE2    | 
| Date              | -           | TIM     | TYER    | TDRC    |
| Arranger          | -           | -       | IPLS    | TIPL    |
| Author            | -           | TOL     | TOLY    | TOLY    |
| BPM               | -           | BPM     | TBPM    | TBPM    |
| Catalog Number    | -           | -       | TXXX    | TXXX    |
| Compilation       | -           | -       | TCMP    | TCMP    |
| Composer          | -           | TCM     | TCOM    | TCOM    |
| Conductor         | -           | TP3     | TPE3    | TPE3    |
| Copyright         | -           | TCR     | TCOP    | TCOP    |
| Description       | -           | TXX     | TIT3    | TIT3    |
| Disc Number       | -           | -       | TPOS    | TPOS    |
| Encoded by        | -           | TEN     | TENC    | TENC    |
| Track Number      | TrackNumber | TRK     | TRCK    | TRCK    |  
| Picture           | -           | PIC     | APIC    | APIC    |
       

# Status 
In progress  
Future features:
*  Convert formats
*  Support all tags (id3 v1, v1.1, v2.2, v2.3, v2.4)
*  Fix errors in files (empty tags, incorrect size, tag size, tag parameters)
*  Command line arguments 

| Format | Read                      | Set                       | Delete                     |  Save                     |
|--------|---------------------------|---------------------------|----------------------------|---------------------------|
| idv1   | <ul><li> - [x] </li></ul> | <ul><li> - [x] </li></ul> | <ul><li> - [x] </li></ul>  | <ul><li> - [x] </li></ul> |
| idv1.1 | <ul><li> - [x] </li></ul> | <ul><li> - [x] </li></ul> | <ul><li> - [x] </li></ul>  | <ul><li> - [x] </li></ul> |
| idv2.2 | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul>  | <ul><li> - [ ] </li></ul> |
| idv2.3 | <ul><li> - [x] </li></ul> | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul>  | <ul><li> - [ ] </li></ul> |
| idv2.4 | <ul><li> - [x] </li></ul> | <ul><li> - [x] </li></ul> | <ul><li> - [ ] </li></ul>  | <ul><li> - [ ] </li></ul> |
| mp4    | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul>  | <ul><li> - [ ] </li></ul> |
| FLAC   | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul> | <ul><li> - [ ] </li></ul>  | <ul><li> - [ ] </li></ul> |

# How to use

```go
tags, err := tag.ReadFile("song.mp3")
if err != nil {
	return err
}
fmt.Println(tags.GetTitle())
```

```tag.ReadFile or tag.Read``` return interface ```Metadata```:  

```go 
type Metadata interface {
	GetMetadata
	SetMetadata
	DeleteMetadata
	SaveMetadata
}

type GetMetadata interface {
	GetAllTagNames() []string
	GetVersion() TagVersion
	GetFileData() []byte // all another file data

	GetTitle() (string, error)
	GetArtist() (string, error)
	GetAlbum() (string, error)
	GetYear() (int, error)
	GetComment() (string, error)
	GetGenre() (string, error)
	GetAlbumArtist() (string, error)
	GetDate() (time.Time, error)
	GetArranger() (string, error)
	GetAuthor() (string, error)
	GetBMP() (int, error)
	GetCatalogNumber() (int, error)
	GetCompilation() (string, error)
	GetComposer() (string, error)
	GetConductor() (string, error)
	GetCopyright() (string, error)
	GetDescription() (string, error)
	GetDiscNumber() (int, int, error) // number, total
	GetEncodedBy() (string, error)
	GetTrackNumber() (int, int, error) // number, total
	GetPicture() (image.Image, error)
}

type SetMetadata interface {
	SetTitle(title string) error
	SetArtist(artist string) error
	SetAlbum(album string) error
	SetYear(year int) error
	SetComment(comment string) error
	SetGenre(genre string) error
	SetAlbumArtist(albumArtist string) error
	SetDate(date time.Time) error
	SetArranger(arranger string) error
	SetAuthor(author string) error
	SetBMP(bmp int) error
	SetCatalogNumber(catalogNumber int) error
	SetCompilation(compilation string) error
	SetComposer(composer string) error
	SetConductor(conductor string) error
	SetCopyright(copyright string) error
	SetDescription(description string) error
	SetDiscNumber(number int, total int) error
	SetEncodedBy(encodedBy string) error
	SetTrackNumber(number int, total int) error
	SetPicture(picture image.Image) error
}

type DeleteMetadata interface {
	DeleteAll() error

	DeleteTitle() error
	DeleteArtist() error
	DeleteAlbum() error
	DeleteYear() error
	DeleteComment() error
	DeleteGenre() error
	DeleteAlbumArtist() error
	DeleteDate() error
	DeleteArranger() error
	DeleteAuthor() error
	DeleteBMP() error
	DeleteCatalogNumber() error
	DeleteCompilation() error
	DeleteComposer() error
	DeleteConductor() error
	DeleteCopyright() error
	DeleteDescription() error
	DeleteDiscNumber() error
	DeleteEncodedBy() error
	DeleteTrackNumber() error
	DeletePicture() error
}

type SaveMetadata interface {
	SaveFile(path string) error
	Save(input io.WriteSeeker) error
}
```   

