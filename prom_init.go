package prom_init

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func init() {
	go func() {
		r := http.NewServeMux()
		r.Handle("/metrics", promhttp.Handler())

		if err := http.ListenAndServe(":7070", r); err != nil {
			log.Printf("error serving prometheus: %v", err)
		}
	}()
}
