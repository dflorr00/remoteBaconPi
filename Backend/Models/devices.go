package models

import (
	"database/sql"
	"log"
)

type Device struct {
	DeviceID   int
	DeviceName string
	Service    string
	Ip         string
	Port       int
	OwnerID    int
	Status     int
}

// Devuelve todos los devices del id indicado
func (d *Device) GetDevices(id int) (result []Device) {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Busqueda de dispositivos
	var rows *sql.Rows
	if id == -1 {
		rows, err = db.Query("SELECT * FROM devices")
	} else {
		rows, err = db.Query("SELECT devices.* FROM devices JOIN views ON devices.DeviceId = views.DeviceId WHERE views.UserId = ?", id)
	}

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Almacenamiento de los dispositivos
	for rows.Next() {

		err = rows.Scan(&d.DeviceID, &d.DeviceName, &d.Service, &d.Ip, &d.Port, &d.OwnerID, &d.Status)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, *d)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return result
}

// Eliminar un dispositivo
func (d *Device) DeleteDevice() error {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Eliminación de la tabla views
	_, err = db.Exec("DELETE FROM views WHERE DeviceId = ?", d.DeviceID)
	if err != nil {
		log.Println("[Error Servidor] Error al eliminar de la tabla views (Función -- deleteDevice())")
	}
	log.Println("Se ha eliminado correctamente de la tabla views: ", d.DeviceID)

	// Eliminación del dispositivo
	_, err = db.Exec("DELETE FROM devices WHERE deviceID=?", d.DeviceID)
	if err != nil {
		log.Println("[Error Servidor] Error al preparar la petición (Función -- deleteDevice())")
	}
	log.Println("Se ha eliminado correctamente el dispositivo: ", d.DeviceID)

	return nil
}

// Añadir un dispositivo a la base de datos devuelve 0 si ya existe, 1 si lo añade y -1 en caso de error
func (d *Device) AddDevice() (int, error) {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Comprobamos si el elemento exite en la base de datos
	count := 0
	err = db.QueryRow("SELECT COUNT(*) FROM devices WHERE deviceName = ? AND Service = ?", d.DeviceName, d.Service).Scan(&count)
	if err != nil {
		log.Println("[Error Servidor] Error al comprobar si existia el dispositivo (Función -- newDevice())")
		return -1, err
	}

	//Si no existe lo insertamos
	if count == 0 {
		stmt, err := db.Prepare("INSERT INTO devices (deviceName,Service,ip,Port,OwnerId) VALUES (?,?,?,?,?)")
		if err != nil {
			log.Println("[Error Servidor] Error al preparar la inserción (Función -- newDevice())")
			return -1, err
		}
		_, err = stmt.Exec(d.DeviceName, d.Service, d.Ip, d.Port, d.OwnerID)
		if err != nil {
			log.Println("[Error Servidor] Error al ejecutar la inserción (Función -- newDevice())")
			return -1, err
		}

		// Obtener el último DeviceId generado
		var last_id int
		err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&last_id)
		if err != nil {
			log.Println("[Error Servidor] Error al obtener el último DeviceId (Función -- newDevice())")
			return -1, err
		}

		// Actualizamos la vista
		_, err = db.Exec("INSERT INTO views (UserId,DeviceId) VALUES (?,?)", d.OwnerID, last_id)
		if err != nil {
			log.Println("[Error Servidor] Error al insertar en views el user (Función -- newDevice())")
			return -1, err
		}
		if d.OwnerID != 1 {
			_, err = db.Exec("INSERT INTO views (UserId,DeviceId) VALUES (?,?)", 1, last_id)
		}
		if err != nil {
			log.Println("[Error Servidor] Error al insertar en views el admin (Función -- newDevice())")
			return -1, err
		}

	} else {
		return 0, nil
	}

	return 1, nil
}

func (d *Device) SetStatus() int {
	// Conexión a la base de datos
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		return -1
	}
	defer db.Close()

	// Actualización del estado del dispositivo en la base de datos
	_, err = db.Exec("UPDATE devices SET status = ? WHERE Ip = ?", d.Status, d.Ip)
	if err != nil {
		log.Println("[Error Servidor] Error al actualizar el estado del dispositivo en la base de datos (Función -- setStatus())")
		return -1
	}
	return 0
}
