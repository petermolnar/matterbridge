// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AppRoleAssignmentRequestBuilder is request builder for AppRoleAssignment
type AppRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppRoleAssignmentRequest
func (b *AppRoleAssignmentRequestBuilder) Request() *AppRoleAssignmentRequest {
	return &AppRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppRoleAssignmentRequest is request for AppRoleAssignment
type AppRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for AppRoleAssignment
func (r *AppRoleAssignmentRequest) Get(ctx context.Context) (resObj *AppRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppRoleAssignment
func (r *AppRoleAssignmentRequest) Update(ctx context.Context, reqObj *AppRoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppRoleAssignment
func (r *AppRoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}