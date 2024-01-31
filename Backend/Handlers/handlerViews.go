package handlers

import (
	models "backend/Models"
	tools "backend/Tools"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HandlerViews(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		// Leer la cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			log.Println(w, "[Error Servidor] Error al obtener la cookie (Función -- handlerDevices(GET))")
			return
		}

		// Obtener el id de la cookie
		var s tools.Session
		id := s.DecodificaIDToken(cookie.Value)

		// Obtener los dispositivos
		var view models.View
		var viewsList []models.View
		view.UserID = id

		if id >= 1 {
			log.Print("Obteniendo las vistas del usuario: ", id)
			viewsList = view.GetAllViews()
		}

		if viewsList == nil {
			http.Error(w, "[Error Servidor] Error al obtener los dispositivos (Función -- handlerDevices(GET))", http.StatusInternalServerError)
			return
		}

		// Almacenar la respuesta
		data, err := json.Marshal(viewsList)
		if err != nil {
			log.Println("[Error Servidor] Fallo al crear el json (Función -- handlerDevices(GET))")
		}

		// Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		w.Write(data)

	// Añade una vista
	case "POST":

		// Leer el cuerpo de la petición
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("[Error Servidor] Error al leer el cuerpo de la petición (Función -- handlerViews(POST))")
			return
		}

		// Almacenar en una estructura views
		var view models.View
		err = json.Unmarshal(body, &view)
		if err != nil {
			log.Println(w, "[Error Servidor] Error al parsear el cuerpo de la petición (Función -- handlerViews(POST))")
			return
		}

		// Añadir view a la bdd
		var codigo int
		codigo, err = view.AddView()
		if err != nil {
			log.Println(w, "[Error Servidor] Error al añadir la vista (Función -- handlerViews(POST))")
			return
		}

		// Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		var mensage string
		if codigo == 0 {
			mensage = "Error al añadir, la vista ya existe"
		} else {
			mensage = "Vista añadida correctamente"
		}

		w.Write([]byte(mensage))
	// Eliminar vista de la bdd
	case "DELETE":

		// Leer el cuerpo de la petición
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("[Error Servidor] Error al leer el cuerpo de la petición (Función -- handlerViews(DELETE))")
			return
		}

		// Parsear el cuerpo de la petición como JSON
		var view models.View
		err = json.Unmarshal(body, &view)
		if err != nil {
			log.Println(w, "[Error Servidor] Error al parsear el cuerpo de la petición (Función -- handlerViews(DELETE))")
			return
		}

		// Realizar la lógica para eliminar la vista
		err = view.DeleteView()
		if err != nil {
			log.Println(w, "[Error Servidor] Error al eliminar la vista (Función -- handlerViews(DELETE))")
			return
		}

		// Enviar una respuesta
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Vista eliminada"))
	}
}
