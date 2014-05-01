#!/usr/bin/env bash
go build
./ggir gobject gobject/gobject.xml
./ggir glib glib/glib.xml
./ggir gtk gtk/gtk.xml
