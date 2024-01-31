package models

import (
	"database/sql"
	"log"
)

type View struct {
	UserID   int
	DeviceID int
}

// Obtiene las vistas indicadas por su ID si es 1 devuelve todas
func (v *View) GetAllViews() (result []View) {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Buscar la vista en la bdd
	var rows *sql.Rows
	if v.UserID == 1 {
		rows, err = db.Query("SELECT * FROM views")
	} else if v.UserID > 1 {
		rows, err = db.Query("SELECT * FROM views WHERE UserId = ?", v.UserID)
	}
	if err != nil {
		log.Println("Erro en la consulta de vistas")
	}
	defer rows.Close()

	// Almacenamiento de las vistas
	for rows.Next() {
		err = rows.Scan(&v.UserID, &v.DeviceID)
		if err != nil {
			log.Println("Error al almacenar las vistas")
		}
		result = append(result, *v)

	}

	return result
}

// Eliminar una vista
func (v *View) DeleteView() error {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Eliminación de la tabla views
	_, err = db.Exec("DELETE FROM views WHERE UserId = ? AND DeviceId = ?", v.UserID, v.DeviceID)
	if err != nil {
		log.Println("[Error Servidor] Error al eliminar de la tabla views (Función -- DeleteView())")
	}
	log.Println("Se ha eliminado correctamente de la tabla views: ", v.DeviceID)

	return nil
}

// Añadir una vista a la base de datos devuelve 0 si ya existe, 1 si lo añade y -1 en caso de error
func (v *View) AddView() (int, error) {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Comprobamos si el elemento exite en la base de datos
	count := 0
	err = db.QueryRow("SELECT COUNT(*) FROM views WHERE UserId = ? AND DeviceId = ?", v.UserID, v.DeviceID).Scan(&count)
	if err != nil {
		log.Println("[Error Servidor] Error al comprobar si existia la vista (Función -- AddView())")
		return -1, err
	}

	//Si no existe se inserta
	if count == 0 {
		_, err := db.Exec("INSERT INTO views (UserId,DeviceId) VALUES (?,?)", v.UserID, v.DeviceID)
		if err != nil {
			log.Println("[Error Servidor] Error al insertar (Función --AddView())")
			return -1, err
		}
	} else {
		return 0, nil
	}

	return 1, nil
}
