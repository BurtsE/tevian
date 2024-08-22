FROM golang:1.22-alpine AS builder

WORKDIR /build



ADD go.mod .
ADD go.sum .
RUN go mod download
RUN pwd


COPY cmd cmd
COPY internal internal

RUN GOOS=linux go build -o app ./cmd

FROM golang:1.22-alpine
WORKDIR /root/

COPY configs configs

ENV FACE_CLOUD_LOGIN=${FACE_CLOUD_LOGIN}
ENV FACE_CLOUD_PASSWORD=${FACE_CLOUD_PASSWORD}
ENV FACE_CLOUD_URL=${FACE_CLOUD_URL}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}

ENV LOGIN=${LOGIN}
ENV PASSWORD=${PASSWORD}




COPY --from=builder /build/app .


CMD ["./app"]