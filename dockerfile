FROM golang:1.25-bullseye

# Install dependencies
RUN apt-get update && apt-get install -y python3 python3-pip ffmpeg \
    && pip3 install -U yt-dlp

# Set working directory
WORKDIR /app

# Copy Go files
COPY . .

# Build Go binary
RUN go build -o app .

# Expose port
EXPOSE 10000

# Run app
CMD ["./app"]
