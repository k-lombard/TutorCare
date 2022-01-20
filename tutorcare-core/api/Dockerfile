FROM golang:latest
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$PATH
RUN mkdir /app
COPY . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
EXPOSE 8080
CMD ["go","run", "main.go"]
