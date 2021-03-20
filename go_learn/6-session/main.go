package main

import (
	"sync"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ := NewManager("memory", "goDessionId", 3600)
}

func main() {
	
}