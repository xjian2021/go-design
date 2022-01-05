package adapter

type MacOSAdapter interface {
	WinFile() string
}

type MacOS interface {
	MacFile() string
}

type mac struct {
	name string
}

func NewMac(name string) MacOS {
	return &mac{name}
}

func (m *mac) MacFile() string {
	return m.name
}

type windowsOS struct {
	MacOS
}

func NewWindowsOS(macOS MacOS) MacOSAdapter {
	return &windowsOS{MacOS: macOS}
}

func (O *windowsOS) WinFile() string {
	return O.MacFile() + ".exe"
}

