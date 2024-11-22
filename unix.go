//go:build !windows
// +build !windows

package main

import (
	"devhosts/sudo"
	"os"
	"slices"
	"strings"
)

const (
	hostsPath = "/etc/hosts"
	loopback  = "127.0.0.1"
)

var (
	hostsContent     string
	localhostAliases = []string{"localhost", loopback, "localhost.localdomain", "broadcasthost"}
)

func init() {
	hostsContentB, err := os.ReadFile(hostsPath)
	if err != nil {
		panic(err)
	}
	hostsContent = string(hostsContentB)
}

func AddHosts(hosts map[string]string) error {
	sb := strings.Builder{}
	sb.WriteString(hostsContent)
	sb.WriteString("\n\n# managed by devhosts\n")
	for host, target := range hosts {
		if slices.Contains(localhostAliases, target) {
			sb.WriteString(loopback)
		} else {
			sb.WriteString(target)
		}
		sb.WriteString(" ")
		sb.WriteString(host)
		sb.WriteString("\n")
	}

	newHostsContent := sb.String()
	return sudo.WriteFile(hostsPath, []byte(newHostsContent), 0644)
}

func ResetHosts() error {
	return sudo.WriteFile(hostsPath, []byte(hostsContent), 0644)
}
