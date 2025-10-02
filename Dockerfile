FROM amazonlinux:2023 as build

WORKDIR /app

# Install Go
RUN yum install -y golang

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 go build -o main main.go

# Copy Delve
COPY dlv_binary dlv

FROM public.ecr.aws/lambda/go:1

# Copy the binary and Delve
COPY --from=build /app/main .
COPY --from=build /app/dlv /tmp/lambci_debug_files/dlv

# Command to run the Lambda function
CMD ["./main"]