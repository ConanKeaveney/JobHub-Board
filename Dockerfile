# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.12.7 as dev

# Install dependencies
# No Longer Needed
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/bson


# copy the local package files to the container workspace
ADD ./src /go/src/

# Setting up working directory
WORKDIR /go/src/

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o boards ./


# alpine image
FROM alpine:3.9.2 as prod

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

# Setting up working directory
WORKDIR /root/
# copy boards binary
COPY --from=dev /go/src/ .

EXPOSE 8080

# Run the boards microservice when the container starts.
ENTRYPOINT ["./boards"]
