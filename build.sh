#!/usr/bin/env bash
go build
./ggir glib
./ggir gobject
./ggir gio
./ggir gdk
./ggir gdkpixbuf
./ggir cairo
./ggir gtk
