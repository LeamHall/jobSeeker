// name    :	jobSeeker.go
// version :	0.0.1
// date    :	20230427
// author  :	Leam Hall
// desc    :	Tools for tracking job applications

package jobSeeker

import (
    "errors"
    //"fmt"
    "os"
    "strconv"
    "strings"
)


// DataFromFile takes a fileName string, and returns a slice of strings.
// TODO:    Remove lines with comment characters like / and #
//          Do not include empty strings.
func DataFromFile(f string) (data []string, err error) {
    handle, err := os.ReadFile(f)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return data, nil
        }
        return data, err
    }
    if len(handle) == 0 {
        return data, err
    }
    lineData := strings.Split(string(handle), "\n")
    for _, line := range lineData {
        line = strings.TrimSpace(line)
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
// TODO:
//  If PocId provided, set LastContact to the same as the POC.
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
    p.Id, _         = strconv.Atoi(data[0]) 
}
 
