package handlers

import (
	models "backend/Models"
	tools "backend/Tools"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Pagina de inicio de sesión
func HandlerSignIn(w http.ResponseWriter, r *http.Request) {

	// Intenta logear al usuario con username y password y devuelve una cookie para la sesion
	if r.Method == "POST" {

		// Leer el cuerpo de la petición
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "[Error Servidor] Error al leer el cuerpo de la petición (Función -- handlerSignIn(POST))", http.StatusBadRequest)
			return
		}

		// Almacenar el cuerpo en un usuario
		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "[Error Servidor] Error al parsear el cuerpo de la petición (Función -- handlerSignIn(POST))", http.StatusBadRequest)
			return
		}

		// Obtener el id del usuario logeado
		var valorToken int = 0
		valorToken = user.LoginUser()
		if valorToken == -1 || valorToken == 0 {
			http.Error(w, "[Error Servidor] Error, no se encuentra el usuario (Función -- handlerSignIn(POST))", http.StatusBadRequest)
			return
		} else {

			// Crear el token de inicio de sesión
			var sesion tools.Session
			token, err := sesion.GeneraToken(valorToken)
			if err != nil {
				http.Error(w, "[Error Servidor] Error al generar el token de sesión (Función -- handlerSignIn(POST))", http.StatusBadRequest)
				return
			}

			// Generar la cookie de sesión
			expiresAt := time.Now().Add(3600 * time.Second)
			cookie := http.Cookie{
				Name:     "jwt",
				Value:    token,
				Expires:  expiresAt,
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
				HttpOnly: false,
			}
			http.SetCookie(w, &cookie)
			// Respuesta del servidor
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(cookie.String()))
		}
	} else if r.Method == "GET" {

		//Llamamos al discovery
		tools.Discovery()

		var id int
		// Leer la cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			log.Println(w, "[Error Servidor] Error al obtener la cookie (Función -- handlerDevices(GET))")
			id = -1
		} else {
			// Obtener el id de la cookie
			var s tools.Session
			id = s.DecodificaIDToken(cookie.Value)
		}

		// Obtener los dispositivos
		var device models.Device
		var devicesList []models.Device
		log.Print("Obteniendo los dispositivos del usuario: ", id)
		devicesList = device.GetDevices(id)

		if devicesList == nil {
			http.Error(w, "[Error Servidor] Error al obtener los dispositivos (Función -- handlerDevices(GET))", http.StatusInternalServerError)
			return
		}

		// Almacenar la respuesta
		data, err := json.Marshal(devicesList)
		if err != nil {
			log.Println("[Error Servidor] Fallo al crear el json (Función -- handlerDevices(GET))")
		}

		// Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		http.Error(w, "[Error Servidor] Método no permitido (Función -- handlerSignIn)", http.StatusMethodNotAllowed)
		return
	}
}
