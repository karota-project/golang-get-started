package main

import(
    "log"
    "log/syslog"
)

func main() {

    // Configure logger to write to the syslog. You could do this in init(), too.
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "karota")
    if e == nil {
        log.SetOutput(logwriter)
    }

    // Now from anywhere else in your program, you can use this:
    log.Print("Hello Logs!")

    // env ) OS X 10.9.4 
    // $ cat /var/log/system.log | grep karota
}
