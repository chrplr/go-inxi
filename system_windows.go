//go:build windows

package main

import (
	"os"
	"path/filepath"
	"time"
)

func printSystem() {
	host, _ := os.Hostname()

	info := first(wmicGet("OS", "Caption", "Version", "LastBootUpTime"))

	osName := info["Caption"]
	kernel := info["Version"]

	uptime := ""
	if boot, err := parseWMIDatetime(info["LastBootUpTime"]); err == nil {
		uptime = fmtDuration(time.Since(boot))
	}

	printSection("System",
		[]string{
			kv("Host", host),
			kv("OS", osName),
			kv("Kernel", kernel),
			kv("Uptime", uptime),
			kv("Shell", currentShell()),
		},
	)
}

func currentShell() string {
	// PowerShell sets PSModulePath; cmd.exe does not.
	if os.Getenv("PSModulePath") != "" {
		return "PowerShell"
	}
	if comspec := os.Getenv("COMSPEC"); comspec != "" {
		return filepath.Base(comspec)
	}
	return "cmd.exe"
}
