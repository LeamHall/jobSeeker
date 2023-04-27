// name    :	jobSeeker_test.go
// version :	0.0.1
// date    :	20230427
// author  :	Leam Hall
// desc    :    Test jobSeeker

package jobSeeker_test

import (
    "os" 
    "testing"
    
    "github.com/LeamHall/jobSeeker"
)

func TestDataFromFile(t *testing.T) {
    tempFile, err := os.CreateTemp("", "datafile")
    if err != nil {
        t.Errorf("Could not create tempFile: %s", err)
    }	
    defer os.Remove(tempFile.Name())

    if _, err := tempFile.Write([]byte("  first\n second \n third \n the rest  ")); err != nil {
        t.Errorf("Could not write to tempFile: %s", err)
    }

    data, err := jobSeeker.DataFromFile(tempFile.Name())
    if err != nil {
        t.Errorf("Could not read tempFile: %s", err)
    }
    if len(data) != 4 {
        t.Errorf("Expected %d items in data, got %d", 4, len(data))
    }

    if err := tempFile.Close(); err != nil {
        t.Errorf("Could not close tempFile: %s", err)
    }
}

func TestFieldsFromLine(t *testing.T) {
    tempLine := "one;   two   ; drei; kvar"
    data, err := jobSeeker.FieldsFromLine(tempLine, ";")
    if err != nil {
        t.Errorf("Could not parse the line: %s", err)
    }
    if len(data) != 4 {
        t.Errorf("Expected %d phrases, got %d", 4, len(data))
    }
    if len(data[1]) != 3 {
        t.Errorf("Expected a trimmed to %d datum, got %d sized", 3, len(data[1]))
    }
}

func TestJBuilder(t *testing.T) {
    data := []string{
        "23", 
        "Sr Automation Eng", 
        "Y",
        "Good person", 
        "MyCo", 
        "myco.com", 
        "1", 
        "20230101", 
        "20230401",
        }
    j       := jobSeeker.Job{}
    j.JBuilder(data)
    if j.Id != 23 {
        t.Errorf("Expected id of %d, got %d", 23, j.Id)
    }
    if !j.Active {
        t.Errorf("Expected j.Active to be true")
    }
    if j.Notes != "Good person" {
        t.Errorf("Expected %s, got %s", "Good person", j.Notes)
    }
    if j.Company != "MyCo" {
        t.Errorf("Expected %s, got %s", "MyCo", j.Company)
    }
    if j.PocId != 1 {
        t.Errorf("Expected PocId of %d, got %d", 1, j.PocId)
    }
    if j.FirstContact != 20230101 {
        t.Errorf("Expected FirstContact of %d, got %d", 20230101, j.FirstContact)
    }
    if j.LastContact != 20230401 {
        t.Errorf("Expected LastContact of %d, got %d", 20230401, j.LastContact
    }
}

