FROM golang

WORKDIR /api

RUN mkdir -p public/uploads

COPY . .

RUN go build -o main .

CMD ["./main"]
