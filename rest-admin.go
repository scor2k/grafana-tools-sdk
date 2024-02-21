package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enuan/grafana-tools-sdk/openapi"
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
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
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

// DisableUser disables a global user
// Requires basic authentication and that the authenticated user ia Grafana Admin
// Reflects POST api/admin/users/{user_id}/disable API call.
func (r *Client) DisableUser(ctx context.Context, uid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.post(ctx, fmt.Sprintf("api/admin/users/%d/disable", uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// EnableUser enables a global user
// Requires basic authentication and that the authenticated user ia Grafana Admin
// Reflects POST api/admin/users/{user_id}/disable API call.
func (r *Client) EnableUser(ctx context.Context, uid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.post(ctx, fmt.Sprintf("api/admin/users/%d/enable", uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// UpdateUserPermissions updates the permissions of a global user.
// Requires basic authentication and that the authenticated user is a Grafana Admin.
// Reflects PUT /api/admin/users/:userId/permissions API call.
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
	err = json.Unmarshal(raw, &reply)
	return reply, err
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
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// UpdateUserPassword updates the password of a global user.
// Requires basic authentication and that the authenticated user is a Grafana Admin.
// Reflects PUT /api/admin/users/:userId/password API call.
func (r *Client) UpdateUserPassword(ctx context.Context, password UserPassword, uid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = json.Marshal(password); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put(ctx, fmt.Sprintf("api/admin/users/%d/password", uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

func (r *Client) GetUserAuthToken(ctx context.Context, id uint) ([]openapi.UserToken, error) {
	var (
		raw   []byte
		reply []openapi.UserToken
		err   error
	)
	if raw, _, err = r.get(ctx, fmt.Sprintf("api/admin/users/%d/auth-tokens", id), nil); err != nil {
		return reply, err
	}
	fmt.Println("auth-tokens: ", string(raw))
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

func (r *Client) RevokeAuthToken(ctx context.Context, id uint, authTokenId int64) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	token := openapi.RevokeAuthTokenCmd{AuthTokenId: &authTokenId}
	if raw, err = json.Marshal(token); err != nil {
		return reply, err
	}

	if raw, _, err = r.post(ctx, fmt.Sprintf("api/admin/users/%d/revoke-auth-token", id), nil, raw); err != nil {
		return reply, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}
