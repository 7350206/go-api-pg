# --------- 1st step: does all steps and tools to compile needed app
FROM golang:1.19 AS builder

RUN mkdir /app
# add all content of curr dir
ADD . /app
WORKDIR /app

# while have access to go cli and source code, so make app binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# -------- 2nd step, execute binary from FAR LIGHTER image
FROM alpine:latest AS production
COPY --from=builder /app .
# execute
CMD ["./app"]

# benefit: final container size will far smaller in terms of mem&reqs
# check docker ps
