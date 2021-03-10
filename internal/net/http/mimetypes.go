package mimetypes

// File signature
type MagicBytes struct {
	offset int
	bytes []byte
}

type MimeType struct {
	MagicBytes []MagicBytes
	FileExt []string
	Value string
}

var (
	MimeTypes = []MimeType{
		{
			MagicBytes: []MagicBytes{
				{
					offset: 0,
					bytes:  []byte{ 0x47, 0x49, 0x46, 0x38, 0x37, 0x61 },
				},
				{
					offset: 0,
					bytes:  []byte{ 0x47, 0x49, 0x46, 0x38, 0x39, 0x61 },
				},
			},
			FileExt:    []string{"gif"},
			Value:      "image/gif",
		},
	}
)
