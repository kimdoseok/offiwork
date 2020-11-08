package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8801"
	userid  = "882966"
	passwd  = "c@sLdGgxI[sE|aJ"
)

func TestMSSQL(t *testing.T) {
	fmt.Println("You must run backend first!")
	_, err := getConfig()
	if err != nil {
		t.Error("Configuration loading Error.", err)
	}
	_, err = connectSQL()
	//fmt.Println(err)
	if err != nil {
		t.Error("MSSQL Connection Error.", err)
	}

}

func TestLDAP(t *testing.T) {
	t.Helper()
	cfg, err := getConfig()
	if err != nil {
		t.Error("Configuration loading Error.", err)
	}
	cfg.Set("HERMES_SERVERS_LDAP_URL", "localhost")
	//fmt.Println("Config:", cfg)

	recs := getUserRecs(userid, passwd)
	fmt.Println("TestLDAP:", recs)
}

func TestKeyCode(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	searchStr := "c"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetKeyCode(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j KeycodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	fmt.Println("j.Errmsg: ", j.Errmsg, err)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("KeyCode Response: %v %s", j, reflect.TypeOf(j))
}

func TestKeyCodeDetail(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "K9171"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetKeyCodeDetail(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	var j KeycodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("KeyCodeDetail Response: %v %s", j, reflect.TypeOf(j))
}

func TestKeyCodeRef(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "K9171"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetKeyCodeRef(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j KeycoderefResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("KeyCodeRef Response: %v %s", j, reflect.TypeOf(j))

}

func TestICD10(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "c"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetICD10(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	fmt.Println("Response: ", r.GetJsondata())
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j IcdCodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("ICD10 Response: %v %s", j, reflect.TypeOf(j))

}
func TestICD10Detail(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "R7881"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetICD10Detail(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j IcdCodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("ICD10Detail Response: %v %s", j, reflect.TypeOf(j))

}
func TestICD10ref(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "R7881"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetICDref(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j IcdDrgCodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("ICD10ref Response: %v %s", j, reflect.TypeOf(j))

}

func TestDRG(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "c"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetDRG(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	log.Printf("DRG Response: %s %s", r.GetJsondata(), reflect.TypeOf(r.GetJsondata()))

	var j DrgCodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("DRG Response: %v %s", j, reflect.TypeOf(j))

}

func TestDRGDetail(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	searchStr := "195"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetDRGDetail(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j DrgCodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("DRGDetail Response: %v %s", j, reflect.TypeOf(j))

}

func TestDRGref(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	authctx, authCancel := context.WithTimeout(context.Background(), time.Second)
	defer authCancel()
	ar, err := c.AuthUser(authctx, &UserRequest{Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	if !ar.GetValid() {
		return
	}

	// Contact the server and print out its response.
	searchStr := "195"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetDRGref(ctx, &SearchRequest{Search: searchStr, Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j IcdDrgCodeResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("DRGref Response: %v %s", j, reflect.TypeOf(j))

}

func TestAuth(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AuthUser(ctx, &UserRequest{Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	log.Printf("Auth Response: %v %s", r.GetValid(), reflect.TypeOf(r.GetValid()))

}

func TestGetUser(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewCdidemoClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &UserRequest{Userid: userid, Passwd: passwd})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	var j UserResp
	err = json.Unmarshal([]byte(r.GetJsondata()), &j)
	if err != nil {
		log.Fatalf("invalid jsonformat: %v", err)
	}
	log.Printf("GetUser Response: %v %v", j, reflect.TypeOf(j))

}
