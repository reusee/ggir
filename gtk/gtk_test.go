package gtk

import (
	"testing"
	"time"
)

func TestSimpleWindow(t *testing.T) {
	Init(nil, nil)
	win := WindowNew(WINDOW_TOPLEVEL)
	win.SetTitle("foobarbaz")
	label := LabelNew("Hello, world!")
	win.Add(label)
	win.ShowAll()
	go func() {
		time.Sleep(time.Second * 3)
		MainQuit()
	}()
	Main()
}
