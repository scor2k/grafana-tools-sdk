package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// SearchUsersWithPaging search users with paging.
// query optional.  query value is contained in one of the name, login or email fields. Query values with spaces need to be url encoded e.g. query=Jane%20Doe
// perpage optional. default 1000
// page optional. default 1
// http://docs.grafana.org/http_api/user/#search-users
// http://docs.grafana.org/http_api/user/#search-users-with-paging
//
// Reflects GET /api/users/search API call.
func (r *Client) SearchTeamWithName(ctx context.Context, name string) (PageTeams, error) {
	var (
		raw       []byte
		pageTeams PageTeams
		code      int
		err       error
	)
	v := url.Values{}
	v.Set("name", name)

	//qs := fmt.Sprintf("api/teams/search?name=%s", name)
	if raw, code, err = r.get(ctx, "api/teams/search", v); err != nil {
		return pageTeams, err
	}
	if code != 200 {
		return pageTeams, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&pageTeams); err != nil {
		return pageTeams, fmt.Errorf("unmarshal users: %s\n%s", err, raw)
	}
	return pageTeams, err
}
