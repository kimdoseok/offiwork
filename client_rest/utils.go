package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func getRandStr(charset string, charlen int) string {
	chars := []rune(charset)
	rand.Seed(time.Now().UnixNano())
	//length := rand.Intn(500) + 500
	var b strings.Builder
	for i := 0; i < charlen; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func getConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("offiworks")
	v.SetConfigType("json")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func connectSQL() (*gorm.DB, error) {
	cfg, err := getConfig()
	dbtype = cfg.GetString("DATABASES_TYPE")
	//fmt.Println(cfg)
	if err != nil {
		return nil, err
	}

	if dbtype == "mssql" {
		connstr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			cfg.GetString("DATABASES_MSSQL_HOST"),
			cfg.GetString("databases_mssql_user"),
			cfg.GetString("DATABASES_MSSQL_PASSWORD"),
			cfg.GetInt("DATABASES_MSSQL_PORT"),
			cfg.GetString("DATABASES_MSSQL_DBNAME"))
		return gorm.Open(dbtype, connstr)
	} else if dbtype == "pgsql" {
		connstr := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
			cfg.GetString("DATABASES_PGSQL_HOST"),
			cfg.GetString("DATABASES_PGSQL_USER"),
			cfg.GetString("DATABASES_PGSQL_PASSWORD"),
			cfg.GetInt("DATABASES_PGSQL_PORT"),
			cfg.GetString("DATABASES_PGSQL_DBNAME"))
		return gorm.Open("postgres", connstr)
	} else {
		dbname := "offiworks.db"
		return gorm.Open("sqlite", dbname)
	}
}

func GetPermRBAC(e *casbin.Enforcer, req []string) bool {
	fmt.Println(*e, req[0], req[1], req[2])
	res := e.Enforce(req[0], req[1], req[2])
	if res {
		return true
	}
	users := e.GetImplicitRolesForUser(req[0])
	for _, user := range users {
		roles := e.GetImplicitRolesForUser(user)
		if len(roles) == 0 {
			return e.Enforce(user, req[1], req[2])
		} else {
			return GetPermRBAC(e, []string{user, req[1], req[2]})
		}
	}
	return false
}

func authrizing(userid, path, action string) bool {
	cfg, err := getConfig()
	if err != nil {
		return false
	}
	//fmt.Println("Configs: ", cfg.GetString("CASBIN_MODEL"), cfg.GetString("CASBIN_POLICY"))
	e := casbin.NewEnforcer(cfg.GetString("CASBIN_MODEL"), cfg.GetString("CASBIN_POLICY"))
	return GetPermRBAC(e, []string{userid, path, action})
}
