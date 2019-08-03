package models

import (
	"github.com/nial26/goto/db"
	"database/sql"
)

type TripInfo struct {
	Id uint8 `json:"id" schema:"id"`
	TripId string `json:"trip_id" schema:"trip_id"`
	FromPosition string `json:"from_position" schema:"from_position"`
	ToPosition string `json:"to_position" schema:"to_position"`
	Vehicle string `json:"vehicle" schema:"vehicle"`
	MaxCapacity int `json:"max_capacity" schema:"max_capacity"`
}



func GetTripInfoById(dbEnv *db.DBEnv, tripId string)(*TripInfo, error) {
	tripInfo := &TripInfo{}
	queryString := "SELECT * FROM trip_info WHERE trip_id = ?"
	err := dbEnv.Db.QueryRow(queryString, tripId).Scan(&tripInfo.Id, &tripInfo.TripId, &tripInfo.FromPosition, &tripInfo.ToPosition, &tripInfo.Vehicle, &tripInfo.MaxCapacity)
	if err != nil {
		return nil, err
	}
	return tripInfo, nil
}

func CreateTripInfo(dbEnv *db.DBEnv, t TripInfo) (sql.Result, error) {
	insertStatement := "INSERT INTO trip_info(trip_id, from_position, to_position, vehicle, max_capacity) VALUES (?, ?, ?, ?, ?)"
	res, err := dbEnv.Db.Exec(insertStatement, t.TripId, t.FromPosition, t.ToPosition, t.Vehicle, t.MaxCapacity)
	if err != nil {
		return nil, err
	}
	return res, nil
}