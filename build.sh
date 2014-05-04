#!/usr/bin/env bash
go build
./ggir glib
./ggir gobject
./ggir gdk
./ggir gdkpixbuf
./ggir cairo
./ggir gtk
