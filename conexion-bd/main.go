package main

import (
	"conection/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// Reemplazar estos valores con los de la base de datos
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v\n", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConfig := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Abre la conexión a la base de datos
	db, err := sql.Open("pgx", dbConfig)
	if err != nil {
		log.Fatalf("No se pudo abrir la conexión a la base de datos: %v\n", err)
	}
	defer db.Close() // Cierra la conexión al finalizar la función main

	// Verifica que la conexión sea válida
	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo hacer ping a la base de datos: %v\n", err)
	}

	fmt.Println("Conexión a PostgreSQL exitosa!")

	query := "SELECT user_name, email, firstnames, lastnames, mobilenumber FROM mosat.cliente_mosat WHERE user_name LIKE '%prue%' ORDER BY user_name ASC LIMIT 10;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserName, &user.Email, &user.FirstName, &user.LastName, &user.Phone); err != nil {
			log.Fatalf("Error al escanear la fila: %v\n", err)
		}
		fmt.Printf("Usuario: %+v\n", user.UserName)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("Nombres: %s\n", user.FirstName)
		fmt.Printf("Apellidos: %s\n", user.LastName)
		fmt.Printf("Teléfono: %s\n", user.Phone)
		fmt.Println("------------------------------")

	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error en las filas: %v\n", err)
	}

}
