package lumberjack

import (
	"os"
	"syscall"
	"time"
)

func ctime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	ctime := time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	return ctime
}
