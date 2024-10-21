FROM golang:alpine AS builder

ARG user=swadmin
RUN adduser -D -g '' $user

RUN mkdir /app

WORKDIR /app

RUN chown -R $user:$user /app

COPY go.mod .
COPY go.sum .

RUN go mod download


COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app .

RUN chmod 0555 /go/bin/app

# start from scratch
FROM scratch

WORKDIR /app


COPY --from=builder /etc/passwd /etc/passwd

USER $user

COPY --from=builder /go/bin/app ./app_binary

EXPOSE 3000

CMD ["./app_binary"]