package dengine

import "bytes"

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
	de := &DesktopData{
		buffer: make([]fileData, 0),
	}

	de.AddField(DesktopEntry, "")
	return de
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
