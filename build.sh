#!/usr/bin/env bash
go build
./ggir glib
./ggir gobject
./ggir gio
./ggir gdk
./ggir gdkpixbuf
./ggir cairo
./ggir gtk
echo 'go get'
go get -u github.com/reusee/ggir/glib
go get -u github.com/reusee/ggir/gobject
go get -u github.com/reusee/ggir/gio
go get -u github.com/reusee/ggir/gdk
go get -u github.com/reusee/ggir/gdkpixbuf
go get -u github.com/reusee/ggir/cairo
go get -u github.com/reusee/ggir/gtk
