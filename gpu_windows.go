//go:build windows

package main

func printGraphics() {
	rows := wmicPath("Win32_VideoController", "Name", "DriverVersion")
	if len(rows) == 0 {
		return
	}
	lines := make([][]string, len(rows))
	for i, r := range rows {
		lines[i] = []string{
			kv("Card", r["Name"]),
			kv("Driver", r["DriverVersion"]),
		}
	}
	printSection("Graphics", lines...)
}
