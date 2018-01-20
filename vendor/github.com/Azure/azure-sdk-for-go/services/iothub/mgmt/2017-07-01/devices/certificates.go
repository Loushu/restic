package devices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// CertificatesClient is the use this API to manage the IoT hubs in your Azure subscription.
type CertificatesClient struct {
	ManagementClient
}

// NewCertificatesClient creates an instance of the CertificatesClient client.
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return NewCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificatesClientWithBaseURI creates an instance of the CertificatesClient client.
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return CertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate adds new or replaces existing certificate.
//
// resourceGroupName is the name of the resource group that contains the IoT hub. resourceName is the name of the IoT
// hub. certificateName is the name of the certificate certificateDescription is the certificate body. ifMatch is eTag
// of the Certificate. Do not specify for creating a brand new certificate. Required to update an existing certificate.
func (client CertificatesClient) CreateOrUpdate(resourceGroupName string, resourceName string, certificateName string, certificateDescription CertificateBodyDescription, ifMatch string) (result CertificateDescription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "devices.CertificatesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, resourceName, certificateName, certificateDescription, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CertificatesClient) CreateOrUpdatePreparer(resourceGroupName string, resourceName string, certificateName string, certificateDescription CertificateBodyDescription, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(certificateDescription),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) CreateOrUpdateResponder(resp *http.Response) (result CertificateDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing X509 certificate or does nothing if it does not exist.
//
// resourceGroupName is the name of the resource group that contains the IoT hub. resourceName is the name of the IoT
// hub. certificateName is the name of the certificate ifMatch is eTag of the Certificate.
func (client CertificatesClient) Delete(resourceGroupName string, resourceName string, certificateName string, ifMatch string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "devices.CertificatesClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, resourceName, certificateName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CertificatesClient) DeletePreparer(resourceGroupName string, resourceName string, certificateName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CertificatesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateVerificationCode generates verification code for proof of possession flow. The verification code will be
// used to generate a leaf certificate.
//
// resourceGroupName is the name of the resource group that contains the IoT hub. resourceName is the name of the IoT
// hub. certificateName is the name of the certificate ifMatch is eTag of the Certificate.
func (client CertificatesClient) GenerateVerificationCode(resourceGroupName string, resourceName string, certificateName string, ifMatch string) (result CertificateWithNonceDescription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "devices.CertificatesClient", "GenerateVerificationCode")
	}

	req, err := client.GenerateVerificationCodePreparer(resourceGroupName, resourceName, certificateName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "GenerateVerificationCode", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateVerificationCodeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "GenerateVerificationCode", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateVerificationCodeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "GenerateVerificationCode", resp, "Failure responding to request")
	}

	return
}

// GenerateVerificationCodePreparer prepares the GenerateVerificationCode request.
func (client CertificatesClient) GenerateVerificationCodePreparer(resourceGroupName string, resourceName string, certificateName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}/generateVerificationCode", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// GenerateVerificationCodeSender sends the GenerateVerificationCode request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GenerateVerificationCodeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GenerateVerificationCodeResponder handles the response to the GenerateVerificationCode request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GenerateVerificationCodeResponder(resp *http.Response) (result CertificateWithNonceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns the certificate.
//
// resourceGroupName is the name of the resource group that contains the IoT hub. resourceName is the name of the IoT
// hub. certificateName is the name of the certificate
func (client CertificatesClient) Get(resourceGroupName string, resourceName string, certificateName string) (result CertificateDescription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "devices.CertificatesClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, resourceName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CertificatesClient) GetPreparer(resourceGroupName string, resourceName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetResponder(resp *http.Response) (result CertificateDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByIotHub returns the list of certificates.
//
// resourceGroupName is the name of the resource group that contains the IoT hub. resourceName is the name of the IoT
// hub.
func (client CertificatesClient) ListByIotHub(resourceGroupName string, resourceName string) (result CertificateListDescription, err error) {
	req, err := client.ListByIotHubPreparer(resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "ListByIotHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByIotHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "ListByIotHub", resp, "Failure sending request")
		return
	}

	result, err = client.ListByIotHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "ListByIotHub", resp, "Failure responding to request")
	}

	return
}

// ListByIotHubPreparer prepares the ListByIotHub request.
func (client CertificatesClient) ListByIotHubPreparer(resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByIotHubSender sends the ListByIotHub request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) ListByIotHubSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByIotHubResponder handles the response to the ListByIotHub request. The method always
// closes the http.Response Body.
func (client CertificatesClient) ListByIotHubResponder(resp *http.Response) (result CertificateListDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Verify verifies the certificate's private key possession by providing the leaf cert issued by the verifying pre
// uploaded certificate.
//
// resourceGroupName is the name of the resource group that contains the IoT hub. resourceName is the name of the IoT
// hub. certificateName is the name of the certificate certificateVerificationBody is the name of the certificate
// ifMatch is eTag of the Certificate.
func (client CertificatesClient) Verify(resourceGroupName string, resourceName string, certificateName string, certificateVerificationBody CertificateVerificationDescription, ifMatch string) (result CertificateDescription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "devices.CertificatesClient", "Verify")
	}

	req, err := client.VerifyPreparer(resourceGroupName, resourceName, certificateName, certificateVerificationBody, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Verify", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Verify", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.CertificatesClient", "Verify", resp, "Failure responding to request")
	}

	return
}

// VerifyPreparer prepares the Verify request.
func (client CertificatesClient) VerifyPreparer(resourceGroupName string, resourceName string, certificateName string, certificateVerificationBody CertificateVerificationDescription, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}/verify", pathParameters),
		autorest.WithJSON(certificateVerificationBody),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// VerifySender sends the Verify request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) VerifySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// VerifyResponder handles the response to the Verify request. The method always
// closes the http.Response Body.
func (client CertificatesClient) VerifyResponder(resp *http.Response) (result CertificateDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}