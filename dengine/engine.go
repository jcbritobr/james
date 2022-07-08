package dengine

import "bytes"

type DesktopField = uint8

const (
	DesktopEntry              = "[Desktop Entry]"
	NameEntry                 = "Name="
	GenericNameEntry          = "GenericName="
	CommentEntry              = "Comment="
	KeywordsEntry             = "Keywords="
	ExecEntry                 = "Exec="
	IconEntry                 = "Icon="
	TerminalEntry             = "Terminal="
	TypeEntry                 = "Type="
	PrefersNonDefaultGpuEntry = "PrefersNonDefaultGPU="
	CategoriesEntry           = "Categories="
	MimeTypeEntry             = "MimeTypes="
)

type fileData struct {
	key   string
	value string
}

type DesktopData struct {
	buffer []fileData
}

func NewDesktopData() *DesktopData {
	return &DesktopData{
		buffer: make([]fileData, 0),
	}
}

func (d *DesktopData) AddField(key, value string) {
	data := fileData{key, value}
	d.buffer = append(d.buffer, data)
}

func (d *DesktopData) BuildFileData() string {
	buffer := bytes.Buffer{}

	for _, v := range d.buffer {
		buffer.WriteString(v.key + v.value)
		buffer.WriteByte('\n')
	}

	return buffer.String()
}