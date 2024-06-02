FROM golang:1.18-alpine3.14
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG EnvironmentVariable
RUN go mod download && go build -o main ./cmd
CMD /app/main
