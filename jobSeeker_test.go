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

    if _, err := tempFile.Write(
        []byte("\n \n \n # Comment #1 \n // Comment \n */ comment\n first\n second \n third \n the rest  ")); err != nil {
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
        "Good reviews", 
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
    if j.Notes != "Good reviews" {
        t.Errorf("Expected %s, got %s", "Good reviews", j.Notes)
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
        t.Errorf("Expected LastContact of %d, got %d", 20230401, j.LastContact)
    }
}

func TestPBuilder(t *testing.T) {
    data := []string{
        "32",
        "John Smythe",
        "Long time recruiter", 
        "MyCo",
        "jsmythe@myco.com",
        "555.1212",
        "20230101",
        "20230401",
        }
    p   := jobSeeker.POC{}
    p.PBuilder(data)
    if p.Id != 32 {
        t.Errorf("Expected an Id of %d, got %d", 32, p.Id)
    }
    if p.Name != "John Smythe" {
        t.Errorf("Expected Name of %s, got %s", "John Smythe", p.Name)
    }
    if p.Notes != "Long time recruiter" {
        t.Errorf("Expected notes of %s, got %s", "Long time recruiter", p.Notes)
    }
    if p.Company != "MyCo" {
        t.Errorf("Expected company to be %s, got %s", "MyCo", p.Company)
    }
    if p.Email != "jsmythe@myco.com" {
        t.Errorf("Expected email to be %s, got %s", "jsmythe@myco.com", p.Email)
    }
    if p.Phone != "555.1212" {
        t.Errorf("Expected phone to be %s, got %s", "555.1212", p.Phone)
    }
    if p.FirstContact != 20230101 {
        t.Errorf("Expected FirstContact to be %d, got %d", 20230101, p.FirstContact)
    }
    if p.LastContact != 20230401 {
        t.Errorf("Expected LastContact to be %d, got %d", 20230401, p.LastContact)
    }
}

func TestHighestId(t *testing.T) {
    data        := []string{"1;one", "2;two", "3;drei", "4;kvar", "27315; lots",}
    sep         := ";"
    expected    := 27315
    result, err := jobSeeker.HighestId(data, sep)
    if err != nil {
        t.Errorf("Could not run HighestId: %s", err)
    }
    if result != expected {
        t.Errorf("Expected HighestId of %d, got %d", expected, result)
    }
}


