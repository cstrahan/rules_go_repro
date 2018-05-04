package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_binary",
    "go_test",
)

cc_library(
    name = "cxxdep",
    srcs = ["dep.cc"],
    alwayslink = 1,
)

cc_library(
    name = "cxxlib",
    srcs = ["lib.cc"],
    hdrs = ["lib.h"],
    deps = [
        "cxxdep",
        ],
        alwayslink = 1 ,
)

cc_binary(
    name = "democc",
    srcs = ["main.cc"],
    deps = [
        "cxxlib",
    ],
)

go_binary(
    name = "demogo",
    srcs = [
        "main.go",
    ],
    cdeps = [
        ":cxxlib",
    ],
    cgo = True,
)
