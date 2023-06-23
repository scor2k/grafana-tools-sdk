package sdk

type CreatedBy struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type UpdatedBy struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type Meta struct {
	FolderName          string    `json:"folderName"`
	FolderUid           string    `json:"folderUid"`
	ConnectedDashboards int       `json:"connectedDashboards"`
	Created             string    `json:"created"`
	Updated             string    `json:"updated"`
	CreatedBy           CreatedBy `json:"createdBy"`
	UpdatedBy           UpdatedBy `json:"updatedBy"`
}

// https://grafana.com/docs/grafana/latest/developers/http_api/folder/
type LibraryElement struct {
	ID          int    `json:"id"`
	OrgId       uint   `json:"orgId"`
	FolderID    int    `json:"folderId"`
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Kind        int    `json:"kind"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Model       string `json:"model"` //TODO verificare cosa ritorna, doc non chiara
	Version     int    `json:"version"`
	Meta        Meta   `json:"meta"`
}
