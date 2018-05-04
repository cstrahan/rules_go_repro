workspace(name = "demo")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.11.1/rules_go-0.11.1.tar.gz",
    sha256 = "1868ff68d6079e31b2f09b828b58d62e57ca8e9636edff699247c9108518570b",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()
go_register_toolchains()
