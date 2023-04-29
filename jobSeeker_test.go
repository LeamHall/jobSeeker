// name    :	jobSeeker_test.go
// version :	0.0.1
// date    :	20230427
// author  :	Leam Hall
// desc    :    Test jobSeeker

package jobSeeker_test

import (
    "os" 
    "strings"
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
    if j.String != strings.Join(data, ";") {
        t.Errorf("Expected a string of %s, got %s", strings.Join(data, ";"), j.String)
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
    if p.String != strings.Join(data, ";") {
        t.Errorf("Expected a string of %s, got %s", strings.Join(data, ";"), p.String)
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

func TestInputType(t *testing.T) {
    typeJ       := jobSeeker.InputType("job;this;that;the other thing")
    typeP       := jobSeeker.InputType("poc;something;allthing; nothing")
    typeN       := jobSeeker.InputType("fred;joe;sam")
    if typeJ != "job" {
        t.Errorf("Expected a %s, got %s", "job", typeJ)
    }
    if typeP != "poc" {
        t.Errorf("Expected a %s, got %s", "poc", typeP)
    }
    if typeN != "" {
        t.Errorf("Expected an empty string, got %s", typeN)
    }
}

func TestAdd(t *testing.T) {
    baseData, err   := jobSeeker.Add("job;some;where;over",";", []string{})
    if err != nil {
        t.Errorf("Expected Add to not error out: %s", err)
    }
    if len(baseData) != 1 {
        t.Errorf("Expected baseData to have %d elements, but it has %d", 1, len(baseData))
    }
    if !strings.HasPrefix(baseData[0], "1;") {
        t.Errorf("Expected the first line to start with '1;', got %s", baseData[0])
    }
}

func TestToday(t *testing.T) {
    today   := jobSeeker.Today()
    if today < 19000000 {
        t.Errorf("Expected a date more recent than %d", today)
    }
    if today > 21000000 {
        t.Errorf("Not ready for stardates yet: %d", today)
    }
}

func TestWriteFile(t *testing.T) {
    tempFile, err := os.CreateTemp("", "datafile_writeable")
    if err != nil {
        t.Errorf("Could not create tempFile: %s", err)
    }	
    data := []string{ "one", "two", "three"}
    err = jobSeeker.WriteFile(tempFile.Name(), data)
    if err != nil {
        t.Errorf("Could not write to %s: %s", tempFile.Name(), err)
    }
    newData, err := jobSeeker.DataFromFile(tempFile.Name())
    if err != nil {
        t.Errorf("Could not read %s: %s", tempFile.Name(), err)
    }
    if len(newData) != 3 {
        t.Errorf("Expected %d elements in newData, got %d", 3, len(newData))
    }
}

