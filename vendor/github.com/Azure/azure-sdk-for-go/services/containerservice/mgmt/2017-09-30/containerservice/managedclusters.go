package containerservice

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ManagedClustersClient is the the Container Service Client.
type ManagedClustersClient struct {
	BaseClient
}

// NewManagedClustersClient creates an instance of the ManagedClustersClient client.
func NewManagedClustersClient(subscriptionID string) ManagedClustersClient {
	return NewManagedClustersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedClustersClientWithBaseURI creates an instance of the ManagedClustersClient client.
func NewManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) ManagedClustersClient {
	return ManagedClustersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a managed cluster with the specified configuration for agents and Kubernetes
// version.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the managed cluster resource.
// parameters is parameters supplied to the Create or Update a Managed Cluster operation.
func (client ManagedClustersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters ManagedCluster) (result ManagedClustersCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ManagedClusterProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ManagedClusterProperties.LinuxProfile", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ManagedClusterProperties.LinuxProfile.AdminUsername", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ManagedClusterProperties.LinuxProfile.AdminUsername", Name: validation.Pattern, Rule: `^[a-z][a-z0-9_-]*$`, Chain: nil}}},
						{Target: "parameters.ManagedClusterProperties.LinuxProfile.SSH", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ManagedClusterProperties.LinuxProfile.SSH.PublicKeys", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
					{Target: "parameters.ManagedClusterProperties.ServicePrincipalProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ManagedClusterProperties.ServicePrincipalProfile.ClientID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ManagedClusterProperties.ServicePrincipalProfile.KeyVaultSecretRef", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ManagedClusterProperties.ServicePrincipalProfile.KeyVaultSecretRef.VaultID", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ManagedClusterProperties.ServicePrincipalProfile.KeyVaultSecretRef.SecretName", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("containerservice.ManagedClustersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedClustersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters ManagedCluster) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) CreateOrUpdateSender(req *http.Request) (future ManagedClustersCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedCluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the managed cluster with a specified resource group and name.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the managed cluster resource.
func (client ManagedClustersClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result ManagedClustersDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagedClustersClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) DeleteSender(req *http.Request) (future ManagedClustersDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the managed cluster with a specified resource group and name.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the managed cluster resource.
func (client ManagedClustersClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result ManagedCluster, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedClustersClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) GetResponder(resp *http.Response) (result ManagedCluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAccessProfiles gets the accessProfile for the specified role name of the managed cluster with a specified
// resource group and name.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the managed cluster resource.
// roleName is the name of the role for managed cluster accessProfile resource.
func (client ManagedClustersClient) GetAccessProfiles(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (result ManagedClusterAccessProfile, err error) {
	req, err := client.GetAccessProfilesPreparer(ctx, resourceGroupName, resourceName, roleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "GetAccessProfiles", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAccessProfilesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "GetAccessProfiles", resp, "Failure sending request")
		return
	}

	result, err = client.GetAccessProfilesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "GetAccessProfiles", resp, "Failure responding to request")
	}

	return
}

// GetAccessProfilesPreparer prepares the GetAccessProfiles request.
func (client ManagedClustersClient) GetAccessProfilesPreparer(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/accessProfiles/{roleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAccessProfilesSender sends the GetAccessProfiles request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) GetAccessProfilesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetAccessProfilesResponder handles the response to the GetAccessProfiles request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) GetAccessProfilesResponder(resp *http.Response) (result ManagedClusterAccessProfile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUpgradeProfile gets the details of the upgrade profile for a managed cluster with a specified resource group and
// name.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the managed cluster resource.
func (client ManagedClustersClient) GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string) (result ManagedClusterUpgradeProfile, err error) {
	req, err := client.GetUpgradeProfilePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "GetUpgradeProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUpgradeProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "GetUpgradeProfile", resp, "Failure sending request")
		return
	}

	result, err = client.GetUpgradeProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "GetUpgradeProfile", resp, "Failure responding to request")
	}

	return
}

// GetUpgradeProfilePreparer prepares the GetUpgradeProfile request.
func (client ManagedClustersClient) GetUpgradeProfilePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeProfiles/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUpgradeProfileSender sends the GetUpgradeProfile request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) GetUpgradeProfileSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetUpgradeProfileResponder handles the response to the GetUpgradeProfile request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) GetUpgradeProfileResponder(resp *http.Response) (result ManagedClusterUpgradeProfile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of managed clusters in the specified subscription. The operation returns properties of each managed
// cluster.
func (client ManagedClustersClient) List(ctx context.Context) (result ManagedClusterListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "List", resp, "Failure sending request")
		return
	}

	result.mclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagedClustersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) ListResponder(resp *http.Response) (result ManagedClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ManagedClustersClient) listNextResults(lastResults ManagedClusterListResult) (result ManagedClusterListResult, err error) {
	req, err := lastResults.managedClusterListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedClustersClient) ListComplete(ctx context.Context) (result ManagedClusterListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists managed clusters in the specified subscription and resource group. The operation returns
// properties of each managed cluster.
//
// resourceGroupName is the name of the resource group.
func (client ManagedClustersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ManagedClusterListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.mclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.mclr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ManagedClustersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedClustersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ManagedClustersClient) ListByResourceGroupResponder(resp *http.Response) (result ManagedClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ManagedClustersClient) listByResourceGroupNextResults(lastResults ManagedClusterListResult) (result ManagedClusterListResult, err error) {
	req, err := lastResults.managedClusterListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ManagedClustersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedClustersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ManagedClusterListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}
