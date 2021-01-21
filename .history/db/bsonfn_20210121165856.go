package db

import (
	"iii/ifactory/compute/util"

	// . "iii/ifactory/compute/util/cch/json"
	"reflect"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// func L() {
// 	// var i []interface{}
// 	var i []model.WorkOrder
// 	err := Lookup(model.C.Workorder, model.C.Workorder_list, "WorkOrderId", &i)
// 	fmt.Println(err)
// 	util.PrintJson(i)
// 	for _, o := range i {
// 		fmt.Println(o)
// 	}
// }

type Agg struct {
}

func (a *Agg) GenUnwind(nestedObject interface{}) interface{} {
	return map[string]interface{}{
		"$unwind": "$" + nestedObject.(string),
	}
}

func (a *Agg) GenGroup(nestedObject string, groups, firsts, sums []string) interface{} {
	nestedObject = "$" + nestedObject + "."

	m:=map[string]interface{}{}
	grouped := map[string]interface{}{}	
	for _, g := range groups {
		grouped[g] = nestedObject + g
	}
	m["_id"]:=grouped

	for _, f := range firsts {
		m[s] = bson.M{"$sum": s}
	}

	for _, s := range sums {
		m[f] = bson.M{"$first": f}
	}

	return map[string]interface{}{
		"$group": m,
	}
}

//關聯查詢 副表物件會存入SubCollection
func Lookup(collection, subCollection string, primaryKey string, result interface{}) error {
	lookup := map[string]interface{}{
		"from":         subCollection,
		"localField":   primaryKey,
		"foreignField": primaryKey,
		"as":           "SubCollection",
	}

	var querys []interface{}
	querys = append(
		querys,
		map[string]interface{}{
			"$lookup": lookup,
		},
	)

	err := MongoDB.UseC(collection).Pipe(querys).All(result)
	if err != nil {
		return err
	}
	return nil
}

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

func UpdateOne(collection string, selector interface{}, update interface{}) error {
	c := MongoDB.UseC(collection)
	updated := bson.M{
		"$set": update,
	}
	err := c.Update(selector, updated)
	if err != nil {
		return err
	}
	return nil
}

func Push(collection string, selector interface{}, update interface{}) error {
	c := MongoDB.UseC(collection)
	updated := bson.M{
		"$push": update,
	}
	// util.PrintJson(selector)
	// util.PrintJson(updated)
	err := c.Update(selector, updated)
	if err != nil {
		return err
	}
	return nil
}

func Pushs(collection string, selector interface{}, field string, update interface{}) error {
	c := MongoDB.UseC(collection)

	updated := bson.M{
		"$push": bson.M{
			field: bson.M{"$each": update},
		},
	}

	err := c.Update(selector, updated)
	if err != nil {
		glog.Error(util.Cerr(err))
		return err
	}
	return nil
}

func ToIfaces(iface interface{}) (ifaces []interface{}) {
	switch reflect.TypeOf(iface).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(iface)
		for i := 0; i < s.Len(); i++ {
			// fmt.Println(s.Index(i))
			ifaces = append(ifaces, s.Index(i).Interface())
		}
	}
	return ifaces
}

func Insert(collection string, v interface{}) {
	c := MongoDB.UseC(collection)
	err := c.Insert(v)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
}

//查詢collection中所有資料(根據deviceid,groupid,parentid其中之一) ---#to be deprecated
func Match(collection string, searchs interface{}, result interface{}) error {
	err := MongoDB.UseC(collection).Pipe(searchs).All(result)
	if err != nil {
		return err
	}
	return nil
}

func Upsert(collection string, query bson.M, setvalue interface{}) (info *mgo.ChangeInfo) {
	c := MongoDB.UseC(collection)
	i, err := c.Upsert(query, setvalue)
	util.PrintJson(i)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return i
}

func Update(collection string, option interface{}, setvalue bson.M) {
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
