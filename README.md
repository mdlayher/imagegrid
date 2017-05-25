imagegrid [![Build Status](https://travis-ci.org/mdlayher/imagegrid.svg?branch=master)](https://travis-ci.org/mdlayher/imagegrid) [![GoDoc](https://godoc.org/github.com/mdlayher/imagegrid?status.svg)](https://godoc.org/github.com/mdlayher/imagegrid) [![Go Report Card](https://goreportcard.com/badge/github.com/mdlayher/imagegrid)](https://goreportcard.com/report/github.com/mdlayher/imagegrid)
=========

Package imagegrid enables composing one or more images into a single image,
using a tiled grid pattern.  MIT Licensed.

Example
-------

The `cmd/gophergrid` utility demonstrates usage of package `imagegrid`,
by generating a tiled grid of Go gophers.

```
$ gophergrid -h
Usage of gophergrid:
  -n int
        number of images per row or column (default 5)

By default, gophergrid will print a PNG image to stdout.

It is recommended to redirect stdout to a file or pipe.
$ gophergrid -n 5 > gopher.png
```

![gopher](https://cloud.githubusercontent.com/assets/1926905/26433716/577f99a2-40d2-11e7-9c5d-85df9a7d9ec8.png)

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
The design is licensed under the Creative Commons 3.0 Attributions license.
Read this article for more details: https://blog.golang.org/gopher
