load("//tools:defs.bzl", "bzl_library", "go_library", "go_test", "select_goarch", "select_goos")
load("//tools/nogo:defs.bzl", "nogo_objdump_tool", "nogo_stdlib", "nogo_target")

package(licenses = ["notice"])

exports_files(["config-schema.json"])

nogo_target(
    name = "target",
    goarch = select_goarch(),
    goos = select_goos(),
    visibility = ["//visibility:public"],
)

nogo_objdump_tool(
    name = "objdump_tool",
    visibility = ["//visibility:public"],
)

nogo_stdlib(
    name = "stdlib",
    visibility = ["//visibility:public"],
)

go_library(
    name = "nogo",
    srcs = [
        "analyzers.go",
        "build.go",
        "config.go",
        "findings.go",
        "nogo.go",
    ],
    nogo = False,
    visibility = ["//:sandbox"],
    deps = [
        "//tools/checkescape",
        "//tools/checklinkname",
        "//tools/checklocks",
        "//tools/checkunsafe",
        "//tools/nogo/objdump",
        "//tools/worker",
        "@co_honnef_go_tools//staticcheck:go_default_library",
        "@co_honnef_go_tools//stylecheck:go_default_library",
        "@org_golang_x_tools//go/analysis:go_default_library",
        "@org_golang_x_tools//go/analysis/internal/facts:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/asmdecl:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/assign:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/atomic:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/bools:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/buildtag:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/cgocall:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/composite:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/copylock:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/errorsas:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/httpresponse:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/loopclosure:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/lostcancel:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/nilfunc:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/nilness:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/printf:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/shadow:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/shift:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/stdmethods:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/stringintconv:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/structtag:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/tests:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unmarshal:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unreachable:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unsafeptr:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unusedresult:go_default_library",
        "@org_golang_x_tools//go/gcexportdata:go_default_library",
        "@org_golang_x_tools//go/types/objectpath:go_default_library",
    ],
)

go_test(
    name = "nogo_test",
    srcs = ["config_test.go"],
    library = ":nogo",
)

bzl_library(
    name = "defs_bzl",
    srcs = ["defs.bzl"],
    visibility = ["//visibility:private"],
)
