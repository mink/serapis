FROM golang:1.16-alpine

WORKDIR /serapis
COPY . .

RUN go get -d -v
RUN go install -v

CMD ["serapis"]
