package db

import (
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

// FindAll will find all resources.
func FindAll(collection string, query interface{}, selector interface{}, result interface{}) error {
	// Query method to specify the fields to query
	// Select method to specify the fields
	err := MongoDB.UseC(collection).Find(query).Select(selector).All(result)
	if err != nil {
		return err
	}
	return nil
}

//查詢collection中所有資料(根據deviceid,groupid,parentid其中之一) ---#to be deprecated
func Match(collection string, searchs interface{}, result interface{}) error {
	err := MongoDB.UseC(collection).Pipe(searchs).All(result)
	if err != nil {
		return err
	}
	return nil
}

func Upsert(collection string, query bson.M, setvalue interface{}) {
	c := MongoDB.UseC(collection)
	i, err := c.Upsert(query, setvalue)
	util.PrintJson(i)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func Insert(collection string, v interface{}) {
	c := MongoDB.UseC(collection)
	err := c.Insert(v)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func Update(collection string, option bson.M, setvalue bson.M) {
	c := MongoDB.UseC(collection)
	err := c.Update(option, setvalue)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func Remove(collection string, selector interface{}) {
	c := MongoDB.UseC(collection)
	err := c.Remove(selector)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}
