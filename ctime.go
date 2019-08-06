// +build !darwin,!linux,!windows

package lumberjack

import (
	"os"
	"time"
)

// ctime returns the current time for platforms which do not support getting the change time from the filesystem.
func ctime(info os.FileInfo) time.Time {
	return time.Now()
}
