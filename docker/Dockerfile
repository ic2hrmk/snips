FROM golang:1.11-alpine3.9 as builder

WORKDIR /go/src/github.com/ic2hrmk/snips

RUN apk add git dep

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .

RUN _PROJECT=/go/src/github.com/ic2hrmk/snips \
    _VERSION=0.1.0 \
    _STAGING=$(git branch | grep \* | cut -d ' ' -f2) \
    _GITHASH=$(git rev-parse --short HEAD) \
    _BUILD_AT=$(date +%Y-%m-%dT%T%z) \
    _LDFLAGS="-X ${_PROJECT}/shared/version.Version=${_VERSION} \
              -X ${_PROJECT}/shared/version.Staging=${_STAGING} \
              -X ${_PROJECT}/shared/version.GitHash=${_GITHASH} \
              -X ${_PROJECT}/shared/version.BuildAt=${_BUILD_AT}" && \
    CGO_ENABLED=0 \
    go build -ldflags "${_LDFLAGS}" -o /bin/myapp cmd/main.go

FROM alpine:3.9

COPY --from=builder /build/myapp /app/
ENTRYPOINT ["/app/myapp"]
