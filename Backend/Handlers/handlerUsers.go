package handlers

import (
	models "backend/Models"
	tools "backend/Tools"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HandlerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Devuelve el usuario indicado en la cookie si es el user 1 devuelve todos
	case "GET":

		// Obtener la cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "[Error Servidor] Error al obtener la cookie (Función -- HandlerUsers(GET))", http.StatusBadRequest)
			return
		}

		// Decodificar la cookie y obtener el id
		var sesion tools.Session
		var id int = 0
		id = sesion.DecodificaIDToken(cookie.Value)

		// Obtener el usuario
		var user models.User
		user.UserID = id
		users := user.GetAllUsers()

		// Devolver el usuario
		data, err := json.Marshal(users)
		if err != nil {
			log.Println(w, "[Error Servidor] Fallo al crear el json (Función -- handlerUsers(GET))")
			return
		}

		// Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		w.Write(data)

	// Añade un usuario a la bdd
	case "POST":

		// Leer el cuerpo de la petición
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("[Error Servidor] Error al leer el cuerpo de la petición (Función -- handlerUsers(POST))")
			return
		}

		// Almacenar el cuerpo en un usuario
		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Println("[Error Servidor] Error al parsear el cuerpo de la petición (Función -- handlerUsers(POST))")
			return
		}

		//Insertar usuario en la bd
		id := user.InsertUser()

		// Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		if id == 0 {
			w.Write([]byte("Usuario registrado con éxito"))
		} else {
			w.Write([]byte("Error, ese usuario ya existe"))
		}
	}
}
