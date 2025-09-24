FROM golang:1.25-bullseye

# Install dependencies (yt-dlp + ffmpeg)
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3 python3-pip ffmpeg ca-certificates \
    && pip3 install --no-cache-dir -U yt-dlp \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy go.mod & go.sum dulu (cache dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy all files
COPY . .

# Build Go binary
RUN go build -o app .

# Expose port untuk Render (Render biasanya set $PORT)
EXPOSE 10000

# Run app
CMD ["./app"]
