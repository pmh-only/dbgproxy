FROM alpine as builder

RUN apk add go

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENV CGO_ENABLED=0

RUN go build -o ./dbgproxy

RUN chmod +x ./dbgproxy

# ---

FROM scratch as runtime 

USER 1000:1000

WORKDIR /app

COPY --from=builder --chown=1000:1000 /app/dbgproxy ./

ENTRYPOINT [ "/app/dbgproxy" ]
