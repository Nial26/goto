package api

import (
	"log"
	"time"
	"github.com/nial26/goto/models"
	"github.com/nial26/goto/db"
)

type TransitSearchFilter struct {
	From string `json:"from"`
	To string `json:"to"`
	Capacity int `json:"capacity"`
	ReachingBefore time.Time `json:"reaching_before"`
}

type Transit struct {
	Routes [][]models.RouteInfo `json:"routes"`
}

func GetTransits(dbEnv *db.DBEnv, transitSearchFilter TransitSearchFilter) (Transit, error) {
	log.Println("[api/GetTransits] Got Transit Search Filter : ", transitSearchFilter)

	var transit Transit

	possiblePaths := [][]models.RouteInfo{}
	err := GetRoutesBetween(dbEnv, transitSearchFilter.From, transitSearchFilter.To, []models.RouteInfo{}, &possiblePaths)
	if err != nil {
		return transit, err
	} 
	transit = Transit{possiblePaths}
	log.Println("[api/GetTransits] Transit Details: ", transit)
	return transit, nil
}

func GetRoutesBetween(dbEnv *db.DBEnv, from string, to string, seenRoutes []models.RouteInfo, possiblePaths *[][]models.RouteInfo) (error) {
	log.Printf("[api/GetRoutesBetween] Searching for Routes between : %s and %s with seenRoutes: %v", from, to, seenRoutes)
	if from == to {
		log.Println("[api/GetRoutesBetween] Got From and To as Same Returning seenRoutes : ", seenRoutes)
		*possiblePaths = append(*possiblePaths, seenRoutes)
		return nil
	}

	routes, err := models.GetRoutesFrom(dbEnv, from)

	if err != nil {
		return err
	}

	log.Printf("[api/GetRoutesBetween] Nodes originating from %s : %v", from, routes)


	for _, route := range routes{
		if !IsPresentIn(route, seenRoutes) {
			seenRoutes = append(seenRoutes, route)
			GetRoutesBetween(dbEnv, route.ToNode, to, seenRoutes, possiblePaths)
		}
	}
	return nil
}

func IsPresentIn(route models.RouteInfo, routes []models.RouteInfo) bool {
	for _, r := range routes {
		if r == route {
			return true
		}
	}
	return false
}