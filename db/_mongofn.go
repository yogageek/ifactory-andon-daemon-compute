package db

import (
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

func FindAll(collection string) []bson.M {
	c := MongoDB.UseC(collection)
	resp := []bson.M{}
	err := c.Find(nil).All(&resp)
	if err != nil {
		glog.Error(util.Cerr(err))
		return resp
	}
	// util.PrintJson(resp)
	return resp
}

//注意有些可以用options ...bson.M ; 有些只能用option bson.M

//查詢collection中所有資料(根據deviceid,groupid,parentid其中之一) ---#to be deprecated
func MatchBySearchOptions(collection string, searchs ...bson.M) []bson.M {
	c := MongoDB.UseC(collection)
	resp := []bson.M{}
	err := c.Pipe(searchs).All(&resp)
	if err != nil {
		glog.Error(util.Cerr(err))
		return resp
	}
	// util.PrintJson(resp)
	return resp
}

//查詢collection中所有資料(根據deviceid,groupid,parentid其中之一)
func MatchBySearchOptionsI(collection string, searchs ...bson.M) (resp []interface{}) {
	c := MongoDB.UseC(collection)
	err := c.Pipe(searchs).All(&resp)
	if err != nil {
		glog.Error(util.Cerr(err))
		return resp
	}
	// util.PrintJson(resp)
	return resp
}

func QueryAll(collection string) (resp []interface{}) {
	resp = MatchBySearchOptionsI(collection)
	return resp
}

func Insert(collection string, i interface{}) {
	c := MongoDB.UseC(collection)
	err := c.Insert(i)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func UpsertByOption(collection string, query bson.M, setvalue interface{}) {
	c := MongoDB.UseC(collection)
	i, err := c.Upsert(query, setvalue)
	util.PrintJson(i)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func UpdateByOption(collection string, option bson.M, setvalue bson.M) {
	c := MongoDB.UseC(collection)
	err := c.Update(option, setvalue)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func UpdateAllByOption(collection string, setvalue bson.M, option bson.M) {
	c := MongoDB.UseC(collection)
	i, err := c.UpdateAll(option, setvalue) //UpdateAll(搜尋條件,修改規則)
	util.PrintJson(i)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

func CountByOption(collection string, option bson.M) int {
	c := MongoDB.UseC(collection)
	num, err := c.Find(option).Count()
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return num
}

func RemoveAllByOption(collection string, option bson.M) {
	c := MongoDB.UseC(collection)
	_, err := c.RemoveAll(option)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

//depre
func FindByOption(collection string, option bson.M) []bson.M {
	c := MongoDB.UseC(collection)
	resp := []bson.M{}
	err := c.Find(option).All(&resp)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return resp
}

func FindOneByOption(collection string, option bson.M) bson.M {
	c := MongoDB.UseC(collection)
	resp := bson.M{}
	err := c.Find(option).One(&resp)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return resp
}

// func FindOneByOptionErr(collection string, option bson.M) (bson.M, error) {
// 	c := MongoDB.UseC(collection)
// 	resp := bson.M{}
// 	err := c.Find(option).One(&resp)
// 	if err != nil {
// 		return resp, err
// 	}
// 	return resp, nil
// }

//Select用法(只取要指定key的值出來) ->"keyname":1
//Use the query Select method to specify the fields to return
func FindOneByOptionAndSelect(collection string, option bson.M, field string) bson.M {
	c := MongoDB.UseC(collection)
	resp := bson.M{}

	selector := bson.M{field: 1}

	err := c.Find(option).Select(selector).One(&resp)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return resp
}

func Drop(collection string) {
	c := MongoDB.UseC(collection)
	err := c.DropCollection()
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}
