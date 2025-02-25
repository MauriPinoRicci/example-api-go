package server

import (
	"github.com/MauriPinoRicci/example-api-go/users/application/create_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/delete_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/get_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/update_srv"
	"github.com/MauriPinoRicci/example-api-go/users/infra/users_dynamo"
)

type Dependencies struct {
	createSrv *create_srv.Service
	getSrv    *get_srv.Service
	updateSrv *update_srv.Service
	deleteSrv *delete_srv.Service
}

func InitDependencies() *Dependencies {

	userDynamo := users_dynamo.New()

	// Init services
	createSrv := create_srv.NewService(userDynamo)
	getSrv := get_srv.NewService(userDynamo)
	updateSrv := update_srv.NewService(userDynamo)
	deleteSrv := delete_srv.NewService(userDynamo)

	return &Dependencies{
		createSrv: createSrv,
		getSrv:    getSrv,
		updateSrv: updateSrv,
		deleteSrv: deleteSrv,
	}

}
