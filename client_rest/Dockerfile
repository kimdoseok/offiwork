FROM golang:1.14.5-alpine as builder
RUN apk update
RUN apk add --no-cache bash git gcc libc-dev

WORKDIR /app
COPY . /app

RUN go build -o app .

# production environment
FROM alpine
COPY --from=builder /app/app /app/app
COPY offiworks.json /app/offiworks.json
#COPY model.conf /app/model.conf
#COPY policy.csv /app/policy.csv
WORKDIR /app
RUN mkdir storage
CMD ["./app"]