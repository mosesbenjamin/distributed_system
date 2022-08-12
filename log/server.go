package log

import (
	"io/ioutil"
	stlog "log" // alias standard log package
	"net/http"
	"os"
)

// custom logger to handle logging for application
var log *stlog.Logger

// handle actual logging
type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// instantiate logger
func Run(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

// register handlers
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil || len(msg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
