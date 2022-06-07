package main
import (
	 "log"
	 "log/syslog"
)
func main() {
	syslog, err := syslog.New(syslog.LOG_SYSLOG, "systemlog.go")
	if err !=nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(syslog)
		log.Println("Everything is fine!")
	}
}