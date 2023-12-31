FROM golang:1.20 as builder

COPY go.mod src/go.mod
COPY go.sum src/go.sum
COPY cmd src/cmd
RUN cd src/ && go mod download

RUN cd src && CGO_ENABLED=false go build -tags osusergo,netgo -o /github cmd/*.go

FROM gcr.io/distroless/static-debian11

COPY --from=builder /github /bin/github
CMD ["/bin/github"]