package model

type collection struct {
	Workorder      string
	Workorder_list string //#已棄用
}

var C collection

func init() {
	C = collection{
		Workorder:      "iii.sfc.workorder",
		Workorder_list: "iii.sfc.workorder_list",
	}
}
