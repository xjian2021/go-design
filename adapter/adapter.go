package adapter

type OSAdapter interface {
	File() string
}

type MacOS interface {
	FileType() string
}

type mac struct {}

func NewMac() *mac {
	return &mac{}
}

func (m *mac) FileType() string {
	return "main"
}

type windowsOS struct {
	MacOS
}

func NewWindowsOS(macOS MacOS) *windowsOS {
	return &windowsOS{MacOS: macOS}
}

func (O *windowsOS) File() string {
	return O.FileType() + " -> " + "main.exe"
}

