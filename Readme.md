github.com/ybkimm/msgpack
=========================
[![godoc][godoc_badge]][godoc]
[![MIT License][license_badge]][license]

Another MessagePack encoding/decoding library.

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

<!-- For example, please see [examples/main.go](./examples/main.go). -->

Benchmark
=========
Testing PC spec:
* AMD Ryzen 5 2400G (3.90 GHz, 4 Cores, 8 Threads)
* 8GB RAM (Samsung DDR4)
* Windows 10 Pro Build 18362.418
* WSL

```plaintext
BenchmarkDecoder_decodeMap-8     	 5990668	       194 ns/op	      64 B/op	       1 allocs/op
BenchmarkDecoder_decodeArray-8   	 2336313	       512 ns/op	     320 B/op	       2 allocs/op
BenchmarkEncoder_encodeMap-8     	 3324308	       362 ns/op	     640 B/op	       3 allocs/op
BenchmarkEncoder_encodeArray-8   	 2410296	       492 ns/op	     592 B/op	       2 allocs/op
```

For benchmark data, see [msgpack_test.go](./msgpack_test.go).

Todo List
=========
* Optimize JSON conversion
* 100% coverage
* New benchmark data
* Examples

License
=======
MIT License. See [License](license).

Some code snippets are comes from [gojay](gojay) - See [it's license](gojay_license).

[godoc]:         https://godoc.org/github.com/ybkimm/msgpack
[godoc_badge]:   https://img.shields.io/badge/godoc-reference-blue.svg
[license]:       ./License
[license_badge]: https://img.shields.io/badge/license-MIT-green.svg
[map]:           https://godoc.org/github.com/ybkimm/msgpack#Map
[array]:         https://godoc.org/github.com/ybkimm/msgpack#Array
[extension]:     https://godoc.org/github.com/ybkimm/msgpack#Extension
[fromjson]:      https://godoc.org/github.com/ybkimm/msgpack#FromJSON
[tojson]:        https://godoc.org/github.com/ybkimm/msgpack#ToJSON

[gojay]:         https://github.com/francoispqt/gojay
[gojay_license]: https://github.com/francoispqt/gojay/blob/decd89f/LICENSE
