FROM golang:1.25-bullseye

# Install dependencies (yt-dlp + ffmpeg)
RUN apt-get update && apt-get install -y python3 ffmpeg curl \
    && curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/bin/yt-dlp \
    && chmod a+rx /usr/bin/yt-dlp

# Set working directory
WORKDIR /app

# Copy go.mod & go.sum dulu (cache dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy all files
COPY . .

# Build Go binary
RUN go build -o app .

# Expose port
EXPOSE 10000

# Run app
CMD ["./app"]
