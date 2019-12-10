load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/ybkimm/mono/packages/encoding/msgpack/examples",
    visibility = ["//visibility:private"],
    deps = ["//packages/encoding/msgpack:go_default_library"],
)

go_binary(
    name = "examples",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
