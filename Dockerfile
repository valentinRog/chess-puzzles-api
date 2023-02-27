FROM golang

WORKDIR /chess-puzzles-api

COPY . .

RUN go build -o chess-puzzles-api

CMD ["./chess-puzzles-api"]