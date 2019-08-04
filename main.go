package main

import (
	"log"
	"github.com/nial26/goto/db"
	"github.com/gorilla/mux"
	"github.com/nial26/goto/api"
	// "strconv"
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
	r.HandleFunc("/register_trip", loggingMiddleWare(registerTripHandler)).Methods("POST")
	r.HandleFunc("/get_transits", loggingMiddleWare(getTransitsHandler)).Methods("GET")
}

func getTransitsHandler(w http.ResponseWriter, r *http.Request) {
	var transitSearchFilter api.TransitSearchFilter
	vals := r.URL.Query()

	log.Println(vals)

	// layout := "2019-08-03 13:37:33"
	// reachingBefore, err := time.Parse(layout, vals["reaching_before"][0])
	
	// if err != nil {
	// 	log.Println(err)
	// }

	// capacity, err := strconv.Atoi(vals["capacity"][0])

	// if err != nil {
	// 	log.Println(err)
	// }


	transitSearchFilter = api.TransitSearchFilter{From: vals["from"][0], To: vals["to"][0]}
	transit, err := api.GetTransits(dbEnv, transitSearchFilter)
	if err != nil {
		log.Panic(err)
	}
	json.NewEncoder(w).Encode(transit)
	

}

func registerTripHandler(w http.ResponseWriter, r *http.Request){
	var tripDetail api.TripDetail
	json.NewDecoder(r.Body).Decode(&tripDetail)
	err := api.CreateTrip(dbEnv, tripDetail)
	if err != nil {
		log.Panic(err)
	}

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