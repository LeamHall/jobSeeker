// name    :	main_test.go
// version :	0.0.1
// date    :	20230428
// author  :	Leam Hall
// desc    :    Test jobSeeker app

package main_test

import (
    "fmt"
    "os"	
    "os/exec"	
    "runtime"
    "testing"

    //"github.com/LeamHall/jobSeeker"
)

var (
    binName     = "jobSeeker"
    jobFile     = "jobSeeker_jobs.test_data"
    pocFile     = "jobSeeker_pocs.test_data"
)

func TestMain(m *testing.M) {
    fmt.Println("Building tool...")

    if runtime.GOOS == "windows" {
        binName += ".exe"
    }
        
    build := exec.Command("go", "build", "-o", binName)
    if err := build.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Cannot build %s: %s", binName, err)
        os.Exit(1)
    }

    fmt.Println("Running tests...")
    result := m.Run()
    
    fmt.Println("Cleaning up...")
    os.Remove(binName)
    os.Remove(jobFile)
    os.Remove(pocFile)
    os.Exit(result)

}


