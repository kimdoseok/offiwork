package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/baidubce/bce-sdk-go/util/log"
	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"gopkg.in/ldap.v2"
)

type (
	Keycode struct {
		gorm.Model
		Key_code       string `gorm: "column:'key_code';type:'varchar(32)';unique_index" json:"Key_code"`
		Key_title      string `gorm: "column:'key_title'" json:"Key_title"`
		Key_status     int    `gorm: "column:'key_status'" json:"Key_status"`
		Key_definition string `gorm: "column:'key_definition'" json:"Key_definition"`
		Key_app        string `gorm: "column:'key_app'" json:"Key_app"`
		Key_cc         string `gorm: "column:'key_cc'" json:"Key_cc"`
		Key_hcc        string `gorm: "column:'key_hcc'" json:"Key_hcc"`
		Key_mcc        string `gorm: "column:'key_mcc'" json:"Key_mcc"`
	}

	KeycodeResp struct {
		Errno  int       `json: "errno"`
		Errmsg string    `json: "errmsg"`
		Data   []Keycode `json: "data"`
	}

	Keycoderef struct {
		gorm.Model
		Kcr_key_code   string `gorm: "column:'kcr_key_code';type:'varchar(32)';unique_index" json:"Kcr_key_code"`
		Kcr_app        string `gorm: "column:'kcr_app';type:'text'" json:"Kcr_app"`
		Kcr_title      string `gorm: "column:'kcr_title';type:'text'" json:"Kcr_title"`
		Kcr_definition string `gorm: "column:'kcr_definition';type:'text'" json:"Kcr_definition"`
		Kcr_status     int    `gorm: "column:'kcr_status';type:'int'" json:"Kcr_status"`
	}

	KeycoderefResp struct {
		Errno  int          `json: "errno"`
		Errmsg string       `json: "errmsg"`
		Data   []Keycoderef `json: "data"`
	}

	DrgCode struct {
		gorm.Model
		Drg_code       string `gorm: "column:'drg_code';type:'varchar(32)';unique_index" json:"Drg_code"`
		Drg_title      string `gorm: "column:'drg_title'" json:"Drg_title"`
		Drg_definition string `gorm: "column:'drg_definition'" json:"Drg_definition"`
	}

	DrgCodeResp struct {
		Errno  int       `json: "errno"`
		Errmsg string    `json: "errmsg"`
		Data   []DrgCode `json: "data"`
	}

	IcdCode struct {
		gorm.Model
		Icd_code       string `gorm: "column:'icd_code';type:'varchar(32)';unique_index" json:"Icd_code"`
		Icd_title      string `gorm: "column:'icd_title'" json:"Icd_title"`
		Icd_definition string `gorm: "column:'icd_definition'" json:"Icd_definition"`
	}
	IcdCodeResp struct {
		Errno  int       `json: "errno"`
		Errmsg string    `json: "errmsg"`
		Data   []IcdCode `json: "data"`
	}

	IcdDrgCode struct {
		gorm.Model
		Icdr_id       int    `gorm: "column:'icdr_id';type:'integer'" json:"Icdr_id"`
		Icdr_icd_code string `gorm: "column:'icdr_icd_code';type:'varchar(32)'" json:"Icdr_icd_code"`
		Icdr_drg_code string `gorm: "column:'icdr_drg_code';type:'varchar(32)'" json:"Icdr_drg_code"`
		Icdr_title    string `gorm: "column:'icdr_title'" json:"Icdr_title"`
	}

	IcdDrgCodeResp struct {
		Errno  int          `json: "errno"`
		Errmsg string       `json: "errmsg"`
		Data   []IcdDrgCode `json: "data"`
	}

	User struct {
		User_id     int    `json:"User_id"`
		User_dn     string `gorm: "user_dn" json:"User_dn"`
		User_code   string `gorm: "user_code "json:"User_code"`
		User_name   string `gorm: "user_name "json:"User_name"`
		User_passwd string `gorm: "user_passwd "json:"User_passwd"`
	}

	UserResp struct {
		Errno  int    `json: "errno"`
		Errmsg string `json: "errmsg"`
		Data   []User `json: "data"`
	}
)

const (
	alphanumset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	dbtype = "mssql"
)

func (Keycode) TableName() string {
	return "keycodes"
}

func (Keycoderef) TableName() string {
	return "keycode_refs"
}

func (IcdCode) TableName() string {
	return "icdcodes"
}

func (DrgCode) TableName() string {
	return "drgcodes"
}

func (IcdDrgCode) TableName() string {
	return "icd10_drgs"
}

func GetRandStr(charset string, charlen int) string {
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
	v.SetConfigName("cdidemo")
	v.SetConfigType("json")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func connectSQL() (*gorm.DB, error) {
	cfg, err := getConfig()
	dbtype = cfg.GetString("HERMES_DATABASES_TYPE")
	//fmt.Println(cfg)
	if err != nil {
		return nil, err
	}
	if dbtype == "mssql" {
		connstr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			cfg.GetString("HERMES_DATABASES_MSSQLDBSIM_HOST"),
			cfg.GetString("hermes_databases_mssqldbsim_user"),
			cfg.GetString("HERMES_DATABASES_MSSQLDBSIM_PASSWORD"),
			cfg.GetInt("HERMES_DATABASES_MSSQLDBSIM_PORT"),
			cfg.GetString("HERMES_DATABASES_MSSQLDBSIM_DBNAME"))
		return gorm.Open(dbtype, connstr)
	} else if dbtype == "pgsql" {
		connstr := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
			cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_HOST"),
			cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_USER"),
			cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_PASSWORD"),
			cfg.GetInt("HERMES_DATABASES_PGSQLDBSIM_PORT"),
			cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_DBNAME"))
		return gorm.Open("postgres", connstr)
	} else {
		dbname := "memory.db"
		return gorm.Open(dbtype, dbname)
	}
}

func GetPermRBAC(e *casbin.Enforcer, req []string) bool {
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
	/*
		dbtype = cfg.GetString("HERMES_DATABASES_TYPE")
		var a *gormadapter.Adapter
		if dbtype == "mssql" {
			connstr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
				cfg.GetString("HERMES_DATABASES_MSSQLDBSIM_HOST"),
				cfg.GetString("hermes_databases_mssqldbsim_user"),
				cfg.GetString("HERMES_DATABASES_MSSQLDBSIM_PASSWORD"),
				cfg.GetInt("HERMES_DATABASES_MSSQLDBSIM_PORT"))
			//,cfg.GetString("HERMES_DATABASES_MSSQLDBSIM_DBNAME"))
			a = gormadapter.NewAdapter("mssql", connstr)
		} else if dbtype == "pgsql" {
			connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
				cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_USER"),
				cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_PASSWORD"),
				cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_HOST"),
				cfg.GetInt("HERMES_DATABASES_PGSQLDBSIM_PORT"))
			//,cfg.GetString("HERMES_DATABASES_PGSQLDBSIM_DBNAME"))
			a = gormadapter.NewAdapter("postgres", connstr)
		} else {
			a = gormadapter.NewAdapter("sqlite3", "./casbin.db")
		}
		log.Info(a)
		//e := casbin.NewEnforcer("model.conf", a)
	*/
	e := casbin.NewEnforcer(cfg.GetString("HERMES_CASBIN_MODEL"), cfg.GetString("HERMES_CASBIN_POLICY"))
	return GetPermRBAC(e, []string{userid, path, action})
}

func getCodeRecs(filterStr string, filterType int) KeycodeResp {
	Outobj := KeycodeResp{0, Errorcodes[0], []Keycode{}}
	db, err := connectSQL()
	if err != nil {
		log.Error("Error creating gorm object in data(): ", err)
		return KeycodeResp{1, Errorcodes[1], []Keycode{}}
	}
	log.Error(db)
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&Keycode{})

	tbl := db.Table("keycodes")
	tbl = tbl.Select("keycodes.id,keycodes.key_code,keycodes.key_title,keycodes.key_status,keycodes.key_definition,keycodes.key_app")
	if filterType == 1 {
		tbl = tbl.Where("keycodes.key_code like ?", filterStr)
	} else {
		filterList := strings.Split(filterStr, " ")
		for _, a := range filterList {
			if len(a) == 0 {
				continue
			}
			tbl = tbl.Where("keycodes.key_title like ? OR keycodes.key_definition like ?", fmt.Sprintf("%%%s%%", a), fmt.Sprintf("%%%s%%", a))
		}
	}
	tbl.Find(&Outobj.Data)
	for i := 0; i < len(Outobj.Data); i++ {
		Outobj.Data[i].Key_definition = strings.Replace(Outobj.Data[i].Key_definition, "\\r\\n", "\n", -1)
		Outobj.Data[i].Key_definition = strings.Replace(Outobj.Data[i].Key_definition, "\\t", "\t", -1)
	}
	return Outobj
}

func getCodeRefRecs(code string) KeycoderefResp {
	Outobj := KeycoderefResp{0, Errorcodes[0], []Keycoderef{}}

	db, err := connectSQL()
	if err != nil {
		log.Error("Error creating gorm object in data: ", err)
		return KeycoderefResp{1, Errorcodes[1], []Keycoderef{}}
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&Keycoderef{})

	tbl := db.Table("keycode_refs")
	tbl = tbl.Select("keycode_refs.id, keycode_refs.kcr_key_code, keycode_refs.kcr_title, keycode_refs.kcr_status, keycode_refs.kcr_definition, keycode_refs.kcr_app")
	tbl = tbl.Where("keycode_refs.kcr_key_code like ?", code)

	tbl.Find(&Outobj.Data)
	return Outobj
}

func getDRGRecs(filterStr string, filterType int) DrgCodeResp {
	Outobj := DrgCodeResp{0, Errorcodes[0], []DrgCode{}}

	db, err := connectSQL()
	if err != nil {
		fmt.Println(err)
		return DrgCodeResp{1, Errorcodes[1], []DrgCode{}}
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&DrgCode{})

	tbl := db.Table("drgs")
	tbl = tbl.Select("drgs.id,drgs.drg_code,drgs.drg_title,drgs.drg_definition")
	if filterType == 1 {
		tbl = tbl.Where("drgs.drg_code like ?", filterStr)
	} else {
		filterList := strings.Split(filterStr, " ")
		for _, a := range filterList {
			if len(a) == 0 {
				continue
			}
			tbl = tbl.Where("drgs.drg_title like ? OR drgs.drg_definition like ?", fmt.Sprintf("%%%s%%", a), fmt.Sprintf("%%%s%%", a))
		}
	}

	tbl.Find(&Outobj.Data)
	for i := 0; i < len(Outobj.Data); i++ {
		Outobj.Data[i].Drg_definition = strings.Replace(Outobj.Data[i].Drg_definition, "\\r\\n", "\n", -1)
		Outobj.Data[i].Drg_definition = strings.Replace(Outobj.Data[i].Drg_definition, "\\t", "\t", -1)
	}
	return Outobj
}

func getICD10Recs(filterStr string, filterType int) IcdCodeResp {
	Outobj := IcdCodeResp{0, Errorcodes[0], []IcdCode{}}

	db, err := connectSQL()
	if err != nil {
		fmt.Println(err)
		return IcdCodeResp{1, Errorcodes[1], []IcdCode{}}
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&IcdCode{})

	tbl := db.Table("icd10s")
	tbl = tbl.Select("icd10s.id,icd10s.icd_code,icd10s.icd_title,icd10s.icd_definition")

	if filterType == 1 {
		tbl = tbl.Where("icd10s.icd_code like ?", filterStr)
	} else {
		filterList := strings.Split(filterStr, " ")
		for _, a := range filterList {
			if len(a) == 0 {
				continue
			}
			tbl = tbl.Where("icd10s.icd_title like ? OR icd10s.icd_definition like ?", fmt.Sprintf("%%%s%%", a), fmt.Sprintf("%%%s%%", a))
		}
	}
	tbl.Find(&Outobj.Data)
	fmt.Println("Outobj.Data:", Outobj.Data)

	for i := 0; i < len(Outobj.Data); i++ {
		Outobj.Data[i].Icd_definition = strings.Replace(Outobj.Data[i].Icd_definition, "\\r\\n", "\n", -1)
		Outobj.Data[i].Icd_definition = strings.Replace(Outobj.Data[i].Icd_definition, "\\t", "\t", -1)

	}
	return Outobj
}

func getICDRGRecs(code string, codetype int) IcdDrgCodeResp { //codetype{0:icd,1:drg}
	Outobj := IcdDrgCodeResp{0, Errorcodes[0], []IcdDrgCode{}}

	db, err := connectSQL()
	if err != nil {
		fmt.Println(err)
		return IcdDrgCodeResp{1, Errorcodes[1], []IcdDrgCode{}}
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&IcdDrgCode{})

	tbl := db.Table("icd10_drgs")
	tbl = tbl.Select("icd10_drgs.id,icd10_drgs.Icdr_icd_code,icd10_drgs.Icdr_drg_code,icd10_drgs.icdr_title")
	if codetype == 0 {
		tbl = tbl.Where("icd10_drgs.Icdr_icd_code like ?", code)
	} else {
		tbl = tbl.Where("icd10_drgs.Icdr_drg_code like ?", code)
	}
	//tbl.Find(&Outobj.Data)

	rows, _ := tbl.Rows()
	defer rows.Close()
	for rows.Next() {
		//line := IcdDrgCode{}
		var icdr_icd_code, icdr_drg_code, icdr_title string
		var icdr_id int
		rows.Scan(&icdr_id, &icdr_icd_code, &icdr_drg_code, &icdr_title)
		line := IcdDrgCode{}
		line.Icdr_id = icdr_id
		line.Icdr_drg_code = icdr_drg_code
		line.Icdr_icd_code = icdr_icd_code
		line.Icdr_title = icdr_title
		Outobj.Data = append(Outobj.Data, line)
	}
	return Outobj
}

func CreateUserResp(e *ldap.Entry) User {
	ur := User{}
	ur.User_dn = e.DN
	for i := 0; i < len(e.Attributes); i++ {
		if e.Attributes[i].Name == "objectClass" {
			ur.User_dn = e.Attributes[i].Values[0]
		} else if e.Attributes[i].Name == "sn" {
			ur.User_name = e.Attributes[i].Values[0]
		} else if e.Attributes[i].Name == "cn" {
			ur.User_code = e.Attributes[i].Values[0]
		} else if e.Attributes[i].Name == "userPassword" {
			ur.User_passwd = e.Attributes[i].Values[0]
		}
		ur.User_id = i
	}
	return ur
}

func getUserRecs(userid, passwd string) UserResp {
	Outobj := UserResp{0, Errorcodes[0], []User{}}
	cfg, err := getConfig()

	l, err := ldap.Dial(cfg.GetString("HERMES_SERVERS_LDAP_PROTOCOL"), fmt.Sprintf("%v:%v", cfg.GetString("HERMES_SERVERS_LDAP_URL"), cfg.GetString("HERMES_SERVERS_LDAP_PORT")))
	//l, err := ldap.Dial("tcp", "localhost:389")
	//fmt.Println("ldap conn:", l, err)
	if err != nil {
		log.Error("LDAP Dial/Connect failure:", err)
	}
	defer l.Close()

	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Error("StartTLS failed:", err)
	}

	err = l.Bind("cn=admin,dc=corp,dc=accds,dc=com", "30dFokig2e0")
	if err != nil {
		log.Error("LDAP Binding error with:", cfg.GetString("HERMES_LDAP_BINDUSERNAME"), "-->", err)
	}

	searchRequest := ldap.NewSearchRequest("dc=corp,dc=accds,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", userid),
		nil, nil,
	)
	//fmt.Println("getUserRecs Search Request:", searchRequest)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Error("User Search Failed:", err)
	}
	//fmt.Println("getUserRecs Search Result:", sr)

	if len(sr.Entries) == 1 {
		ur := CreateUserResp(sr.Entries[0])
		//fmt.Println("getUserRecs UserResp:", ur)
		//fmt.Println("getUserRecs Passwd:", ur.User_passwd, passwd)
		if ur.User_passwd == passwd {
			Outobj.Data = []User{ur}
		}
	}
	return Outobj
}
