package api

import (
	"github.com/nial26/goto/db"
	"github.com/nial26/goto/models"
	"log"
)

type TripDetail struct {
	Trip models.TripInfo `json:"trip"`
	Routes []models.RouteInfo `json"routes"`
}

func GetTripDetail(dbEnv *db.DBEnv, tripId string) (TripDetail, error) {
	var tripDetail TripDetail
	log.Println("[api/GetTripDetail] Got Trip Id: ", tripId)
	tripInfo, err := models.GetTripInfoById(dbEnv, tripId)
	if err != nil {
		return tripDetail, err
	}
	log.Println("[api/GetTripDetail] Got TripInfo : ", tripInfo)

	routesInfo, err := models.GetRoutesForTrip(dbEnv, tripId)
	if err != nil {
		return tripDetail, err
	}
	log.Println("[api/GetTripDetail] Got RoutesInfo : ", routesInfo)
	tripDetail = TripDetail{*tripInfo, routesInfo}
	log.Println("[api/GetTripDetail] Got TripDetail : ", tripDetail)
	return tripDetail, nil
}