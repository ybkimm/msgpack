github.com/ybkimm/msgpack
=========================
[![godoc][godoc_badge]][godoc]
[![MIT License][license_badge]][license]

Fast MessagePack encoding/decoding library.

This library doesn't use reflection, instead it relies on small interface:
[Map][map], [Array][array], and [Extension][extension].

Heavily inspired by [GoJay][gojay].

Features
========
* Encoding primitive types, arrays, maps, and `time.Time`.
* No reflection.
* [Extension interface for custom types][extension].
* JSON conversion: [FromJSON][fromjson], [ToJSON][tojson].
  * NotOptimizedYet™ - json-related functions are slow at this time.

Get Started
===========
```
go get -u github.com/ybkimm/msgpack
```

For example, please see [examples/main.go](./examples/main.go).

Benchmark
=========
Testing PC spec:
* AMD Ryzen 5 2400G (3.90 GHz, 4 Cores, 8 Threads)
* 8GB RAM (Samsung DDR4)
* Windows 10 Pro Build 18362.418

```plaintext
BenchmarkDecoder_DecodeBool-8             260881              4296 ns/op            4144 B/op          2 allocs/op
BenchmarkDecoder_decode-8                 222208              5005 ns/op            4224 B/op          4 allocs/op
BenchmarkEncoder_encodeMap-8              599858              2056 ns/op             576 B/op          2 allocs/op
BenchmarkEncoder_encodeArray-8            266655              4531 ns/op             576 B/op          2 allocs/op
```

Todo List
=========
* Optimizing JSON conversion, decoding. Maybe, sometime...
* Refactoring decode functions

License
=======
MIT License.

[godoc]:         https://godoc.org/github.com/ybkimm/msgpack
[godoc_badge]:   https://img.shields.io/badge/godoc-reference-blue.svg
[license]:       ./License
[license_badge]: https://img.shields.io/badge/license-MIT-green.svg
[map]:           https://godoc.org/github.com/ybkimm/msgpack#Object
[array]:         https://godoc.org/github.com/ybkimm/msgpack#Array
[extension]:     https://godoc.org/github.com/ybkimm/msgpack#Extension
[fromjson]:      https://godoc.org/github.com/ybkimm/msgpack#FromJSON
[tojson]:        https://godoc.org/github.com/ybkimm/msgpack#ToJSON

[gojay]:         https://github.com/francoispqt/gojay
