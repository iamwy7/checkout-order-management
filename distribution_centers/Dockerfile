# Link útil sobre as flags https://stackoverflow.com/questions/52640304/standard-init-linux-go190-exec-user-process-caused-no-such-file-or-directory
FROM golang:1.24-alpine AS builder
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -tags netgo -ldflags="-s -w" main.go

FROM scratch AS runner
WORKDIR /go/src/
COPY --from=builder /go/src .
ENTRYPOINT [ "./main"]