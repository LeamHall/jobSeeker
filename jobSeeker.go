// name    :	jobSeeker.go
// version :	0.0.1
// date    :	20230427
// author  :	Leam Hall
// desc    :	Tools for tracking job applications

package jobSeeker

import (
    //"errors"
    "os"
    "strconv"
    "strings"
)


// DataFromFile takes a fileName string, and returns a slice of strings.
func DataFromFile(f string) (data []string, err error) {
    comments := []string{"#", "/", "*",}
    handle, err := os.ReadFile(f)
    if err != nil {
        //if errors.Is(err, os.ErrNotExist) {
        //    return data, err
        //}
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

// InitFile creates the initial file, and returns an error if it fails.
//func InitFile(fileName string) (err error) {
//    os.
//}

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
    j.FirstContact, _   = strconv.Atoi(data[7])
    j.LastContact, _    = strconv.Atoi(data[8])
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
}

// PBuilder takes a []string and assigns data to attributes.
func (p *POC) PBuilder(data []string) { 
    p.Id, _             = strconv.Atoi(data[0]) 
    p.Name              = data[1]
    p.Notes             = data[2]
    p.Company           = data[3]
    p.Email             = data[4]
    p.Phone             = data[5]
    p.FirstContact, _   = strconv.Atoi(data[6])
    p.LastContact, _    = strconv.Atoi(data[7])
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
 
