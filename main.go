package main

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Env struct{
	db *sql.DB
}

func InitDB(dataStoreName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataStoreName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main(){
	var (
		id uint8
		name string
	)

	db, err := InitDB("root:@tcp(127.0.0.1:3306)/goto")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	rows, err := env.db.Query("SELECT * FROM blah")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer db.Close()

	for rows.Next(){
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

}




/*

1. Someone comes and seeds the DB 

tripInfo {
	"vehicle" : ,
	"from": ,
	"to": ,
	"capacity": 
	"trip": [ {
		"from_node": ,
		"to_node": ,
		"leaving_from_at": ,
		"arriving_to_at":
	}, ...

	]
}

addTrip(tripInfo){

}


vehicleSearchFilter{
	capacity
	LeavingFromAt
	ArrivingToAt
}


Get Vehicles Available (from, to , vehicleSearchFilter){
	{
		"vehicles_available": [{
			"vehicle_id": ,
			"trip_id": ,
			"vehicle_capacity": ,
			..
		},

		]


	}

}


*/