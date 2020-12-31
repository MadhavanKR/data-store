package httpServer

import (
	"encoding/json"
	"net/http"

	"github.com/milkmocha-datastore/pkg/redisAdapter"
)

func NewHttpServer() *http.Server {
	httpServer := &http.Server{
		Addr:    ":2527",
		Handler: getHandler(),
	}
	return httpServer
}

func getHandler() http.Handler {
	serverMux := http.DefaultServeMux
	serverMux.HandleFunc("/", helloWorld)
	return serverMux
}

type keyValues struct {
	Key   string
	Value string
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	keyValueMap := redisAdapter.GetAllKeys()
	keys := make([]keyValues, 0)
	for key, value := range keyValueMap {
		keys = append(keys, keyValues{
			Key:   key,
			Value: value,
		})
	}
	data, _ := json.Marshal(keys)
	w.Write(data)
}
