package model

type collection struct {
	Workorder      string
	Workorder_list string
}

var C collection

func init() {
	C = collection{
		Workorder:      "iii.sfc.workorder",
		Workorder_list: "iii.sfc.workorder_list",
	}
}
