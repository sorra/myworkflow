FROM golang:1.16.0-alpine3.13
WORKDIR /
COPY . .
#RUN GOOS=linux GOARCH=amd64 go build -o myworkflow . # Need network proxy in China
CMD ["./myworkflow worker"]