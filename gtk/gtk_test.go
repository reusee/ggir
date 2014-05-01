package gtk

import (
	"fmt"
	"testing"
)

func TestSimpleWindow(t *testing.T) {
	Init(nil, nil)
	win := WindowNew(WINDOW_TOPLEVEL)
	win.SetTitle("foobarbaz")
	label := LabelNew("Hello, world!")
	win.Add(label)
	win.ShowAll()
	win.Connect("destroy", func() {
		fmt.Printf("Quit\n")
		MainQuit()
	})
	Main()
}
