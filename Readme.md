github.com/ybkimm/msgpack
=========================
[![godoc][godoc_badge]][godoc]
[![MIT License][license_badge]][license]

Fast MessagePack encoding/decoding library.

This library doesn't use reflection, instead it relies on small interface:
[Object][object], [Array][array], and [Extension][extension].

Heavily inspired by [GoJay][gojay].

Features
========
* Encoding primitive types, arrays, maps, and `time.Time`.
* No reflection.
* [Extension interface for custom types][extension].
* JSON conversion: [FromJSON][fromjson], [ToJSON][tojson].

Get Started
===========
```
go get -u github.com/ybkimm/msgpack
```

[godoc]:         https://godoc.org/github.com/ybkimm/msgpack
[godoc_badge]:   https://img.shields.io/badge/godoc-reference-blue.svg
[license]:       ./License
[license_badge]: https://img.shields.io/badge/license-MIT-green.svg
[object]:        https://godoc.org/github.com/ybkimm/msgpack#Object
[array]:         https://godoc.org/github.com/ybkimm/msgpack#Array
[extension]:     https://godoc.org/github.com/ybkimm/msgpack#Extension
[fromjson]:      https://godoc.org/github.com/ybkimm/msgpack#FromJSON
[tojson]:        https://godoc.org/github.com/ybkimm/msgpack#ToJSON

[gojay]:         https://github.com/francoispqt/gojay