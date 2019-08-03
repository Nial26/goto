package main

import (
	// "fmt"
	"log"
	"github.com/nial26/goto/db"
	// "github.com/nial26/goto/models"
	"github.com/gorilla/mux"
	"github.com/nial26/goto/api"
	// "time"
	"net/http"
	"encoding/json"
)

var dbEnv *db.DBEnv

func main(){

	dbEnv = initDb("root:@tcp(127.0.0.1:3306)/goto?parseTime=true")

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v0.1").Subrouter()
	registerApiRoutes(apiRouter)
	log.Println("Starting Server...")
	http.ListenAndServe(":80", r)

	// getTripInfo(dbEnv, "7D73A463-1CA1-497B-B193-B961CAF3FAD8")

	// creatingTripInfo := models.TripInfo{TripId: "Blah", FromPosition: "B", ToPosition: "C", Vehicle: "Mercedes"}

	// createTripInfo(dbEnv, creatingTripInfo)

	// getRoutesFromAndTo(dbEnv, "A", "B")

	// routes := []models.RouteInfo{}
	// route1 := models.RouteInfo{TripId: "Blah", FromNode: "A", ToNode: "B", LeavingFromAt:time.Now(), ArrivingToAt:time.Now(), Capacity:5}
	// routes = append(routes, route1)
	// addRoutes(dbEnv, routes)
}


func loggingMiddleWare(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		log.Println(r.Method, r.URL.Path)
		f(w, r)
	}
}

func registerApiRoutes(r *mux.Router){
	log.Println("Registering Routes to /api/v0.1 ...")
	r.HandleFunc("/get_trip/{trip_id}", loggingMiddleWare(getTripInfoHandler)).Methods("GET")
}


func getTripInfoHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tripId := vars["trip_id"]
	tripDetail, err := api.GetTripDetail(dbEnv, tripId)
	if err != nil {
		log.Panic(err)
	}
	json.NewEncoder(w).Encode(tripDetail)
}


func initDb(dataStoreName string) *db.DBEnv {
	dbEnv, err := db.InitDB(dataStoreName)
	if err != nil {
		log.Panic(err)
	}
	log.Println("DB Initialized!")
	return dbEnv
}


// func addRoutes(dbEnv *db.DBEnv, routes []models.RouteInfo) {
// 	models.CreateRoutes(dbEnv, routes)
// }

// func getRoutesFromAndTo(dbEnv *db.DBEnv, from string, to string) {
// 	routeInfos, err := models.GetRoutes(dbEnv, "A", "B")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Println(routeInfos)
// }


// func getTripInfo(dbEnv *db.DBEnv, id string){
// 	tripInfo, err := models.GetTripInfoById(dbEnv, id)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Println(tripInfo)
// }


// func createTripInfo(dbEnv *db.DBEnv, creatingTripInfo models.TripInfo){

// 	res, err := models.CreateTripInfo(dbEnv, creatingTripInfo)

// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Println(res)
// }
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