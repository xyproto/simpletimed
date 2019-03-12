# The Simple Timed Wallpaper Format

[![GoDoc](https://godoc.org/github.com/xyproto/simpletimed?status.svg)](https://godoc.org/github.com/xyproto/simpletimed)

## The specification

### 1.0.0

* [Markdown](https://github.com/xyproto/simpletimed/blob/master/stw-1.0.0.md)
* [PDF](https://github.com/xyproto/simpletimed/blob/master/stw-1.0.0.pdf)

## Go module

The `simpletimed` Go module can be used for parsing the file format and for running an event loop for setting the wallpaper, given a function with this signature:

    func(string) error

Where the given string is the image filename to be set.

## General Info

* Version: 1.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
