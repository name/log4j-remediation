package backend

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/shirou/gopsutil/disk"
)

// Discover : Starts discovery of '.ear|.jar|.war' files
func Discover() []string {
	var detections []string
	partitions, _ := disk.Partitions(true)
	for _, partition := range partitions {
		log.Println("Discovering on " + partition.Mountpoint)
		bar := progressbar.NewOptions(-1,
			progressbar.OptionSetDescription("Discovering files"),
			progressbar.OptionShowCount(),
			progressbar.OptionSpinnerType(40),
		)
		e := filepath.Walk(partition.Mountpoint+"/", func(path string, f os.FileInfo, err error) error {
			bar.Add(1)
			if strings.Contains(path, "log4j-core-") && !strings.Contains(path, "2.17") {
				if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".ear") || strings.HasSuffix(path, ".war") {
					if isVulnerable(path) {
						detections = append(detections, path)
					}
				}
			}
			return nil
		})

		if e != nil {
			log.Println(e)
		}
		fmt.Println()
	}
	return detections
}

func isVulnerable(file string) bool {
	zip, _ := zip.OpenReader(file)
	for _, f := range zip.File {
		if strings.Contains(strings.ToLower(f.Name), "jndilookup.class") {
			return true
		}
	}
	return false
}
