package util

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/host"
)

// GetUpTime get uptime
func GetUpTime() string {
	uptime, _ := host.Uptime()
	fmt.Println("Total:", uptime, "seconds")

	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	totalUptime := strconv.FormatUint(days, 10) + " days, " + strconv.FormatUint(hours, 10) + " hours, " + strconv.FormatUint(minutes, 10) + " minutes"
	return totalUptime
}
