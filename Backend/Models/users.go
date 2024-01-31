package models

import (
	"database/sql"
	"log"
)

type User struct {
	UserID   int
	UserName string
	Password string
	Email    string
	Type     int
}

// Obtiene el usuario indicado por su ID si es 1 devuelve todos
func (u *User) GetAllUsers() (result []User) {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Buscar al usuario en la bdd
	var rows *sql.Rows
	if u.UserID == 1 {
		rows, err = db.Query("SELECT * FROM users")
	} else if u.UserID > 1 {
		rows, err = db.Query("SELECT * FROM users WHERE UserId = ?", u.UserID)
	}
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Almacenamiento de los usuarios
	for rows.Next() {

		err = rows.Scan(&u.UserID, &u.UserName, &u.Password, &u.Email, &u.Type)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, *u)

	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return result
}

// Autentica al usuario y devuelve su id como respuesta si devuelve -1 es que no existe
func (u *User) LoginUser() int {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Comprobar si la autenticación existe en la base de datos
	var index int = 0

	err = db.QueryRow("SELECT UserId FROM users WHERE UserName = ? AND Password = ?", u.UserName, u.Password).Scan(&index)
	if err != nil {
		log.Println("[Error Servidor] Datos de inicio de sesión incorrectos:", err)
		return -1
	}

	return index
}

// Añade un usuario a la base de datos y devuelve 0 si todo va bien
func (u *User) InsertUser() int {

	// Conexión a la bdd
	db, err := sql.Open("mysql", "root:admin@/tfg")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Comprobar si el usuario ya existe en la base de datos
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE userName = ?", u.UserName).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	// Insertar el usuario si no existe
	if count == 0 {
		_, err = db.Exec("INSERT INTO users (UserName, Email, Password, Type) VALUES (?, ?, ?, ?)", u.UserName, u.Email, u.Password, 2)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Usuario añadido con éxito:", u.UserName, " ", u.Password, " ", u.Email)
	}

	return count
}
