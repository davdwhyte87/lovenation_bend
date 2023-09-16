package routes

import (
	"context"
	"lovenation_bend/configs"
	"lovenation_bend/dao"
	"lovenation_bend/models"
)

// this contains all the setups for data acess for all application routes


var factoryDAO *dao.FactoryDAO
var UserDAO *dao.UserDAO
var FactoryDAO *dao.FactoryDAO
var RoleDAO *dao.RoleDAO
var VisaApplicationDAO *dao.VisaApplicationDAO

func GetDAO () (*dao.FactoryDAO, *dao.UserDAO){
	return factoryDAO, UserDAO
}

func SetupDAO (){
	factoryDAO = dao.InitializeFactory(configs.DB, context.TODO())
	UserDAO = &dao.UserDAO{
		 Collection: configs.GetCollection(configs.DB, models.UserCollection),
		 Context: context.TODO(),
	}

	RoleDAO = &dao.RoleDAO{
		Collection: configs.GetCollection(configs.DB, models.RoleCollection),
		Context: context.TODO(),	
	}
	VisaApplicationDAO = &dao.VisaApplicationDAO{
		Collection: configs.GetCollection(configs.DB, models.VisaApplicationsCollection),
		Context: context.TODO(),	
	}
}