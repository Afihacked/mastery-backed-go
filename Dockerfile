FROM golang:1.25-bullseye

# Install dependencies (ffmpeg + wget untuk ambil yt-dlp binary)
RUN apt-get update && apt-get install -y ffmpeg wget python3 && rm -rf /var/lib/apt/lists/*

# Install yt-dlp binary langsung (bukan pip)
RUN wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp \
    && chmod a+rx /usr/local/bin/yt-dlp

# Set working directory
WORKDIR /app

# Copy go.mod & go.sum dulu (cache dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy all project files
COPY . .

# Build Go binary
RUN go build -o app .

# Expose port untuk Render (Render set $PORT)
EXPOSE 10000

# Run app
CMD ["./app"]
