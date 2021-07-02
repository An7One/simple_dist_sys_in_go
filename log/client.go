package log

import (
	"bytes"
	"fmt"
	stdlog "log"
	"net/http"

	"github.com/an7one/tutorial/simple_dist_sys_in_go/registry"
)

func SetClientLogger(serviceUrl string, clientService registry.ServiceName) {
	stdlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stdlog.SetFlags(0)
	stdlog.SetOutput(&clientLogger{url: serviceUrl})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to send log messages. Service responsded with the status code: %d", res.StatusCode)
	}
	return len(data), nil
}
