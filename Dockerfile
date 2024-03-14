FROM golang:1.22
WORKDIR /
COPY . .
#RUN GOOS=linux GOARCH=amd64 go build -o myworkflow . # Need network proxy in China
CMD ["./myworkflow"]