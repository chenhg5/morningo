package alipay

import (
	//"net/http"
	"../../config"
)

type aliSdkType struct {
	ticket authKeyType
}

type authKeyType struct {
	APP_PRIVATE string
	APP_PUBLIC  string
}

var aliSdk aliSdkType

func init() {
	aliSdk.ticket = authKeyType{
		APP_PRIVATE: config.GetEnv().DATABASE_IP,
		APP_PUBLIC:  config.GetEnv().DATABASE_IP,
	}
}

func (aliSdk aliSdkType) withdraw(amount int) {

}

func (aliSdk aliSdkType) pay(amount int) {

}
