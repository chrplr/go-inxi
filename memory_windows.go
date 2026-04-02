//go:build windows

package main

import (
	"fmt"
	"strconv"
)

func printMemory() {
	osInfo := first(wmicGet("OS", "TotalVisibleMemorySize", "FreePhysicalMemory"))
	// Values are in KiB.
	totalKB, _ := strconv.ParseInt(osInfo["TotalVisibleMemorySize"], 10, 64)
	freeKB, _ := strconv.ParseInt(osInfo["FreePhysicalMemory"], 10, 64)
	if totalKB == 0 {
		return
	}
	usedKB := totalKB - freeKB
	pct := float64(usedKB) / float64(totalKB) * 100
	ram := fmt.Sprintf("total: %s  used: %s (%.1f%%)",
		fmtBytes(totalKB*1024), fmtBytes(usedKB*1024), pct)

	row := []string{kv("RAM", ram)}

	// Page file (Windows equivalent of swap).
	// AllocatedBaseSize and CurrentUsage are in MiB.
	pf := first(wmicPath("Win32_PageFileUsage", "AllocatedBaseSize", "CurrentUsage"))
	if pfTotal, err := strconv.ParseInt(pf["AllocatedBaseSize"], 10, 64); err == nil && pfTotal > 0 {
		pfUsed, _ := strconv.ParseInt(pf["CurrentUsage"], 10, 64)
		pfPct := float64(pfUsed) / float64(pfTotal) * 100
		swap := fmt.Sprintf("total: %s  used: %s (%.1f%%)",
			fmtBytes(pfTotal<<20), fmtBytes(pfUsed<<20), pfPct)
		row = append(row, kv("PageFile", swap))
	}

	printSection("Memory", row)
}
