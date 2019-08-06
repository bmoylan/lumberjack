package lumberjack

import (
	"os"
	"syscall"
	"time"
)

func ctime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Win32FileAttributeData)
	ctime := time.Unix(0, stat.CreationTime.Nanoseconds())
	return ctime
}
