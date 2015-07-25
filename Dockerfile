FROM golang:cross

ENV CGO_ENABLED 0

# Recompile the standard library without CGO
RUN go install -a std

RUN apt-get install -y -q git

# Declare the maintainer
MAINTAINER Millipede Team <business@millipede.io>

# For convenience, set an env variable with the path of the code
ENV APP_DIR /go/src/github.com/moul/millipede-go
WORKDIR $APP_DIR

ADD . /go/src/github.com/moul/millipede-go


# Compile the binary and statically link
RUN  GOOS=darwin   GOARCH=amd64          go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Darwin-x86_64
RUN  GOOS=darwin   GOARCH=386            go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Darwin-i386
RUN  GOOS=linux    GOARCH=386            go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Linux-i386
RUN cp /go/bin/millipede-Linux-i386 /go/bin/millipede-Linux-x86_64
RUN  GOOS=linux    GOARCH=arm   GOARM=5  go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Linux-arm
RUN  GOOS=linux    GOARCH=arm   GOARM=6  go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Linux-armv6
RUN  GOOS=linux    GOARCH=arm   GOARM=7  go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Linux-armv7
RUN  GOOS=freebsd  GOARCH=amd64          go build -a -v -ldflags '-w -s'    -o /go/bin/millipede-Freebsd-x86_64
RUN  GOOS=freebsd  GOARCH=386            go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Freebsd-i386
RUN  GOOS=freebsd  GOARCH=arm   GOARM=5  go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Freebsd-arm
RUN  GOOS=windows  GOARCH=amd64          go build -a -v -ldflags '-w -s'    -o /go/bin/millipede-Windows-x86_64.exe
RUN  GOOS=linux    GOARCH=amd64          go build -a -v -ldflags '-w -s'    -o /go/bin/millipede-Linux-x86_64

#RUN GOOS=openbsd  GOARCH=amd64          go build -a -v -ldflags '-w -s'    -o /go/bin/millipede-Openbsd-x86_64
#RUN GOOS=openbsd  GOARCH=386            go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Openbsd-i386
#RUN GOOS=openbsd  GOARCH=arm   GOARM=5  go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Openbsd-arm
#RUN GOOS=windows  GOARCH=386            go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Windows-i386.exe
#RUN GOOS=windows  GOARCH=arm   GOARM=5  go build -a -v -ldflags '-d -w -s' -o /go/bin/millipede-Windows-arm.exe
