package middleware

import (
	"fmt"
	"time"
)

var formatTemplate = "2006-09-27 16:00:00"

func ServiceLog(content string) {
	t := time.Now()
	fmt.Println("[Service] [" + t.Format(formatTemplate) + "] " + content)
}

func LibraryLog(content string) {
	t := time.Now()
	fmt.Println("[Library] [" + t.Format(formatTemplate) + "] " + content)
}

func DBLog(content string) {
	t := time.Now()
	fmt.Println("[DataBase] [" + t.Format(formatTemplate) + "] " + content)
}
