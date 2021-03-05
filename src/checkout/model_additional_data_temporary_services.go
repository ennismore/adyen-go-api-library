/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 67
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// AdditionalDataTemporaryServices struct for AdditionalDataTemporaryServices
type AdditionalDataTemporaryServices struct {
	// Customer code, if supplied by a customer. * Encoding: ASCII * maxLength: 25
	EnhancedSchemeDataCustomerReference string `json:"enhancedSchemeData.customerReference,omitempty"`
	// Name or ID associated with the individual working in a temporary capacity. * maxLength: 40
	EnhancedSchemeDataEmployeeName string `json:"enhancedSchemeData.employeeName,omitempty"`
	// Description of the job or task of the individual working in a temporary capacity. * maxLength: 40
	EnhancedSchemeDataJobDescription string `json:"enhancedSchemeData.jobDescription,omitempty"`
	// Amount paid per regular hours worked, minor units. * maxLength: 7
	EnhancedSchemeDataRegularHoursRate string `json:"enhancedSchemeData.regularHoursRate,omitempty"`
	// Amount of time worked during a normal operation for the task or job. * maxLength: 7
	EnhancedSchemeDataRegularHoursWorked string `json:"enhancedSchemeData.regularHoursWorked,omitempty"`
	// Name of the individual requesting temporary services. * maxLength: 40
	EnhancedSchemeDataRequestName string `json:"enhancedSchemeData.requestName,omitempty"`
	// Date for the beginning of the pay period. * Format: ddMMyy * maxLength: 6
	EnhancedSchemeDataTempStartDate string `json:"enhancedSchemeData.tempStartDate,omitempty"`
	// Date of the end of the billing cycle. * Format: ddMMyy * maxLength: 6
	EnhancedSchemeDataTempWeekEnding string `json:"enhancedSchemeData.tempWeekEnding,omitempty"`
	// Total tax amount, in minor units. For example, 2000 means USD 20.00 * maxLength: 12
	EnhancedSchemeDataTotalTaxAmount string `json:"enhancedSchemeData.totalTaxAmount,omitempty"`
}
