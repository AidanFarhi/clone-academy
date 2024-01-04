FROM golang:1.20
WORKDIR /clone-academy
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY app/repository app/repository
COPY app/handlers app/handlers
COPY app/service app/service
COPY app/app.go app/app.go
COPY go.mod .
COPY main.go .
RUN go build main.go
COPY app/web app/web
CMD [ "./main" ]