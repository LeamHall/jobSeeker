// name    :	main.go
// version :	0.0.1
// date    :	20230428
// author  :	Leam Hall
// desc    :    jobSeeker app

package main

import (
    "errors"
    "fmt"
    "os" 

    "github.com/LeamHall/jobSeeker"
)


const jobFile = "jobSeeker_jobs.data"
const pocFile = "jobSeeker_pocs.data"

func main() {

    jobData, err    := jobSeeker.DataFromFile(jobFile)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            _, err := os.Create(jobFile)
            if err != nil {
                fmt.Printf("Can't create %s file: %s\n", jobFile, err)
                os.Exit(1)
            }
        } else {
            fmt.Printf("Can't open %s file: %s\n", jobFile, err)
            os.Exit(1)
        }
    }

    pocData, err    := jobSeeker.DataFromFile(pocFile)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            _, err := os.Create(pocFile)
            if err != nil {
                fmt.Printf("Can't create %s file: %s\n", pocFile, err)
                os.Exit(1)
            }
        } else {
            fmt.Printf("Can't open %s file: %s\n", pocFile, err)
            os.Exit(1)
        }
    }

    if len(jobData) > 0 {
        for _, line := range jobData {
            fmt.Println(line)
        }
    }
    if len(pocData) > 0 {
        for _, line := range pocData {
            fmt.Println(line)
        }
    }
}


	
