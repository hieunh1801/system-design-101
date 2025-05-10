# Stage 1: Biên dịch ứng dụng
FROM golang:tip-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go get
RUN  go build -o myapp .

# Stage 2: Tạo image nhỏ nhất
FROM scratch
COPY --from=builder /app/myapp /myapp
ENTRYPOINT ["/myapp"]