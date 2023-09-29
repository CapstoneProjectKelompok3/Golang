FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "ecci"
RUN go build -o ecci

# run executable file
CMD ["./ecci"]