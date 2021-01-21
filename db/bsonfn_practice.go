package db

import (
	model "iii/ifactory/compute/model/sfc"
	"iii/ifactory/compute/util"

	// . "iii/ifactory/compute/util/cch/json"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

type AggregateBoss struct {
	AggregateKind
	AggregateKindTest
}

type AggregateKind func(searchs interface{}) interface{}
type AggregateKindTest func(searchs interface{}) interface{}

func (a AggregateKind) Pipe(collection string, result interface{}, nestedObject interface{}) error {
	search := a(nestedObject)
	searchs := []interface{}{search}

	err := MongoDB.UseC(collection).Pipe(searchs).All(result)
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func unwind(nestedObject interface{}) interface{} {
	return map[string]interface{}{
		"$unwind": "$" + nestedObject.(string),
	}
}

func group(i interface{}) interface{} {
	return i
}

func MakeGroup(groupFields []string, sumFields []string, firstFields []string) interface{} {
	m := map[string]interface{}{}

	for _, f := range groupFields {
		m[f] = f
	}

	for _, s := range sumFields {
		m[s] = bson.M{"$sum": s}
	}

	for _, f := range firstFields {
		m[f] = bson.M{"$first": f}
	}

	return map[string]interface{}{
		"$group": m,
	}
}

func Do() {
	c := model.C.Workorder
	var r []interface{}
	A := AggregateKind(unwind)

	err := A.Pipe(c, &r, "WorkOrderList")
	if err != nil {
		panic(err)
	}
	util.PrintJson(r)

	mg := MakeGroup(
		[]string{"$WorkOrderId", "$WorkOrderList.StationName"},
		[]string{"$WorkOrderId", "$WorkOrderList.StationName"},
		[]string{"$WorkOrderList.CompletedQty", "$WorkOrderList.NonGoodQty"},
	)

	// #dont know why nil pointer
	// B := AggregateBoss{}
	// B.AggregateKind(unwind)

	B := AggregateKind(group)
	err = B.Pipe(c, &r, mg)
	if err != nil {
		panic(err)
	}
	util.PrintJson(r)

}
