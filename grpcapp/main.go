package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	//	"cdidemo-grpc/sample"

	"google.golang.org/grpc"
)

type (
	server struct {
		UnimplementedCdidemoServer
	}
)

const (
	port = ":8801"
)

func (s *server) AuthUser(ctx context.Context, in *UserRequest) (*ValidResponse, error) {
	log.WithFields(log.Fields{"userid": in.GetUserid(), "type": "request"}).Info("AuthUser")
	rec := getUserRecs(in.GetUserid(), in.GetPasswd())
	var valid bool
	if len(rec.Data) == 1 {
		valid = true
	} else {
		valid = false
	}
	log.WithFields(log.Fields{"valid": valid, "userid": in.GetUserid(), "type": "response"}).Info("AuthUser")
	return &ValidResponse{Valid: valid}, nil
}

func (s *server) GetUser(ctx context.Context, in *UserRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"userid": in.GetUserid()}).Info("GetUser1")
	rec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if !authrizing(in.GetUserid(), "/user", "read") {
		log.WithFields(log.Fields{"search": in.GetUserid(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(UserResp{3, Errorcodes[3], []User{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recsJSON, _ := json.Marshal(rec)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetUser2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetKeyCode(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	fmt.Println("GetKeyCode()")
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetKeyCode1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	fmt.Println(userrec)
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(KeycodeResp{3, Errorcodes[3], []Keycode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}

	if !authrizing(in.GetUserid(), "/code", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(KeycodeResp{3, Errorcodes[3], []Keycode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}

	rec := getCodeRecs(in.GetSearch(), 0)
	recsJSON, _ := json.Marshal(rec)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetKeyCode2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetKeyCodeDetail(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetKeyCodeDetail1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(KeycodeResp{3, Errorcodes[3], []Keycode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/code/detail", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(KeycodeResp{3, Errorcodes[3], []Keycode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getCodeRecs(in.GetSearch(), 1)
	//fmt.Println(len(recs))
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetKeyCodeDetail2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetKeyCodeRef(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetKeyCodeRef1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(KeycoderefResp{3, Errorcodes[3], []Keycoderef{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/code/ref", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(KeycoderefResp{3, Errorcodes[3], []Keycoderef{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getCodeRefRecs(in.GetSearch())
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetKeyCodeRef2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetICD10(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetICD101")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(IcdCodeResp{3, Errorcodes[3], []IcdCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/icd10", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(IcdCodeResp{3, Errorcodes[3], []IcdCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getICD10Recs(in.GetSearch(), 0)
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetICD102")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetICD10Detail(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetICD10Detail1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(IcdCodeResp{3, Errorcodes[3], []IcdCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/icd10/detail", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(IcdCodeResp{3, Errorcodes[3], []IcdCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getICD10Recs(in.GetSearch(), 1)
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetICD10Detail2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetICDref(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetICD10Ref1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(IcdDrgCodeResp{3, Errorcodes[3], []IcdDrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/icd10/ref", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(IcdDrgCodeResp{3, Errorcodes[3], []IcdDrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getICDRGRecs(in.GetSearch(), 0)
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetICD10Ref2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetDRG(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetDRG1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(DrgCodeResp{3, Errorcodes[3], []DrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/drg", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(DrgCodeResp{3, Errorcodes[3], []DrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getDRGRecs(in.GetSearch(), 0)
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetDRG2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetDRGDetail(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetDRGDetail1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(DrgCodeResp{3, Errorcodes[3], []DrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/drg/detail", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(DrgCodeResp{3, Errorcodes[3], []DrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getDRGRecs(in.GetSearch(), 1)
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetDRGDetail2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func (s *server) GetDRGref(ctx context.Context, in *SearchRequest) (*JsonResponse, error) {
	log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "request"}).Info("GetDRGRef1")
	userrec := getUserRecs(in.GetUserid(), in.GetPasswd())
	if len(userrec.Data) != 1 {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authenticated")
		recsJSON, _ := json.Marshal(IcdDrgCodeResp{3, Errorcodes[3], []IcdDrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	if !authrizing(in.GetUserid(), "/drg/ref", "read") {
		log.WithFields(log.Fields{"search": in.GetSearch(), "userid": in.GetUserid(), "type": "error"}).Panic("Not authorized")
		recsJSON, _ := json.Marshal(IcdDrgCodeResp{3, Errorcodes[3], []IcdDrgCode{}})
		return &JsonResponse{Jsondata: string(recsJSON)}, nil
	}
	recs := getICDRGRecs(in.GetSearch(), 1)
	recsJSON, _ := json.Marshal(recs)
	log.WithFields(log.Fields{"result": string(recsJSON), "userid": in.GetUserid(), "type": "response"}).Info("GetICD10Ref2")
	return &JsonResponse{Jsondata: string(recsJSON)}, nil
}

func main() {
	//	go sample.KeepRun()
	lf, err := os.OpenFile(fmt.Sprintf("log/cdidemo_%s.log", time.Now().Format(time.RFC3339)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Info("Failed to log to file, using default stderr")
	}
	defer lf.Close()
	mw := io.MultiWriter(os.Stdout, lf)
	log.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(mw)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterCdidemoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
