package dengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dataToTest = `[Desktop Entry]
Name=Blender
Exec=blender
Icon=blender.svg
`
)

func TestItWorks(t *testing.T) {
	assert.True(t, true)
}

func TestBuildFileData(t *testing.T) {
	// Given Im a component object from application
	// I want to build a file with the data I pass to DesktopDataObject
	// Then a file data as string must be built with all buffer content

	ddata := NewDesktopData()
	ddata.AddField(DesktopEntry, "")
	ddata.AddField(NameEntry, "Blender")
	ddata.AddField(ExecEntry, "blender")
	ddata.AddField(IconEntry, "blender.svg")
	filedata := ddata.BuildFileData()
	assert.Equal(t, dataToTest, filedata)
}

func TestNewDesktopData(t *testing.T) {
	// Given Im a component object from application
	// I want to create a DesktopData object
	// Then it and his buffer cant be nil

	ddata := NewDesktopData()
	assert.NotNil(t, ddata)
	assert.NotNil(t, ddata.buffer)
}

func TestFillDesktopDataBuffer(t *testing.T) {
	// Given Im a component object from application]
	// I want to create a DesktopData object and fill desktop fields in its buffer
	// Then the buffer cant be nil or empy and the buffer needs to have the exact
	// fields thats was appended to its buffer

	ddata := NewDesktopData()
	ddata.AddField(DesktopEntry, "")
	ddata.AddField(NameEntry, "Blender")
	assert.NotNil(t, ddata.buffer)
	assert.Equal(t, 2, len(ddata.buffer))
	assert.Equal(t, "", ddata.buffer[0].value)
	assert.Equal(t, "Blender", ddata.buffer[1].value)
}
