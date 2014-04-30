package gtk

import "testing"

func TestSimpleWindow(t *testing.T) {
	Init(nil, nil)
	win := WindowNew(WINDOW_TOPLEVEL)
	win.SetTitle("foobarbaz")
	label := LabelNew("Hello, world!")
	win.Add(label)
	win.ShowAll()
	Main()
}
