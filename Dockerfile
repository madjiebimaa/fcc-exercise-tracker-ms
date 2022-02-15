FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o fcc-exercise-tracker-ms

EXPOSE 3000

CMD ./fcc-exercise-tracker-ms