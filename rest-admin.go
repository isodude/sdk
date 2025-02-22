package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

// CreateUser creates a new global user.
// Requires basic authentication and that the authenticated user is a Grafana Admin.
// Reflects POST /api/admin/users API call.
func (r *Client) CreateUser(ctx context.Context, user User) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(user); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(ctx, "api/admin/users", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	return resp, json.Unmarshal(raw, &resp)
}

// DeleteUser deletes a global user
// Requires basic authentication and that the authenticated user ia Grafana Admin
// Reflects DELETE /api/admin/users/:userId API call.
func (r *Client) DeleteUser(ctx context.Context, uid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.delete(ctx, fmt.Sprintf("api/admin/users/%d", uid)); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// UpdateUserPermissions updates the permissions of a global user.
// Requires basic authentication and that the authenticated user is a Grafana Admin.
// Reflects PUT /api/admin/users/:userId/password API call.
func (r *Client) UpdateUserPermissions(ctx context.Context, permissions UserPermissions, uid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = json.Marshal(permissions); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put(ctx, fmt.Sprintf("api/admin/users/%d/permissions", uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	return reply, json.Unmarshal(raw, &reply)
}

// SwitchUserContext switches user context to the given organization.
// Requires basic authentication and that the authenticated user is a Grafana Admin.
// Reflects POST /api/users/:userId/using/:organizationId API call.
func (r *Client) SwitchUserContext(ctx context.Context, uid uint, oid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)

	if raw, _, err = r.post(ctx, fmt.Sprintf("/api/users/%d/using/%d", uid, oid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	return resp, json.Unmarshal(raw, &resp)
}
