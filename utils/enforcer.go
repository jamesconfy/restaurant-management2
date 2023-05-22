package utils

import (
	"log"

	"github.com/casbin/casbin/v2"
)

var Enforcer *casbin.Enforcer

func init() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Println("Cashbin: ", err)
	}

	Enforcer = e
}
