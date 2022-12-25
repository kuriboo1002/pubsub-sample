FROM golang:1.19-alpine

WORKDIR /

COPY . .

# COPY go.mod ./
# COPY go.sum ./

RUN go mod download

# COPY *.go ./

RUN go build -o /pull

CMD [ "/pull" ]