//go:build windows

package main

func printAudio() {
	devices := wmicPath("Win32_SoundDevice", "Name", "Manufacturer", "DriverVersion")
	if len(devices) == 0 {
		return
	}
	lines := make([][]string, len(devices))
	for i, d := range devices {
		name := d["Name"]
		if mfr := d["Manufacturer"]; mfr != "" && mfr != name {
			name = mfr + " " + name
		}
		lines[i] = []string{kv("Card", name), kv("Driver", d["DriverVersion"])}
	}
	printSection("Audio", lines...)
}
