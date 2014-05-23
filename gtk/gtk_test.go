package gtk

import (
	"fmt"
	"testing"
)

func TestSimpleWindow(t *testing.T) {
	Init(nil, nil)
	win := WindowNew(WINDOW_TOPLEVEL)
	win.SetTitle("foobarbaz")
	button := ButtonNewWithLabel("Hello, world!")
	button.OnClicked(func() {
		fmt.Printf("clicked\n")
	})
	win.Add(button)
	win.ShowAll()
	win.Connect("destroy", func() {
		fmt.Printf("Quit\n")
		MainQuit()
	})
	Main()
}
