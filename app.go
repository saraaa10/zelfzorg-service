package main

import "service-api/src/main/util"

func main() {
	db := util.InitDB()
	util.AuthMigrateDB(db)
}
