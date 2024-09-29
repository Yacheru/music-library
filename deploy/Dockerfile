FROM golang:alpine AS builder

WORKDIR /music-library

ADD ../go.mod .

COPY . .

RUN go build -o ./build/music.exe ./cmd/main.go

FROM alpine

WORKDIR /music-library

COPY ./configs/.env /music-library/configs/.env
COPY --from=builder /music-library/build /music-library

CMD ["/music-library/music.exe"]