package main

import (
	"flag"
	"fmt"

	"github.com/M00NLIG7/ChopChopGo/maps/auditd"
	"github.com/M00NLIG7/ChopChopGo/maps/journald"
	"github.com/M00NLIG7/ChopChopGo/maps/syslog"
)

func main() {
	var target string
	var path string

	flag.StringVar(&target, "target", "auditd", "where should the scan be conducted")
	flag.StringVar(&path, "rules", "rules/linux/auditd", "where to pull the yaml rules youre applying")

	flag.Parse()
	a := `  ▄████▄   ██░ ██  ▒█████   ██▓███      ▄████▄   ██░ ██  ▒█████   ██▓███       ▄████  ▒█████  
 ▒██▀ ▀█  ▓██░ ██▒▒██▒  ██▒▓██░  ██▒   ▒██▀ ▀█  ▓██░ ██▒▒██▒  ██▒▓██░  ██▒    ██▒ ▀█▒▒██▒  ██▒
 ▒▓█    ▄ ▒██▀▀██░▒██░  ██▒▓██░ ██▓▒   ▒▓█    ▄ ▒██▀▀██░▒██░  ██▒▓██░ ██▓▒   ▒██░▄▄▄░▒██░  ██▒
 ▒▓▓▄ ▄██▒░▓█ ░██ ▒██   ██░▒██▄█▓▒ ▒   ▒▓▓▄ ▄██▒░▓█ ░██ ▒██   ██░▒██▄█▓▒ ▒   ░▓█  ██▓▒██   ██░
 ▒ ▓███▀ ░░▓█▒░██▓░ ████▓▒░▒██▒ ░  ░   ▒ ▓███▀ ░░▓█▒░██▓░ ████▓▒░▒██▒ ░  ░   ░▒▓███▀▒░ ████▓▒░
 ░ ░▒ ▒  ░ ▒ ░░▒░▒░ ▒░▒░▒░ ▒▓▒░ ░  ░   ░ ░▒ ▒  ░ ▒ ░░▒░▒░ ▒░▒░▒░ ▒▓▒░ ░  ░    ░▒   ▒ ░ ▒░▒░▒░ 
   ░  ▒    ▒ ░▒░ ░  ░ ▒ ▒░ ░▒ ░          ░  ▒    ▒ ░▒░ ░  ░ ▒ ▒░ ░▒ ░          ░   ░   ░ ▒ ▒░ 
 ░         ░  ░░ ░░ ░ ░ ▒  ░░          ░         ░  ░░ ░░ ░ ░ ▒  ░░          ░ ░   ░ ░ ░ ░ ▒  
 ░ ░       ░  ░  ░    ░ ░              ░ ░       ░  ░  ░    ░ ░                    ░     ░ ░  
 ░                                     ░                                                    
			By Keyboard Cowboys (M00NL1G7)
 `
	fmt.Println(a)
	switch target {
	case "auditd":
		auditd.Chop(path)
	case "syslog":
		syslog.Chop(path)
	case "journald":
		journald.Chop(path)
	default:
		fmt.Printf("Invalid Target Type")
		return
	}
}
