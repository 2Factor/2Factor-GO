/*
 * m2factor
 *
 * This file was automatically generated for 2Factor by APIMATIC BETA v2.0 on 02/18/2016
 */

package addonservices


import "m2factor/models"

/*
 * Interface for the ADDONSERVICES_IMPL
 */
type ADDONSERVICES interface {
    GetCheckBalanceAddonServices (string, string) (*models.CheckBalanceAddonServicesModel, error)

    GetPullDeliveryReport (string, string) (interface{}, error)

    CreateSendTransactionalSMS (string, string, string, string, *string, map[string]interface{}) (*models.SendTransactionalSmsModel, error)
}

/*
 * Factory for the ADDONSERVICES interaface returning ADDONSERVICES_IMPL
 */
func NewADDONSERVICES() ADDONSERVICES {
    return &ADDONSERVICES_IMPL{}
}