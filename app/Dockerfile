FROM golang:1.20.5-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o app-presensi /app/main.go


FROM alpine

RUN mkdir /app

COPY --from=builder /app/app-presensi /app

CMD [ "/app/app-presensi" ]