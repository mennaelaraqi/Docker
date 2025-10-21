FROM golang:1.20

# set work directory
WORKDIR /app

# copy application files
COPY . .


RUN go mod download
RUN go build -o main .


EXPOSE 3000


CMD ["./main"]
