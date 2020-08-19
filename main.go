package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// MountInfo contains data about which drive to mount and where to mount it.
type MountInfo struct {
	DriveName string `json:"drive_name"`
	Media     string `json:"media"`
}

func main() {
	fmt.Println("Starting Auto-Mounting...\n")

	file, err := ioutil.ReadFile("/home/mdrokz/.mount_info.json")

	fmt.Println(string(file))

	var mountInfo *[]MountInfo

	json.Unmarshal(file, &mountInfo)

	var cmd *exec.Cmd

	// var output []byte

	var outputErr error

	for _, drive := range *mountInfo {
		cmd = exec.Command("sudo", "mount", drive.DriveName, drive.Media)
		_, outputErr = cmd.Output()

		if outputErr != nil {
			fmt.Printf("ERROR: Mounting of drive - %s at location - %s failed with %v ... \n", drive.DriveName, drive.Media, outputErr)
		} else {
			fmt.Printf("Successfully mounted drive - %s at location - %s...\n", drive.DriveName, drive.Media)
		}
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("AutoMount of drives completed...")
}
