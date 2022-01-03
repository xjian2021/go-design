package adapter

import "testing"

func TestOSAdapter(t *testing.T) {
	m := NewMac()
	t.Log(m.FileType())
	win := NewWindowsOS(m)
	t.Log(win.File())
}