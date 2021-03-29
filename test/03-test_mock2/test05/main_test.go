package main

import (

)

type Tgbotapi interface {
	NewObj(string, int) *Person
	NewMessage(string) string
}