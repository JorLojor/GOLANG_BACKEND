package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)





var Connection *pgx.Conn

func ConnectToPostgresql() {
	connectionString := "postgres://710193:Telkom.0329@localhost:5432/newdb"

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Terjadi kesalahan saat koneksi ke database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Koneksi ke database berhasil .....")
	Connection = conn
}

func CloseConnection() {
	if Connection != nil {
		err := Connection.Close(context.Background())
		if err != nil {
			log.Printf("Terjadi kesalahan saat menutup koneksi: %v\n", err)
		}
	}
}
