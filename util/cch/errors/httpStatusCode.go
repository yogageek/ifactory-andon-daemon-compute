package errors

var StatusCode = map[int]int{
	Invalid_Auth_Token:         401,
	Response_Parse_Error:       400,
	Service_Not_Found:          404,
	Enable_Service_Plan_Failed: 400,
	Enable_Service_Failed:      400,
	Disable_Service_Failed:     400,
	Http_Request_Failed:        400,
	Invalid_Request:            400,
	Json_Parser_Error:          400,
	Plan_ID_Not_Found:          404,
	Create_Instance_Failed:     400,
	Resume_Org_Failed:          400,
	Canceled_Org_Failed:        400,
	Json_Validate_Error:        400,
	DB_Unavailable:             503,
	MARKETPLACE_RESPONSE_ERROR: 400,
}

func GetHttpCode(code int) int {
	httpCode, ok := StatusCode[code]
	if ok {
		return httpCode
	}

	if code > 200 && code < 600 {
		return code
	}

	return 400
}
