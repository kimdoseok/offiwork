package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
type (
	Parent struct {
		Level string `json: level`
		Type string `json: type`
		Msg string `json: msg`
		Result string  `json:result`
		Valid bool `json: valid`
		Time string `json: time`
		Userid string `json: userid`
	}

	KeycodeResp struct {
		Parent
		Key_id         int    `json: "key_id"`
		Key_code       string `json: "key_code"`
		Key_title      string `json: "key_title"`
		Key_status     int    `json: "key_status"`
		Key_definition string `json: "key_definition"`
		Key_app        string `json: "key_app"`
		Key_cc         string `json: "key_cc"`
		Key_hcc        string `json: "key_hcc"`
		Key_mcc        string `json: "key_mcc"`
	}

	DrgCodeResp struct {
		Drg_id         int    `json: "drg_id"`
		Drg_code       string `json: "drg_code"`
		Drg_title      string `json: "drg_title"`
		Drg_definition string `json: "drg_definition"`
	}

	IcdCodeResp struct {
		Icd_id         int    `json: "icd_id"`
		Icd_code       string `json: "icd_code"`
		Icd_title      string `json: "icd_title"`
		Icd_definition string `json: "icd_definition"`
	}

	UserResp struct {
		User_id     int
		User_dn     string `json: "user_dn"`
		User_code   string `json: "user_code"`
		User_name   string `json: "user_name"`
		User_passwd string `json: "user_passwd"`
	}

	ValidResp struct {
		Parent
		Valid bool
	}

)
const (
	dirname = "./"
)
func saveKeylog(line Parent) {
	recs := []KeycodeResp{}
	err := json.Unmarshal([]byte(line.Result),&recs)
	if err != nil {
		log.Println("saveKeylog unmarshal error:",err)
	}
	//fmt.Println(line.Result,err)
	for i,rec := range recs {
		fmt.Println(i,line.Level)
		fmt.Println(i,line.Msg)
		fmt.Println(i,line.Time)
		fmt.Println(i,line.Userid)
		fmt.Println(i,rec.Key_id)
		fmt.Println(i,rec.Key_code)
		fmt.Println(i,rec.Key_app)
		fmt.Println(i,rec.Key_title)
		fmt.Println(i,rec.Key_definition)
		fmt.Println(i,rec.Key_status)
		fmt.Println(i,rec.Key_cc)
		fmt.Println(i,rec.Key_mcc)
		fmt.Println(i,rec.Key_hcc)
	}
}

func saveICDlog(line Parent) {
	recs := []IcdCodeResp{}
	err := json.Unmarshal([]byte(line.Result),&recs)
	if err != nil {
		log.Println("saveICDlog unmarshal error:",err)
	}
	for i,rec := range recs {
		fmt.Println(i,line.Level)
		fmt.Println(i,line.Msg)
		fmt.Println(i,line.Time)
		fmt.Println(i,line.Userid)
		fmt.Println(i,rec.Icd_id)
		fmt.Println(i,rec.Icd_code)
		fmt.Println(i,rec.Icd_title)
		fmt.Println(i,rec.Icd_definition)
	}
}

func saveDRGlog(line Parent) {
	recs := []DrgCodeResp{}
	err := json.Unmarshal([]byte(line.Result),&recs)
	if err != nil {
		log.Println("saveDRGlog Unmarshal Error:",err)
	}
	//fmt.Println(line.Result,err)
	for i,rec := range recs {
		fmt.Println(i,line.Level)
		fmt.Println(i,line.Msg)
		fmt.Println(i,line.Time)
		fmt.Println(i,line.Userid)
		fmt.Println(i,rec.Drg_id)
		fmt.Println(i,rec.Drg_code)
		fmt.Println(i,rec.Drg_title)
		fmt.Println(i,rec.Drg_definition)
	}
}

func saveUserlog(line Parent) {
	recs := []UserResp{}
	_ = json.Unmarshal([]byte(line.Result),&recs)
	//fmt.Println(line.Result,err)
	for i,rec := range recs {
		fmt.Println(i,line.Level)
		fmt.Println(i,line.Type)
		fmt.Println(i,line.Msg)
		fmt.Println(i,line.Time)
		fmt.Println(i,line.Userid)
		fmt.Println(i,rec.User_id)
		fmt.Println(i,rec.User_code)
		fmt.Println(i,rec.User_name)
		fmt.Println(i,rec.User_dn)
		//fmt.Println(i,rec.User_passwd)
	}

}

func loadJSON(jsonbytes []byte) {
	line := Parent{}
	err := json.Unmarshal(jsonbytes,&line)
	if err != nil {
        log.Fatal("loadJSON Unmarshal Error:",err)
	}
	msgrunes := []rune(line.Msg)
	if line.Msg=="GetUser" {
		saveUserlog(line)
	} else if string(msgrunes[0:6])=="GetDRG" {
		saveDRGlog(line)
	} else if string(msgrunes[0:8])=="GetICD10" {
		saveICDlog(line)
	} else if string(msgrunes[0:9])=="GetKeyCode" {
		saveKeylog(line)
	}
	
}

func loadfile(fn string) {
	//fmt.Println(fn)
	f, err := os.Open(dirname+fn)
	if err != nil {
        log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		line := scanner.Bytes()
		//fmt.Println(string(line))
		loadJSON(line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main() {
	files, err := ioutil.ReadDir("./")
    if err != nil {
        fmt.Println(err)
	}
	for _,fs := range files {
		fn := fs.Name()
		if fn[len(fn)-4:len(fn)]==".log" {
			loadfile(fn)
		}
	}
	fmt.Println("Done!")
}