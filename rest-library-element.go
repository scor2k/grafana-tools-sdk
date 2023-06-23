package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Element struct {
	ID          uint   `json:"id"`
	OrgId       uint   `json:"orgId"`
	FolderId    uint   `json:"folderId"`
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Kind        uint   `json:"kind"`
	Type        string `json:"type"`
	Description string `json:"description"`
	//Model       string `json:"model"` //TODO verificare cosa ritorna, doc non chiara
	Version uint `json:"version"`
	Meta    Meta `json:"meta"`
}

type Result struct {
	TotalCount uint      `json:"totalCount"`
	Page       uint      `json:"page"`
	PerPage    uint      `json:"perPage"`
	Elements   []Element `json:"elements"`
}

// FoundBoard keeps result of search with metadata of a dashboard.
type FoundLibraryElement struct {
	Result Result `json:"result"`
}

// SearchLibraryElements gets all library element.
// Reflects GET /api/library-elements API call.
func (r *Client) SearchLibraryElements(ctx context.Context) (FoundLibraryElement, error) {
	var (
		raw      []byte
		libElems FoundLibraryElement
		code     int
		err      error
	)
	u := url.URL{}
	q := u.Query()
	q.Set("perpage", strconv.FormatUint(uint64(1000), 10))
	if raw, code, err = r.get(ctx, "api/library-elements", q); err != nil {
		return libElems, err
	}
	if code != 200 {
		return libElems, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &libElems)
	return libElems, err
}

// GetLibraryElementByUID gets library element by uid.
// Reflects GET /api/library-elements/:uid API call.
func (r *Client) GetLibraryElementByUID(ctx context.Context, UID string) (LibraryElement, error) {
	var (
		raw     []byte
		libElem LibraryElement
		err     error
	)
	/*if raw, code, err = r.get(ctx, fmt.Sprintf("/api/library-elements/%s", UID), nil); err != nil {
		return libElem, err
	}
	if code != 200 {
		return libElem, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}*/
	if raw, err = r.GetRawLibraryElementByUID(ctx, UID); err != nil {
		return libElem, err
	}
	err = json.Unmarshal(raw, &libElem)
	return libElem, err
}

func (r *Client) GetRawLibraryElementByUID(ctx context.Context, UID string) ([]byte, error) {
	var (
		raw  []byte
		code int
		err  error
	)
	if raw, code, err = r.get(ctx, fmt.Sprintf("/api/library-elements/%s", UID), nil); err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	return raw, err
}

// GetLibraryElementByName gets library element by uid.
// Reflects GET /api/library-elements/name/:uid API call.
func (r *Client) GetLibraryElementByName(ctx context.Context, Name string) (LibraryElement, error) {
	var (
		raw     []byte
		libElem LibraryElement
		code    int
		err     error
	)
	if raw, code, err = r.get(ctx, fmt.Sprintf("/api/library-elements/name/%s", Name), nil); err != nil {
		return libElem, err
	}
	if code != 200 {
		return libElem, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	//fmt.Printf("raw: >>>\n%s\n<<<\n", string(raw[:]))
	err = json.Unmarshal(raw, &libElem)
	return libElem, err
}

func (r *Client) GetRawLibraryElementByName(ctx context.Context, name string) ([]byte, error) {
	var (
		raw  []byte
		code int
		err  error
	)
	if raw, code, err = r.get(ctx, fmt.Sprintf("/api/library-elements/name/%s", name), nil); err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	return raw, err
}

// UpdateFolderByUID update folder by uid
// Reflects PATCH /api/library-elements/:uid API call.
func (r *Client) UpdateRawLibraryElementByUID(ctx context.Context, uid string, rawLibElem []byte) error {
	var (
		rf   Folder
		code int
		err  error
	)
	raw, code, err := r.patch(ctx, fmt.Sprintf("api/library-elements/%s", uid), nil, rawLibElem)
	if err != nil {
		return err
	}
	if code != 200 {
		return fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &rf)
	return err
}
