package log

import "github.com/fatih/color"

func Print(text, status string) {
	switch status {
	case "info":
		color.New(color.FgGreen).Println("[+] " + text)
	case "warn":
		color.New(color.FgYellow).Println("[!] " + text)
	case "error":
		color.New(color.FgRed).Println("[-] " + text)
	default:
		color.New(color.FgCyan).Println("[*] " + text)
	}
}
