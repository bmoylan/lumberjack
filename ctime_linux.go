package lumberjack

import (
	"os"
	"syscall"
	"time"
)

func ctime(info os.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return ctime
}
