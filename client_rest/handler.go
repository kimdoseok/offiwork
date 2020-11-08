package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type (
	PostData struct {
		Userid string `json: "userid"`
		Passwd string `json: "passwd"`
		data   interface{}
	}
	Outobj struct {
		Errno  int
		Errmsg string
		data   interface{}
	}
)

func getClientList(w http.ResponseWriter, r *http.Request) {
}
func getClientDetail(w http.ResponseWriter, r *http.Request) {
}
func getClientDelete(w http.ResponseWriter, r *http.Request) {
}
func getClientSave(w http.ResponseWriter, r *http.Request) {
}

func getRestCode(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/code", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	filterStr := ctx.Value("filterStr").(string)
	recs := getCodeRecs(filterStr, 0)
	recsJSON, _ := json.Marshal(recs)
	w.Write(recsJSON)
}

func getRestCodeDetail(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/code/detail", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	codeStr := ctx.Value("filterStr").(string)
	recs := getCodeRecs(codeStr, 1)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func authUser(userid, passwd, resource, action string) bool {
	log.WithFields(log.Fields{"userid": userid, "type": "request"}).Info("AuthUser")
	rec := getUserRecs(userid, passwd)
	if rec.Errno != 0 && len(rec.Data) != 1 {
		log.WithFields(log.Fields{"search": "", "userid": userid, "type": "error"}).Panic("Fetching data error.")
		return false
	}
	if !authrizing(userid, resource, action) {
		log.WithFields(log.Fields{"search": "", "userid": userid, "type": "error"}).Panic("Not authorized")
		return false
	}
	return true
}

func getRestUser(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	fmt.Println("authUser:", authUser(postData.Userid, postData.Passwd, "/user", "read"))
	if !authUser(postData.Userid, postData.Passwd, "/user", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	recs := getUserRecs(postData.Userid, postData.Passwd)
	recsJSON, _ := json.Marshal(recs)
	w.Write(recsJSON)
}

func getRestCodeRef(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/code/ref", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	codeStr := ctx.Value("filterStr").(string)
	recs := getCodeRefRecs(codeStr)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func getRestICD10(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/icd10", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	filterStr := ctx.Value("filterStr").(string)
	recs := getICD10Recs(filterStr, 0)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func getRestICD10Detail(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/drg/detail", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	codeStr := ctx.Value("filterStr").(string)
	recs := getICD10Recs(codeStr, 1)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func getRestICDref(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/icd10/ref", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	filterStr := ctx.Value("filterStr").(string)
	recs := getICDRGRecs(filterStr, 0)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func getRestDRG(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	//fmt.Println("authUserInput:", postData, err)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/drg", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	filterStr := ctx.Value("filterStr").(string)
	recs := getDRGRecs(filterStr, 0)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func getRestDRGDetail(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/drg/detail", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	codeStr := ctx.Value("filterStr").(string)
	recs := getDRGRecs(codeStr, 1)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}

func getRestDRGref(w http.ResponseWriter, r *http.Request) {
	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Decode error.")
		w.WriteHeader(503)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	if !authUser(postData.Userid, postData.Passwd, "/drg/ref", "read") {
		log.WithFields(log.Fields{"search": "", "userid": postData.Userid, "type": "error"}).Panic("Not authorized")
		w.WriteHeader(401)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	filterStr := ctx.Value("filterStr").(string)
	recs := getICDRGRecs(filterStr, 1)
	recsJSON, _ := json.Marshal(recs)
	fmt.Println(string(recsJSON))
	w.Write(recsJSON)
}
