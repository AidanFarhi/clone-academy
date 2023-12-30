FROM golang:1.20
WORKDIR /app
COPY /src /app
RUN go build main.go
CMD [ "./main" ]