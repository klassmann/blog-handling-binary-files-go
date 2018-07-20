package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

type Header struct {
	Magic  [4]byte // Magic literal, the constant "icns"
	Length uint32  // Length of data, in bytes
}

type IconData struct {
	Type   [4]byte // OSType
	Length uint32  // Length of Data
	Data   []byte  // Icon Data
}

// TypeStr transforms Type as []byte to string
func (i *IconData) TypeStr() string {
	return fmt.Sprintf("%s", i.Type)
}

// AppleIcon is the format of icns file
type AppleIcon struct {
	Header
	Icons []IconData // All Icons
}

// Print show the information inside the file
func (i *AppleIcon) Print() {
	fmt.Printf("Header Magic: %s\n", i.Header.Magic)
	fmt.Printf("Header Length: %d\n", i.Header.Length)
	fmt.Println("[Icons]")
	for i, icon := range i.Icons {
		fmt.Printf("%d - %s - Len: %d\n", i, icon.Type, icon.Length)
		if icon.TypeStr() == "ic09" {
			ioutil.WriteFile("icone.jpg", icon.Data, 0666)
		}
	}
}

// ReadAppleIcon uses the reader to read bytes into de AppleIcon structure
func ReadAppleIcon(r *bytes.Reader) (*AppleIcon, error) {
	var icns AppleIcon

	binary.Read(r, binary.BigEndian, &icns.Header)

	// We have to iterate until end of the file
	for {
		var icon IconData
		err := binary.Read(r, binary.LittleEndian, &icon.Type)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("error reading icons: %s", err)
		}

		binary.Read(r, binary.BigEndian, &icon.Length)
		fmt.Printf("%d\n", icon.Length)
		data := make([]byte, icon.Length-8)
		binary.Read(r, binary.BigEndian, data)
		w := bytes.Buffer{}
		w.Write(data)
		icon.Data = w.Bytes()
		icns.Icons = append(icns.Icons, icon)
	}

	return &icns, nil
}

func main() {

	data, err := ioutil.ReadFile("OpenEmu.icns")

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)
	icon, err := ReadAppleIcon(reader)
	if err != nil {
		panic(err)
	}
	icon.Print()
}
