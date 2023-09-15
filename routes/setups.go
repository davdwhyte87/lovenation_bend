package routes

import (
	"context"
	"lovenation_bend/configs"
	"lovenation_bend/dao"
)

// this contains all the setups for data acess for all application routes


var factoryDAO *dao.FactoryDAO

func GetDAO () (*dao.FactoryDAO){
	return factoryDAO
}

func SetupDAO (){
	factoryDAO = dao.InitializeFactory(configs.DB, context.TODO())
}