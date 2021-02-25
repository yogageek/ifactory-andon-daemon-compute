package db

import (
	// _ "iii/m2i/v1/setenv" //可以使用_ "github.com/xxx/xxx"引入包，提前在這支前做init裡的方法

	"os"
	"time"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

//inition
var (
	mongodb_url      string
	mongodb_database string
	mongodb_username string
	mongodb_password string

	//public mongo connection
	MongoDB *mongoDB
)

type mongoDB struct {
	Db *mgo.Database
}

func MongoHealCheckLoop() {
	d := time.Second * time.Duration(60)

	check := func() {
		s := MongoDB.Db.Session
		if err := s.Ping(); err != nil {
			glog.Error("MongoHealCheckLoop Err:", err)
			MongoDB = NewMongo()
		}
	}

	for {
		check()
		time.Sleep(d)
	}
}

func StartMongo() {
	mongodb_url = os.Getenv("MONGODB_URL")
	mongodb_database = os.Getenv("MONGODB_DATABASE")
	mongodb_username = os.Getenv("MONGODB_USERNAME")
	mongodb_password = os.Getenv("MONGODB_PASSWORD")
	MongoDB = NewMongo()
}

func NewMongo() *mongoDB {
	return &mongoDB{
		Db: createConnection(),
	}
}

func createConnection() *mgo.Database {
	glog.Infoln("create mongodb connection...")
	session, err := mgo.Dial(mongodb_url)

	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB(mongodb_database)
	err = db.Login(mongodb_username, mongodb_password)
	if err != nil {
		panic(err)
	}
	return db
}

// mgo method---------------------------------------------------

func (mongodb *mongoDB) UseC(collection string) *mgo.Collection {
	return mongodb.Db.C(collection)
}

// -------------------------------------------------

// Go original mongo library (more compplecated)
/*
type mongoDB struct {
	mongoClient *mongo.Client
}

func NewMongo() *mongoDB {
	return &mongoDB{
		mongoClient: createConnection(),
	}
}

func createConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://10.0.1.19:27017,10.0.1.20:27017,10.0.1.21:27017")
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
*/
