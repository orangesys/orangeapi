[![CircleCI](https://circleci.com/gh/orangesys/orangeapi.svg?style=svg)](https://circleci.com/gh/orangesys/orangeapi)
[![](https://images.microbadger.com/badges/image/orangesys/alpine-orangeapi.svg)](https://microbadger.com/images/orangesys/alpine-orangeapi "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/orangesys/alpine-orangeapi.svg)](https://microbadger.com/images/orangesys/alpine-orangeapi "Get your own version badge on microbadger.com")
## Orangesys new API with Sling

# Prerequisites
### set env
- FirebaseAuth
- FirebaseURL
- KONG_URL
- KONG_HOST
- KONG_PORT


# Development
Go 1.7 or higher is required
```
make deps
make build
bin/orangeapi
```