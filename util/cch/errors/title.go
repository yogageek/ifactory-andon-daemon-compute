package errors
var Title = map[int]string {
    Invalid_Auth_Token : "Invalid_Auth_Token",
    Response_Parse_Error : "Response_Parse_Error",
    Service_Not_Found : "Service_Not_Found",
    Enable_Service_Plan_Failed : "Enable_Service_Plan_Failed",
    Enable_Service_Failed : "Enable_Service_Failed",
    Disable_Service_Failed : "Disable_Service_Failed",
    Http_Request_Failed : "Http_Request_Failed",
    Invalid_Request : "Invalid_Request",
    Json_Parser_Error : "Json_Parser_Error",
    Plan_ID_Not_Found : "Plan_ID_Not_Found",
    Create_Instance_Failed : "Create_Instance_Failed",
    Resume_Org_Failed : "Resume_Org_Failed",
    Canceled_Org_Failed : "Canceled_Org_Failed",
    Json_Validate_Error : "Json_Validate_Error",
    DB_Unavailable : "DB_Unavailable",
    MARKETPLACE_RESPONSE_ERROR : "MARKETPLACE_RESPONSE_ERROR",
}
func GetName(code int) string {
    msg, ok := Title[code]
    if ok {
        return msg
    }
    return ""
}
