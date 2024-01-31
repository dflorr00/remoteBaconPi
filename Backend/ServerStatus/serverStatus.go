package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/correct", handlerCorrect)
	mux.HandleFunc("/warning", handlerWarning)
	mux.HandleFunc("/failure", handlerFailure)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8001", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Custom-Header", "Cookie"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	port := ":8081"
	log.Printf("Iniciando servidor en el puerto %s...\n", port)

	http.ListenAndServe(port, handler)
}

func handlerCorrect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprint(w,
		`<Serviceability>
		<Printer_Status>
			<PrintEngine>
				<Status>Idle</Status>
			</PrintEngine>
		</Printer_Status>
	</Serviceability>`)
}

func handlerWarning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprint(w,
		`<Serviceability>
		<Printer_Status>
			<PrintEngine>
				<Status>CartridgeExpired</Status>
			</PrintEngine>
		</Printer_Status>
	</Serviceability>`)
}

func handlerFailure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprint(w,
		`<Serviceability>
		<Printer_Status>
			<PrintEngine>
				<Status>Failure</Status>
			</PrintEngine>
		</Printer_Status>
	</Serviceability>`)
}
