git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.7.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()

go_register_toolchains()

go_repository(
    name = "com_github_spf13_cobra",
    commit = "1be1d2841c773c01bee8289f55f7463b6e2c2539",
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_rs_zerolog",
    commit = "1251b38a892a7049cc68c578448ee5313ba8caf8",
    importpath = "github.com/rs/zerolog",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "f15c970de5b76fac0b59abb32d62c17cc7bed265",
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_labstack_echo",
    commit = "0473c51f1dbd83487effce00702571d19033a6e5",
    importpath = "github.com/labstack/echo",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    commit = "462fda1f11d8cad3660e52737b8beefd27acfb3f",
    importpath = "github.com/kelseyhightower/envconfig",
)

go_repository(
    name = "com_github_spf13_cobra",
    commit = "1be1d2841c773c01bee8289f55f7463b6e2c2539",
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "4c012f6dcd9546820e378d0bdda4d8fc772cdfea",
    importpath = "github.com/spf13/pflag",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "365904b0f3154c6e11a9cf541c9803d1dca0445a",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "com_github_labstack_gommon",
    commit = "57409ada9da0f2afad6664c49502f8c50fbd8476",
    importpath = "github.com/labstack/gommon",
)

go_repository(
    name = "com_github_influxdata_influxdb",
    commit = "c59c4b231ecb38a78508b0167bf8f7151c4c418e",
    importpath = "github.com/influxdata/influxdb",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    commit = "dbeaa9332f19a944acb5736b4456cfcc02140e29",
    importpath = "github.com/dgrijalva/jwt-go",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    commit = "dcecefd839c4193db0d35b88ec65b4c12d360ab0",
    importpath = "github.com/valyala/fasttemplate",
)

go_repository(
    name = "com_github_dghubble_sling",
    commit = "80ec33c6152a53edb5545864ca37567b506c4ca5",
    importpath = "github.com/dghubble/sling",
)

go_repository(
    name = "com_github_JustinTulloss_firebase",
    commit = "da65bf12d9019bd021f4583130e3bd2ca1416434",
    importpath = "github.com/JustinTulloss/firebase",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "6fcc0c1fd9b620311d821b106a400b35dc95c497",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    commit = "e746df99fe4a3986f4d4f79e13c1e0117ce9c2f7",
    importpath = "github.com/valyala/bytebufferpool",
)

go_repository(
    name = "com_github_ancientlore_go_avltree",
    commit = "4ba4b949e04ae520ba87de82bb82d8f3a6be7e11",
    importpath = "github.com/ancientlore/go-avltree",
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

# You *must* import the Go rules before setting up the go_image rules.
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
