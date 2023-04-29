// name    :	jobSeeker.go
// version :	0.0.1
// date    :	20230427
// author  :	Leam Hall
// desc    :	Tools for tracking job applications

package jobSeeker

import (
    "os"
    "strconv"
    "strings"
    "time"
)


// Add adds a line of data to the proper dataset.
// It errors if the line does not meet standards.
func Add(s string, sep string, data []string) (newData []string, err error) {
    newData = data
    highestId, err   := HighestId(newData, sep)
    if err != nil {
        return newData, err
    } 
    fields, err := FieldsFromLine(s, sep)
    if err != nil {
        return newData, err
    }
    fields[0]   = strconv.Itoa(highestId + 1)
    newS        := strings.Join(fields[:], sep)
    newData = append(newData, newS)
    return newData, err
}

// DataFromFile takes a fileName string, and returns a slice of strings.
func DataFromFile(f string) (data []string, err error) {
    comments := []string{"#", "/", "*",}
    handle, err := os.ReadFile(f)
    if err != nil {
        return data, err
    }
    if len(handle) == 0 {
        return data, err
    }
    lineData := strings.Split(string(handle), "\n")
    CheckPrefix:
        for _, line := range lineData {
            line = strings.TrimSpace(line)
            if len(line) == 0 {
                continue
            }
            for _,c := range comments {
                if strings.HasPrefix(line, c) {
                    continue CheckPrefix
                }
            } 
            data = append(data, line)
        }
    return data, err
}

// FieldsFromLine takes a single line string and returns the data in a slice
//  of space trimmed fields, based on the given separator string.
func FieldsFromLine(line string, sep string) (data []string, err error) {
    err = nil
    rawData := strings.Split(line, sep)
    for _, datum := range rawData {
        data = append(data, strings.TrimSpace(datum))
    }
    return data, err
}

// Input type takes an input string and returns if it is a job or poc input,
//  returns nil otherwise.
func InputType(s string) (iType string) {
    s = strings.ToLower(s)
    if strings.HasPrefix(s, "job") { 
        iType = "job"
    } else if strings.HasPrefix(s, "poc") {
        iType = "poc"
    } else {
        iType = ""
    }
    return 
}

// Job holds attributes of potential jobs.
type Job struct {
    Id              int
    Title           string
    Active          bool
    Notes           string
    Company         string
    URL             string
    PocId           int
    FirstContact    int
    LastContact     int
    String          string
}

// JBuilder takes a []string and assigns data to attributes.
// Some fields are munged; string to int, 'Y' to boolean true
func (j *Job) JBuilder(data []string) {
    j.Id, _             = strconv.Atoi(data[0])
    j.Title             = data[1]
    if strings.ToLower(data[2]) == "y" {
        j.Active    = true
    }
    j.Notes             = data[3]
    j.Company           = data[4]
    j.URL               = data[5]
    j.PocId, _          = strconv.Atoi(data[6])
    if len(data[7]) != 8 {
        data[7] = strconv.Itoa(Today())
    }
    j.FirstContact, _   = strconv.Atoi(data[7])
    if len(data[8]) != 8 {
        data[8] = strconv.Itoa(Today())
    }
    j.LastContact, _    = strconv.Atoi(data[8])
    j.String            = strings.Join(data, ";")
}


// POC holds attributes of Points Of Contact.
type POC struct {
    Id              int
    Name            string
    Notes           string
    Company         string
    Email           string
    Phone           string
    FirstContact    int
    LastContact     int
    String          string
}

// PBuilder takes a []string and assigns data to attributes.
func (p *POC) PBuilder(data []string) { 
    p.Id, _             = strconv.Atoi(data[0]) 
    p.Name              = data[1]
    p.Notes             = data[2]
    p.Company           = data[3]
    p.Email             = data[4]
    p.Phone             = data[5]
    if len(data[6]) != 8 {
        data[6] = strconv.Itoa(Today())
    }
    p.FirstContact, _   = strconv.Atoi(data[6])
    if len(data[7]) != 8 {
        data[7] = strconv.Itoa(Today())
    }
    p.LastContact, _    = strconv.Atoi(data[7])
    p.String            = strings.Join(data, ";")
}

// HighestId parses the data given splits on the seperator, and returns the 
//  highest int in the first field. 
// This assumes the first field elements are all ints.
func HighestId(data []string, sep string) (id int, err error) {
    if len(data) < 1 {
        return id, err
    }
   
    for _, line := range data {
        d, err := FieldsFromLine(line, sep) 
        if err != nil {
            return 0, err
        }
        thisId, err := strconv.Atoi(d[0])
        if err != nil {
            return 0, err
        }
        if thisId > id {
            id = thisId
        }
    }
    return id, err
}

// Today returns the current date as an int in yyyymmdd format, with padded 0. 
func Today() (today int) {
    todayString := time.Now().Format("2006-01-02")
    todayString = strings.ReplaceAll(todayString, "-", "")
    today, err := strconv.Atoi(todayString)
    if err != nil {
        today = 0
    }
    return
}

// WriteFile writes the given data to the given file.
func WriteFile(fileName string, data []string) error {
    f, err  := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err == nil {
        for _, line := range data {
            line += "\n"
            _, err := f.WriteString(line)
            if err != nil {
                return err
            }
        }
    }
    return err
}


