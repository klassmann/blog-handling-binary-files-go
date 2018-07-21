package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReadAppleIcon(t *testing.T) {

	Magic := [4]byte{'i', 'c', 'n', 's'}

	f, err := ioutil.ReadFile("OpenEmu.icns")

	if err != nil {
		t.Error(err)
	}
	reader := bytes.NewReader(f)

	icon, err := ReadAppleIcon(reader)

	if err != nil {
		t.Error(err)
	}

	if icon.Header.Magic != Magic {
		t.Errorf("Magic Signature must be 'icns' but %s", icon.Header.Magic)
	}

	if icon.Header.Length <= 0 {
		t.Errorf("Length must be more than 0, but %d.", icon.Header.Length)
	}

	if len(icon.Icons) != 13 {
		t.Errorf("The file OpenEmu.icns must have 13 icons, but %d.", len(icon.Icons))
	}
}
