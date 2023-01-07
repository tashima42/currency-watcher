FROM golang:1.19.1-alpine

# Install go releaser
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o currencywatcher.out
CMD [ "./currencywatcher.out"]