[![CircleCI](https://circleci.com/gh/orangesys/orangeapi.svg?style=svg)](https://circleci.com/gh/orangesys/orangeapi)
[![](https://images.microbadger.com/badges/image/orangesys/alpine-orangeapi.svg)](https://microbadger.com/images/orangesys/alpine-orangeapi "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/orangesys/alpine-orangeapi.svg)](https://microbadger.com/images/orangesys/alpine-orangeapi "Get your own version badge on microbadger.com")

# orangesys

ornagesys api

# prerequisites

- Kubernetes 1.6
- Deploy helm Tiller server 2.5.x on kubernetes
- [Wheel](https://github.com/appscode/wheel)

# set environment value with run docker

- FirebaseAuth
- FirebaseURL
- KONG_URL
- KONG_HOST
- KONG_PORT

# development

Go 1.8 or higher is required

```bash
make deps
make build
```

# run

## linux

```bash
dist/orangeapi_linux-amd64
```

## osx

```bash
dist/orangeapi_darwin-amd64
```