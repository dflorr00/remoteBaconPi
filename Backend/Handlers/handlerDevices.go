package handlers

import (
	models "backend/Models"
	tools "backend/Tools"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HandlerDevices(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	// Devuelve los dispositivos del usuario indicado en la cookie de la bdd, si el id = -1 devuelve todos
	case "GET":

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

	//Añade un dipositivo a la base de datos
	case "POST":

		// Leer la cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			log.Println(w, "[Error Servidor] Error al obtener la cookie (Función -- handlerDevices(POST))")
			return
		}

		// Obtener el id de la cookie
		var s tools.Session
		id := s.DecodificaIDToken(cookie.Value)

		// Leer el cuerpo de la petición
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("[Error Servidor] Error al leer el cuerpo de la petición (Función -- handlerDevices(POST))")
			return
		}

		// Almacenar en una estructura device
		var device models.Device
		err = json.Unmarshal(body, &device)
		if err != nil {
			log.Println(w, "[Error Servidor] Error al parsear el cuerpo de la petición (Función -- handlerDevices(POST))")
			return
		}
		device.OwnerID = id

		// Añadir device a la bdd
		var codigo int
		codigo, err = device.AddDevice()
		if err != nil {
			log.Println(w, "[Error Servidor] Error al añadir dispositivo (Función -- handlerDevices(POST))")
			return
		}

		// Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		var mensage string
		if codigo == 0 {
			mensage = "Error al añadir, el dispositivo ya existe"
		} else {
			mensage = "Dispositivo añadido correctamente"
		}

		w.Write([]byte(mensage))

	// Eliminar el device indicado en la url de la bdd
	case "DELETE":

		// Obtener el ID del device en la petición
		var device models.Device
		idParam := r.URL.Path[len("/devices/"):]
		id, _ := strconv.ParseInt(idParam, 10, 64)
		device.DeviceID = int(id)

		// Eliminar el device
		err := device.DeleteDevice()
		if err != nil {
			log.Println(w, "[Error Servidor] Error al eliminar el dispositivo: ", id, " (Función -- handlerDevices(DELETE))")
			return
		}

		//Respuesta del servidor
		w.WriteHeader(http.StatusOK)
		return
	}
}
