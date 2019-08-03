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

func CreateTrip(dbEnv *db.DBEnv, tripDetail TripDetail) (error) {
	log.Println("[api/GetTripDetail] Got Trip Detail: ", tripDetail)

	_, err := models.CreateTripInfo(dbEnv, tripDetail.Trip)
	if err != nil {
		return err
	}
	log.Println("[api/GetTripDetail] Persisted Trip Info...")

	err = models.CreateRoutes(dbEnv, tripDetail.Routes)

	if err != nil {
		return err
	}

	log.Println("[api/GetTripDetail] Persisted Routes...l")
	return nil
}