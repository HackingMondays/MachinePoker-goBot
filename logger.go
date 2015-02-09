package main

import (
	// "io/ioutil"
	"log"
	"os"
)

// Custom logger that can easily be disabled, see:
// http://www.goinggo.net/2013/11/using-log-package-in-go.html

// var logger = log.New(ioutil.Discard,
var logger = log.New(os.Stdout,
	"DEBUG: ", 0)
