FROM golang:1.14.5-alpine as builder
RUN apk update
RUN apk add --no-cache bash git gcc libc-dev

RUN mkdir /app
COPY . /app
WORKDIR /app

#ENV GO111MODULE=off
ENV GOBIN /app/bin
RUN go build -o app .

# production environment
FROM alpine
COPY --from=builder /app/app /app/app
COPY cdidemo.json /app/cdidemo.json
COPY model.conf /app/model.conf
COPY policy.csv /app/policy.csv
WORKDIR /app
RUN mkdir log
CMD ["./app"]