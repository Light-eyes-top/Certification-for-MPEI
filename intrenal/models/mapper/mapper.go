package mapper

import (
	db_models "certification/intrenal/models/db-models"
	service_models "certification/intrenal/models/service-models"
	transport_models "certification/intrenal/models/transport-models"
)

func UserServiceToDb(in *service_models.User) *db_models.User {
	return &db_models.User{
		Id:       in.Id,
		Username: in.Username,
		Password: in.Password,
	}
}

func UserDbToService(in *db_models.User) *service_models.User {
	return &service_models.User{
		Id:       in.Id,
		Username: in.Username,
		Password: in.Password,
	}
}

func UserTransportToService(in *transport_models.User) *service_models.User {
	return &service_models.User{
		Id:       in.Id,
		Username: in.Username,
		Password: in.Password,
	}
}
