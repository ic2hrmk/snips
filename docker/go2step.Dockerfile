FROM golang:1.14.1-alpine3.11 AS BUILDER
ENV CGO_ENABLED=0
WORKDIR /app
RUN apk add git
COPY . .
RUN go build -a -o app ./microdemo-web/*.go

FROM alpine:3.11.5 AS RUNNER
WORKDIR /app
COPY --from=BUILDER /app/app .
RUN apk --no-cache add ca-certificates
ENV USER=microdemo
RUN addgroup ${USER} && \
    adduser --ingroup ${USER} --no-create-home --disabled-password ${USER}
USER ${USER}
ENV MICRO_SERVER_ADDRESS=:8080
EXPOSE 8080
ENTRYPOINT ["./app"]
