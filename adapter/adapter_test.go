package adapter

import "testing"

func TestOSAdapter(t *testing.T) {
	m := NewMac("cmd")
	t.Logf("mac file :%s",m.MacFile())
	win := NewWindowsOS(m)
	t.Logf("windows file :%s",win.WinFile())
}