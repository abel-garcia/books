FROM golang:1.17-alpine AS builder
WORKDIR /go/src/books
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/books cmd/main.go


FROM golang:1.17-alpine
COPY --from=builder /go/src/books/bin bin/.
COPY --from=builder /go/src/books/docs docs/.
COPY --from=builder /go/src/books/books.env .
CMD ["./bin/books"]