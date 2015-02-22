package main


import (
    "testing"
)

func TestLogging(t *testing.T) {
    Trace.Println("I have something standard to say")
    Info.Println("Special Information")
    Warning.Println("There is something you need to know about")
    Error.Println("Something has failed")
}