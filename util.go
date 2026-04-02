package main

import (
	"fmt"
	"strings"
	"time"
)

func fmtDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	mins := int(d.Minutes()) % 60
	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if mins > 0 || len(parts) == 0 {
		parts = append(parts, fmt.Sprintf("%dm", mins))
	}
	return strings.Join(parts, " ")
}

func fmtBytes(b int64) string {
	const (
		KiB = 1024
		MiB = 1024 * KiB
		GiB = 1024 * MiB
		TiB = 1024 * GiB
	)
	switch {
	case b >= TiB:
		return fmt.Sprintf("%.2f TiB", float64(b)/TiB)
	case b >= GiB:
		return fmt.Sprintf("%.2f GiB", float64(b)/GiB)
	case b >= MiB:
		return fmt.Sprintf("%.1f MiB", float64(b)/MiB)
	default:
		return fmt.Sprintf("%d KiB", b/KiB)
	}
}
