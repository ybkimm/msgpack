load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
	name = "msgpack",
	srcs = [
		"./array.go",
		"./const.go",
		"./decode.go",
		"./decode_array.go",
		"./decode_binary.go",
		"./decode_bool.go",
		"./decode_ext.go",
		"./decode_float32.go",
		"./decode_float64.go",
		"./decode_int.go",
		"./decode_int16.go",
		"./decode_int32.go",
		"./decode_int64.go",
		"./decode_int8.go",
		"./decode_json.go",
		"./decode_map.go",
		"./decode_string.go",
		"./decode_time.go",
		"./decode_uint.go",
		"./decode_uint16.go",
		"./decode_uint32.go",
		"./decode_uint64.go",
		"./decode_uint8.go",
		"./encode.go",
		"./encode_array.go",
		"./encode_binary.go",
		"./encode_bool.go",
		"./encode_ext.go",
		"./encode_float32.go",
		"./encode_float64.go",
		"./encode_int.go",
		"./encode_int16.go",
		"./encode_int32.go",
		"./encode_int64.go",
		"./encode_int8.go",
		"./encode_json.go",
		"./encode_map.go",
		"./encode_string.go",
		"./encode_time.go",
		"./encode_uint.go",
		"./encode_uint16.go",
		"./encode_uint32.go",
		"./encode_uint64.go",
		"./encode_uint8.go",
		"./errors.go",
		"./examples/main.go",
		"./extension.go",
		"./map.go",
		"./nullable.go",
		"./time.go",
		"./util.go",
	],
	importpath = "github.com/ybkimm/mono/packages/encoding/msgpack",
	visivility = [ "//visibility:public" ],
	deps = [],
)

go_test(
	name = "test",
	srcs = [
		"./benchmark_test.go",
		"./decode_array_test.go",
		"./decode_binary_test.go",
		"./decode_bool_test.go",
		"./decode_ext_test.go",
		"./decode_float32_test.go",
		"./decode_float64_test.go",
		"./decode_int16_test.go",
		"./decode_int32_test.go",
		"./decode_int64_test.go",
		"./decode_int8_test.go",
		"./decode_int_test.go",
		"./decode_json_test.go",
		"./decode_map_test.go",
		"./decode_string_test.go",
		"./decode_test.go",
		"./decode_time_test.go",
		"./decode_uint16_test.go",
		"./decode_uint32_test.go",
		"./decode_uint64_test.go",
		"./decode_uint8_test.go",
		"./decode_uint_test.go",
		"./encode_array_test.go",
		"./encode_binary_test.go",
		"./encode_bool_test.go",
		"./encode_ext_test.go",
		"./encode_float32_test.go",
		"./encode_float64_test.go",
		"./encode_int16_test.go",
		"./encode_int32_test.go",
		"./encode_int64_test.go",
		"./encode_int8_test.go",
		"./encode_int_test.go",
		"./encode_json_test.go",
		"./encode_map_test.go",
		"./encode_string_test.go",
		"./encode_test.go",
		"./encode_time_test.go",
		"./encode_uint16_test.go",
		"./encode_uint32_test.go",
		"./encode_uint64_test.go",
		"./encode_uint8_test.go",
		"./encode_uint_test.go",
	],
)
