package types

const (
	TextFile    = '0'
	SubMenu     = '1'
	CCSO        = '2'
	Error       = '3'
	BinHex      = '4'
	DOSBinary   = '5'
	UnixEncoded = '6'
	Search      = '7'
	Telnet      = '8'
	BinaryFile  = '9'
	Server      = '+'
	GIFFile     = 'g'
	Image       = 'I'
	Telnet3270  = 'T'
	BitmapImage = ':'
	MovieFile   = ';'
	Sound       = '<'
	Doc         = 'd'
	HTMLFile    = 'h'
	Information = 'i'
	ImageAlt    = 'p'
	RTFFile     = 'r'
	SoundAlt    = 's'
	PDFFile     = 'P'
	XMLFile     = 'X'
	End         = '.'
)

/*
type TextFile struct{}
type SubMenu struct{}
type CCSONameserver struct{}
type Error struct{}
type BinHexEncodedFile struct{}
type DOSFile struct{}
type UUEncodedFile struct{}
type Search struct{}
type Telnet struct{}
type BinaryFile struct{}
type Mirror struct{}
type GIFFile struct{}
type ImageCanonical struct{}
type Telnet3270 struct{}
type BitmapImage struct{}
type MovieFile struct{}
type SoundFile struct{}
type Doc struct{}
type HTMLFile struct{}
type Information struct{}
type ImageAlt struct{}
type RTFFile struct{}
type SoundFileAlt struct{}
type PDFFile struct{}
type XMLFile struct{}

var typeMap = map[byte]any{
	'0': TextFile{},
	'1': SubMenu{},
	'2': CCSONameserver{},
	'3': Error{},
	'4': BinHexEncodedFile{},
	'5': DOSFile{},
	'6': UUEncodedFile{},
	'7': Search{},
	'8': Telnet{},
	'9': BinaryFile{},
	'+': Mirror{},
	'g': GIFFile{},
	'I': ImageCanonical{},
	'T': Telnet3270{},
	':': BitmapImage{},
	';': MovieFile{},
	'<': SoundFile{},
	'd': Doc{},
	'h': HTMLFile{},
	'i': Information{},
	'p': ImageAlt{},
	'r': RTFFile{},
	's': SoundFileAlt{},
	'P': PDFFile{},
	'X': XMLFile{},
}

func GetType(c byte) any {
	res := typeMap[c]
	if res != nil {
		return res
	}

	return typeMap['3']
}
*/
