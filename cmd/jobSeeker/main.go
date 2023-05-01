// name    :	main.go
// version :	0.0.1
// date    :	20230428
// author  :	Leam Hall
// desc    :    jobSeeker app

package main

import (
    "errors"
    "flag"
    "fmt"
    "os" 
    "strings"

    "github.com/LeamHall/jobSeeker"
)


const jobFile = "jobSeeker_jobs.data"
const pocFile = "jobSeeker_pocs.data"

func main() {
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "%s tool.\n", os.Args[0])
        fmt.Fprintf(flag.CommandLine.Output(), "Job add format\n\t\"job;title;active(y/n);notes;company;url;pocId;firstContact;lastContact\"\n")
        fmt.Fprintf(flag.CommandLine.Output(), "POC add format\n\t\"poc;name;notes;company;email;phone;firstContact;lastContact\"\n\n")
        fmt.Fprintln(flag.CommandLine.Output(), "Usage:")
        flag.PrintDefaults()
    }

    add := flag.String("add", "", "Double quoted line to add")
    jobs    := flag.Bool("jobs", false, "Look at jobs")
    pocs    := flag.Bool("pocs", false, "Look at POCs")
    search  := flag.String("search", "", "String to search for")

    flag.Parse()
    
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

    switch {
    case len(*search) > 0:
        if !*jobs && !*pocs {
            fmt.Println("Please specify --jobs and/or --pocs to search through")
            os.Exit(1)
        }
        if *jobs {
            jobData := jobSeeker.Search(jobData, strings.ToLower(*search))
            if len(jobData) > 0 {
                for _, j := range jobData {
                    fmt.Println(j)
                }
            }
        }
        if *pocs {
            pocData := jobSeeker.Search(pocData, strings.ToLower(*search))
            if len(pocData) > 0 {
                for _, p := range pocData {
                    fmt.Println(p)
                }
            }
        }
 
    case len(*add) > 0:
        newData := make([]string, 10)
        addData, err := jobSeeker.FieldsFromLine(*add, ";")
        if err != nil {
            fmt.Printf("Can't convert %s to good data: %s\n", *add, err)
        } 
        count := copy(newData, addData)
        if count < len(addData) {
            fmt.Printf("Not sure why the copy failed. Just did %d", count)
        }
        if jobSeeker.InputType(*add) == "job" {
            j   := jobSeeker.Job{}
            j.JBuilder(newData)
            jobData, err = jobSeeker.Add(j.String, ";", jobData)
            if err != nil {
                fmt.Printf("Could not add %s to jobs\n", newData)
                os.Exit(1)
            }
            err = jobSeeker.WriteFile(jobFile, jobData)
            if err != nil {
                fmt.Printf("Could not write to job file: %s\n", err)
            } 
            if len(jobData) > 0 {
                for _, line := range jobData {
                    fmt.Println(line)
                }
            }
        } else if jobSeeker.InputType(*add) == "poc" {
            p   := jobSeeker.POC{}
            p.PBuilder(newData)
            pocData, err = jobSeeker.Add(p.String, ";", pocData)
            if err != nil {
                fmt.Printf("Could not add %s to pocs\n", *add)
                os.Exit(1)
            }
            err = jobSeeker.WriteFile(pocFile, pocData)
            if err != nil {
                fmt.Printf("Could not write to poc file: %s\n", err)
            } 
            if len(pocData) > 0 {
                for _, line := range pocData {
                    fmt.Println(line)
                }
            }
        }
    default:
        fmt.Println("Not sure what to do here.")
        os.Exit(1)
    
    } 
    
    //if len(jobData) > 0 {
    //    for _, line := range jobData {
    //        fmt.Println(line)
    //    }
    //}
    //if len(pocData) > 0 {
    //    for _, line := range pocData {
    //        fmt.Println(line)
    //    }
    //}
}


	
