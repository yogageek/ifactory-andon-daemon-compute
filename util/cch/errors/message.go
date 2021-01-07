package errors
var Msg = map[int]string {
    Invalid_Auth_Token : "Invalid Auth Token",
    Response_Parse_Error : "Response invalid due to parse error.",
    Service_Not_Found : "Service not found.",
    Enable_Service_Plan_Failed : "Enable service plan failed.",
    Enable_Service_Failed : "Enable service failed.",
    Disable_Service_Failed : "Disable service failed.",
    Http_Request_Failed : "Send http request failed.",
    Invalid_Request : "The request is invalid",
    Json_Parser_Error : "The Json format is invalid",
    Plan_ID_Not_Found : "The plan id not found",
    Create_Instance_Failed : "Create Instance Failed",
    Resume_Org_Failed : "Resume org failed",
    Canceled_Org_Failed : "Canceled org Failed",
    Json_Validate_Error : "Json_Validate_Error",
    DB_Unavailable : "DB is unavailable",
    MARKETPLACE_RESPONSE_ERROR : "Marketplace response error",
}
func GetMsg(code int) string {
    msg, ok := Msg[code]
    if ok {
        return msg
    }
    return ""
}
