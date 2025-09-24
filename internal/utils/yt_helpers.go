package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type YtFormat struct {
	FormatID   string  `json:"format_id"`
	Ext        string  `json:"ext"`
	Resolution string  `json:"format_note"`
	Height     int     `json:"height"`
	Filesize   int64   `json:"filesize"`
	URL        string  `json:"url"`
}

type YtInfo struct {
	Title     string     `json:"title"`
	Thumbnail string     `json:"thumbnail"`
	Duration  int        `json:"duration"`
	URL       string     `json:"url"`
	Ext       string     `json:"ext"`
	Formats   []YtFormat `json:"formats"`
}

func ExtractInfo(url string, format string, cookies string) (map[string]interface{}, error) {
	args := []string{
		"-J", // JSON output
		"--no-warnings",
		"--no-check-certificate",
		"--ignore-errors",
		"--no-playlist",
	}

	if format != "" {
		args = append(args, "-f", format)
	}

	if cookies != "" {
		tmpFile := "/tmp/cookies.txt"
		_ = os.WriteFile(tmpFile, []byte(cookies), 0644)
		args = append(args, "--cookies", tmpFile)
	}

	args = append(args, url)

	// 🔍 cek lokasi yt-dlp di container
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		return map[string]interface{}{"success": false, "error": "yt-dlp not found in PATH"}, err
	}
	fmt.Println("yt-dlp path:", path)

	// 🔧 jalankan command
	cmd := exec.Command(path, args...)
	out, err := cmd.Output()
	if err != nil {
		return map[string]interface{}{"success": false, "error": err.Error()}, err
	}

	var info YtInfo
	if err := json.Unmarshal(out, &info); err != nil {
		return map[string]interface{}{"success": false, "error": err.Error()}, err
	}

	if format != "" {
		return map[string]interface{}{
			"success":   true,
			"title":     info.Title,
			"thumbnail": info.Thumbnail,
			"duration":  info.Duration,
			"url":       info.URL,
			"ext":       info.Ext,
		}, nil
	}

	return map[string]interface{}{
		"success":   true,
		"title":     info.Title,
		"thumbnail": info.Thumbnail,
		"duration":  info.Duration,
		"formats":   info.Formats,
	}, nil
}
