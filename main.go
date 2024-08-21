package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	"tmp-cleaner/pkg"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func cleanOldFiles(dir string, days int) {
	pkg.Log.Infof("Cleaning old files in %s older than %d days\n", dir, days)

	now := time.Now()
	cutoff := now.AddDate(0, 0, -days)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			pkg.Log.Errorf("Error walking the path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && info.ModTime().Before(cutoff) {
			pkg.Log.Infof("Deleting: %s\n", path)
			return os.Remove(path)
		}
		return nil
	})
	if err != nil {
		pkg.Log.Errorf("Error cleaning old files: %v\n", err)
	}
}

func main() {
	tmpDirFlag := flag.String("tmpDir", getEnv("TMP_DIR", "/tmp"), "Temporary directory to clean")
	daysFlag := flag.Int("days", func() int {
		days, err := strconv.Atoi(getEnv("DAYS", "7"))
		if err != nil {
			return 7
		}
		return days
	}(), "Number of days to keep files")
	intervalFlag := flag.Int("interval", func() int {
		interval, err := strconv.Atoi(getEnv("INTERVAL", "86400"))
		if err != nil {
			return 86400
		}
		return interval
	}(), "Interval between cleanups in seconds")

	help := flag.Bool("h", false, "Show help")
	flag.Usage = func() {
		fmt.Println("Usage: tmp_cleaner [options]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || flag.NFlag() == 1 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
		flag.Usage()
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(*intervalFlag) * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				cleanOldFiles(*tmpDirFlag, *daysFlag)
			}
		}
	}()

	<-done
	pkg.Log.Info("Received exit signal, shutting down...")
}
