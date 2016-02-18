/*
 * m2factor
 *
 * This file was automatically generated for 2Factor by APIMATIC BETA v2.0 on 02/18/2016
 */
package addonservices

import(
    "encoding/json"
    "github.com/apimatic/unirest-go"
    "m2factor"
    "m2factor/apihelper"
    "m2factor/models"
)
/*
 * Client structure as interface implementation
 */
type ADDONSERVICES_IMPL struct { }

/**
 * Check Balance For Addon Services
 * @param    string        apiKey           parameter: Required
 * @param    string        serviceName      parameter: Required
 * @return	Returns the *models.CheckBalanceAddonServicesModel response from the API call
 */
func (me *ADDONSERVICES_IMPL) GetCheckBalanceAddonServices (
            apiKey string,
            serviceName string) (*models.CheckBalanceAddonServicesModel, error) {
    //the base uri for api requests
    queryBuilder := m2factor.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/V1/{api_key}/ADDON_SERVICES/BAL/{service_name}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "api_key" : apiKey,
        "service_name" : serviceName,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "2Factor",
        "accept" : "application/json",
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.CheckBalanceAddonServicesModel = &models.CheckBalanceAddonServicesModel{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Pull Delivery Report - Transactional SMS
 * @param    string        apiKey         parameter: Required
 * @param    string        sessionId      parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *ADDONSERVICES_IMPL) GetPullDeliveryReport (
            apiKey string,
            sessionId string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := m2factor.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/V1/{api_key}/ADDON_SERVICES/RPT/TSMS/{session_id}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "api_key" : apiKey,
        "session_id" : sessionId,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "2Factor",
        "accept" : "application/json",
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Send Single / Bulk Transactional Messages / Schedule SMS
 * @param    string         apiKey      parameter: Required
 * @param    string         from        parameter: Required
 * @param    string         msg         parameter: Required
 * @param    string         to          parameter: Required
 * @param    *string        sendAt      parameter: Optional
 * @param    fieldParameters    Additional optional form parameters are supported by this endpoint
 * @return	Returns the *models.SendTransactionalSmsModel response from the API call
 */
func (me *ADDONSERVICES_IMPL) CreateSendTransactionalSMS (
            apiKey string,
            from string,
            msg string,
            to string,
            sendAt *string,
            fieldParameters map[string]interface{}) (*models.SendTransactionalSmsModel, error) {
    //the base uri for api requests
    queryBuilder := m2factor.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/V1/{api_key}/ADDON_SERVICES/SEND/TSMS"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "api_key" : apiKey,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "2Factor",
        "accept" : "application/json",
    }

    //form parameters
    parameters := map[string]interface{} {

        "From" : from,
        "Msg" : msg,
        "To" : to,
        "SendAt" : apihelper.ToString(*sendAt, "2019-01-01 00:00:01"),

    }

    //append optional form parameters
    if fieldParameters != nil {
        for k, v := range fieldParameters {
            parameters[k] = v
        }
    }

    //prepare API request
    request := unirest.Post(queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal *models.SendTransactionalSmsModel = &models.SendTransactionalSmsModel{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

