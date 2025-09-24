FROM golang:1.25-bullseye

# Install dependencies (yt-dlp + ffmpeg)
RUN apt-get update && apt-get install -y python3 python3-pip ffmpeg \
    && pip3 install -U yt-dlp \
    && ln -s /usr/local/bin/yt-dlp /usr/bin/yt-dlp


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
