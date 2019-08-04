package models

import (
	"github.com/nial26/goto/db"
	"time"
)

type RouteInfo struct {
	Id uint8 `json:"id"`
	TripId string `json:"trip_id"`
	FromNode string `json:"from_node"`
	ToNode string `json:"to_node"`
	LeavingFromAt time.Time `json:"leaving_from_at"`
	ArrivingToAt time.Time `json:"arriving_to_at"`
	Capacity int `json:"capacity"`
}



func GetRoutes(dbEnv *db.DBEnv, fromNode string, toNode string)([]RouteInfo, error) {

	routes := []RouteInfo{}

	queryString := "SELECT * FROM route_info WHERE from_node = ? AND to_node = ?"
	rows, err := dbEnv.Db.Query(queryString, fromNode, toNode)
	defer rows.Close()

	for rows.Next() {
		var route RouteInfo
		err = rows.Scan(&route.Id, &route.TripId, &route.FromNode, &route.ToNode, &route.LeavingFromAt, &route.ArrivingToAt, &route.Capacity)
		if err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err != nil {
		return nil, err
	}
	return routes, nil
}


func GetRoutesFrom(dbEnv *db.DBEnv, fromNode string)([]RouteInfo, error){
	routes := []RouteInfo{}

	queryString := "SELECT * FROM route_info WHERE from_node = ?"
	rows, err := dbEnv.Db.Query(queryString, fromNode)
	defer rows.Close()

	for rows.Next() {
		var route RouteInfo
		err = rows.Scan(&route.Id, &route.TripId, &route.FromNode, &route.ToNode, &route.LeavingFromAt, &route.ArrivingToAt, &route.Capacity)
		if err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err != nil {
		return nil, err
	}
	return routes, nil
}


func CreateRoutes(dbEnv *db.DBEnv, routes []RouteInfo) error {
	stmt, err := dbEnv.Db.Prepare("INSERT INTO route_info (trip_id, from_node, to_node, leaving_from_at, arriving_to_at, capacity) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	for _, route := range routes {
		_, err := stmt.Exec(route.TripId, route.FromNode, route.ToNode, route.LeavingFromAt, route.ArrivingToAt, route.Capacity)
		if err != nil {
			return err
		}
	}
	return nil

}

func GetRoutesForTrip(dbEnv *db.DBEnv, tripId string) ([]RouteInfo, error) {
	routes := []RouteInfo{}

	queryString := "SELECT * FROM route_info WHERE trip_id = ?"
	rows, err := dbEnv.Db.Query(queryString, tripId)
	defer rows.Close()

	for rows.Next() {
		var route RouteInfo
		err = rows.Scan(&route.Id, &route.TripId, &route.FromNode, &route.ToNode, &route.LeavingFromAt, &route.ArrivingToAt, &route.Capacity)
		if err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err != nil {
		return nil, err
	}
	return routes, nil
}