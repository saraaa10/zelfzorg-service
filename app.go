package main

import "service-api/src/main/infra"

func main() {
	db := infra.InitDB()
	infra.AuthMigrateDB(db)
}
