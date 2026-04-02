//go:build windows

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printCPU() {
	rows := wmicGet("cpu",
		"Name", "NumberOfCores", "NumberOfLogicalProcessors",
		"CurrentClockSpeed", "MaxClockSpeed",
	)
	if len(rows) == 0 {
		return
	}

	model := strings.Join(strings.Fields(rows[0]["Name"]), " ")
	var totalCores, totalThreads int
	var totalMHz, maxMHz float64

	for _, r := range rows {
		if n, err := strconv.Atoi(r["NumberOfCores"]); err == nil {
			totalCores += n
		}
		if n, err := strconv.Atoi(r["NumberOfLogicalProcessors"]); err == nil {
			totalThreads += n
		}
		if mhz, err := strconv.ParseFloat(r["CurrentClockSpeed"], 64); err == nil {
			totalMHz += mhz
		}
		if mhz, err := strconv.ParseFloat(r["MaxClockSpeed"], 64); err == nil && mhz > maxMHz {
			maxMHz = mhz
		}
	}

	topology := ""
	switch {
	case totalCores > 0 && totalThreads > totalCores:
		topology = fmt.Sprintf("%d cores / %d threads", totalCores, totalThreads)
	case totalCores > 0:
		topology = fmt.Sprintf("%d cores", totalCores)
	}

	speed := ""
	if avg := totalMHz / float64(len(rows)); avg > 0 {
		speed = fmt.Sprintf("%.0f MHz", avg)
		if maxMHz > 0 {
			speed += fmt.Sprintf(" (max: %.0f)", maxMHz)
		}
	}

	printSection("CPU",
		[]string{
			kv("Model", model),
			kv("Info", topology),
			kv("Speed", speed),
		},
	)
}
