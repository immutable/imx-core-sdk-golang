package imx

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
GetProject Gets a project detail

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param projectID Project ID
@return Balance
*/
func (c *Client) GetProject(ctx context.Context, l1signer L1Signer, projectID string) (*api.Project, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	response, httpResponse, err := c.projectsAPI.GetProject(ctx, projectID).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the details of a project: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

/*
GetProjects Gets projects info

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@return GetProjectsResponse
*/
func (c *Client) GetProjects(
	ctx context.Context,
	l1signer L1Signer,
	pageSize *int32,
	cursor, orderBy, direction *string,
) (*api.GetProjectsResponse, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	getProjectsRequest := c.projectsAPI.GetProjects(ctx).
		IMXTimestamp(timestamp).
		IMXSignature(signature)

	if pageSize != nil {
		getProjectsRequest.PageSize(*pageSize)
	}
	if cursor != nil {
		getProjectsRequest.Cursor(*cursor)
	}
	if orderBy != nil {
		getProjectsRequest.OrderBy(*orderBy)
	}
	if direction != nil {
		getProjectsRequest.Direction(*direction)
	}

	response, httpResponse, err := getProjectsRequest.Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the projects info: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

/*
CreateProject Creates a new project

In order to create a collection of NFTs, you must first register a project as the creator of the collection.
A project is an administrative level entity that is associated with an owner address, i.e. the address of the Ethereum wallet used to register a user account.
Only the project owner will be authorized to perform administrative tasks such as creating and updating collections and metadata schema.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param projectName Project name
@param companyName Company name to whom this project belongs to.
@param contactEmail Contact email for this project.
@return CreateProjectResponse
*/
func (c *Client) CreateProject(
	ctx context.Context,
	l1signer L1Signer,
	projectName, companyName, contactEmail string,
) (*api.CreateProjectResponse, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	createProjectRequest := api.NewCreateProjectRequest(companyName, contactEmail, projectName)
	createProjectResponse, httpResp, err := c.projectsAPI.CreateProject(ctx).
		CreateProjectRequest(*createProjectRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `CreateProject`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return createProjectResponse, nil
}
