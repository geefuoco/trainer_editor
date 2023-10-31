package logging

import "log"

var isLoggingEnabled bool

func IsLoggingEnabled() bool {
    return isLoggingEnabled
}

func EnableLogging() {
    isLoggingEnabled = true
}

func DisableLogging() {
    isLoggingEnabled = false
}

func InfoLog(message string, args...interface{}) {
    if isLoggingEnabled {
        log.Printf("[INFO]: "+message, args...)
    }
}

func WarnLog(message string, args...interface{}) {
    if isLoggingEnabled {
        log.Printf("[WARN]: "+message, args...)
    }
}

func ErrorLog(message string, args...interface{}) {
    log.Printf("[ERROR]: "+message, args...)
}
