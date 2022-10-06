package imx

import (
	"context"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

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
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	createProjectRequest := api.NewCreateProjectRequest(companyName, contactEmail, projectName)
	createProjectResponse, httpResponse, err := c.projectsAPI.CreateProject(ctx).
		CreateProjectRequest(*createProjectRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return createProjectResponse, nil
}

/*
GetProject Gets a project detail

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param id Project ID
@return Balance
*/
func (c *Client) GetProject(ctx context.Context, l1signer L1Signer, id string) (*api.Project, error) {
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	response, httpResponse, err := c.projectsAPI.GetProject(ctx, id).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}

/*
GetProjects Gets projects owned by given user

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param pageSize The page size of the result
@param cursor The cursor
@param orderBy The property to sort by
@param direction Direction to sort (asc/desc)
@return GetProjectsResponse
*/
func (c *Client) GetProjects(
	ctx context.Context,
	l1signer L1Signer,
	pageSize *int32,
	cursor, orderBy, direction *string,
) (*api.GetProjectsResponse, error) {
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
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
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}
