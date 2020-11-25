package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var port = envOrString("SRE_STUB_PORT", "8080")

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "STUB IMAGE RECEIVED REQUEST ON PORT %s", port)
}

func envOrString(key, defaultValue string) (value string) {
	value, found := os.LookupEnv(key)
	if found {
		return
	}
	value = defaultValue
	return
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/metrics", promhttp.Handler())

	errCh := make(chan error)

	go func() {
		log.WithFields(log.Fields{
			"Port": port,
		}).Info("Starting service")
		errCh <- http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}()

	go func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT)
		errCh <- fmt.Errorf("%s", <-signalCh)
	}()

	log.WithFields(log.Fields{
		"Error": <-errCh,
	}).Info("Shutdown service")
}
