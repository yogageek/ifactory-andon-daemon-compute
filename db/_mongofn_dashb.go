package db

import (
	"github.com/thoas/go-funk"
	"gopkg.in/mgo.v2/bson"
)

//查詢欄位名稱(mongodb有一筆資料專門放要顯示的columns(是string array))
func QueryColumnNames(collection string, key string) (sortedColumnNames []string) {
	search1 := bson.M{"$match": bson.M{"type": "columnname"}}
	resp := MatchBySearchOptions(collection, search1)
	sortedColumnNamesIface := resp[0][key] //只有一筆,但有多個值(是一個array)
	funk.ForEach(sortedColumnNamesIface, func(x interface{}) {
		sortedColumnNames = append(sortedColumnNames, x.(string))
	})
	return sortedColumnNames
}

func QueryData(collection string) []bson.M {
	search1 := bson.M{"$match": bson.M{"type": "Data"}}
	return MatchBySearchOptions(collection, search1)
}

func QueryStatistics(collection string) []bson.M {
	search1 := bson.M{"$match": bson.M{"type": "Statistics"}}
	return MatchBySearchOptions(collection, search1)
}

func QueryDataByTypeId(collection string, column string, nodeId int) []bson.M {
	search1 := bson.M{"$match": bson.M{"type": "Data"}}
	search2 := bson.M{"$match": bson.M{column: nodeId}}
	return MatchBySearchOptions(collection, search1, search2)
}

func QueryStatisticsByTypeId(collection string, column string, nodeId int) []bson.M {
	search1 := bson.M{"$match": bson.M{"type": "Statistics"}}
	search2 := bson.M{"$match": bson.M{column: nodeId}}
	return MatchBySearchOptions(collection, search1, search2)
}

func QueryDeviceids(collection string) (ids []int) {
	resp := QueryData(collection)
	for _, d := range resp {
		// id := d["groupid"].(string)
		id := d["deviceid"].(int)
		ids = append(ids, id)
	}
	return ids
}
