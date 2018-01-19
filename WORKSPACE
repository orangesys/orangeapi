http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.8.1/rules_go-0.8.1.tar.gz",
    sha256 = "90bb270d0a92ed5c83558b2797346917c46547f6f7103e648941ecdb6b9d0e72",
)
load("@io_bazel_rules_go//go:def.bzl",
  "go_rules_dependencies",
  "go_register_toolchains",
  "go_repository")
go_rules_dependencies()
go_register_toolchains()

# You *must* import the Go rules before setting up the bazel_gazelle rules.
http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.8/bazel-gazelle-0.8.tar.gz",
    sha256 = "e3dadf036c769d1f40603b86ae1f0f90d11837116022d9b06e4cd88cae786676",
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

# You *must* import the Go rules before setting up the go_image rules.
git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    commit = "8aeab63328a82fdb8e8eb12f677a4e5ce6b183b1",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

go_repository(
    name = "com_github_spf13_cobra",
    commit = "b95ab734e27d33e0d8fbabf71ca990568d4e2020",
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_rs_zerolog",
    commit = "b53826c57a6a1d8833443ebeacf1cfb62b229c64",
    importpath = "github.com/rs/zerolog",
)

go_repository(
    name = "com_github_labstack_echo",
    commit = "b6040409eeceaac178c95fe73498b5f992e6b668",
    importpath = "github.com/labstack/echo",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    commit = "462fda1f11d8cad3660e52737b8beefd27acfb3f",
    importpath = "github.com/kelseyhightower/envconfig",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "e881fd58d78e04cf6d0de1217f8707c8cc2249bc",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "4c012f6dcd9546820e378d0bdda4d8fc772cdfea",
    importpath = "github.com/spf13/pflag",
)

go_repository(
    name = "com_github_labstack_gommon",
    commit = "57409ada9da0f2afad6664c49502f8c50fbd8476",
    importpath = "github.com/labstack/gommon",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    commit = "dcecefd839c4193db0d35b88ec65b4c12d360ab0",
    importpath = "github.com/valyala/fasttemplate",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    commit = "dbeaa9332f19a944acb5736b4456cfcc02140e29",
    importpath = "github.com/dgrijalva/jwt-go",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "13931e22f9e72ea58bb73048bc752b48c6d4d4ac",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "com_github_JustinTulloss_firebase",
    commit = "da65bf12d9019bd021f4583130e3bd2ca1416434",
    importpath = "github.com/JustinTulloss/firebase",
)

go_repository(
    name = "com_github_dghubble_sling",
    commit = "80ec33c6152a53edb5545864ca37567b506c4ca5",
    importpath = "github.com/dghubble/sling",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "586e6dcca296d085100876beb6dbf02287a247f7",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    tag = "v0.0.2",
    # commit = "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_facebookgo_httpcontrol",
    commit = "ccde4420e1fee9af5da365cb76075b95683f39a6",
    importpath = "github.com/facebookgo/httpcontrol",
)

go_repository(
    name = "com_github_google_go_querystring",
    commit = "53e6ce116135b80d037921a7fdd5138cf32d7a8a",
    importpath = "github.com/google/go-querystring",
)

go_repository(
    name = "com_github_ancientlore_go_avltree",
    commit = "4ba4b949e04ae520ba87de82bb82d8f3a6be7e11",
    importpath = "github.com/ancientlore/go-avltree",
)

go_repository(
    name = "com_github_influxdata_influxdb",
    commit = "2169ad680e170bfbe0d732467523a50776d27f71",
    importpath = "github.com/influxdata/influxdb",
)
