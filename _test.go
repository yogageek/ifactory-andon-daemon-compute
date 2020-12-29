package main

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	host     = "52.187.110.12:27017"
	source   = "ifp-data-hub"
	username = "8676b401-a6ce-417b-a0f0-8dee6dee0a67"
	password = "9TuSZ7CD3ah0aQmdHbGqNjrr"
)

var (
	db         = "ifp-data-hub"
	collection = "iii.dae.AbnormalMachineLatest"
)

var session *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host}, // 資料庫位址
		Source:   source,         // 設置權限的資料庫
		Username: username,       // 帳號
		Password: password,       // 密碼
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	session = s
}

func connect(db string, collection string) (*mgo.Session, *mgo.Collection) {
	// 每一次操作，都複製一份 session，避免每次操作都創建 session，導致連線數超過設置的最大值
	s := session.Copy()
	// 獲取資料表
	c := s.DB(db).C(collection)
	return s, c
}

// FindAll will find all resources.
func FindAll(db string, collection string, query interface{}, selector interface{}, result interface{}) error {
	s, c := connect(db, collection)
	// 主動關閉 session
	defer s.Close()

	return c.Find(query).Select(selector).All(result)
}

// abnormalMachineLatest struct
type abnormalMachineLatest struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	UpdateTime time.Time     `json:"UpdateTime" bson:"UpdateTime,omitempty"`
	Type       string        `json:"Type" bson:"Type,omitempty"`
	GroupID    string        `json:"GroupID" bson:"GroupID,omitempty"`
	GroupName  string        `json:"GroupName" bson:"GroupName,omitempty"`
	MachineID  string        `json:"MachineID" bson:"MachineID,omitempty"`
}

// FindAll will find all movies.
func (m *abnormalMachineLatest) FindAll() ([]abnormalMachineLatest, error) {
	var a []abnormalMachineLatest
	err := FindAll(db, collection, nil, nil, &a)
	return a, err
}

func FindTest() {
	//直接拿struct
	var i []interface{}
	err := FindAll(db, collection, nil, nil, &i)
	if err != nil {
		log.Fatal(err)
	}

	//亂碼
	bb, err := bson.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshal:", string(bb))

	//interface -> json of bsondata
	b, err := bson.MarshalJSON(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MarshalJSON:", string(b))

	//json of bsondata -> struct
	a := []abnormalMachineLatest{}
	err = bson.UnmarshalJSON(b, &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("unmarshalJson:", a)
}

func main() {
	FindTest()
	var (
		model = abnormalMachineLatest{}
	)

	var ab []abnormalMachineLatest

	ab, err := model.FindAll()
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", ab)

}
