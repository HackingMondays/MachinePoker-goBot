package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	// test setup
	testMessage := "Special Information"
	out := new(bytes.Buffer)
	InitLogger(ioutil.Discard, out, ioutil.Discard, ioutil.Discard)

	// log and test
	Info.Println(testMessage)
	AssertContains(t, out, "INFO: "+testMessage)
}

// AssertContains returns true if a buffer contains a line containing message.
func AssertContains(t *testing.T, out *bytes.Buffer, message string) {
	if !strings.Contains(out.String(), message) {
		t.Errorf("Logs do not contain %q, content: %q", message, out.String())
	}
}
