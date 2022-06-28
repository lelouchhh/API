package structures

import (
	"net/http"
	"time"
	"context"
)

type User struct {
	Name     	string `json:"name" binding:"required"`
	Password 	string `json:"password" binding:"required"`
	DeviceId 	string `json:"device_id" binding:"required"`
}
type Tokens struct {
	Access		string `json:"access_token" db:"c_access_token"`
	Refresh		string `json:"refresh_token" db:"c_refresh_token"`
	DeviceInfo  string `json:"devId" db:"a_device_info"`
}
type New struct {
	Title			string 		`json:"title" db:"a_title"`
	Description		string 		`json:"description" db:"a_description"`
	Picture			string 	 	`json:"picture" db:"a_picture"`
	TypeTags		[]string 	`json:"typeTags" db:"a_type_tags"`
	Tags			string 		`json:"tags" db:"a_tags"`
}
type Server struct {
	httpServer *http.Server
}
type Entity struct {
	Id				string	`json:"entity_id" db:"entity_id"`
	PurchaseNumber 	string	`json:"purchasenumber" db:"purchasenumber"`
	Price			string	`json:"price" db:"price"`
	Etp 			string	`json:"etp" db:"etp"`
	Href 			string	`json:"href" db:"href"`
	PurchaseInfo 	string	`json:"purchaseobjectinfo" db:"purchaseobjectinfo"`
	DocPublishDate 	string	`json:"docpublishdate" db:"docpublishdate"`
}
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}