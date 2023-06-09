/*
Copyright 2022 QuanxiangCloud Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resp

import "github.com/quanxiang-cloud/appcenter/internal/models"

// AdminAppCenter AdminAppCenter
type AdminAppCenter struct {
	ID          string                 `json:"id"`
	AppName     string                 `json:"appName,omitempty"`
	AccessURL   string                 `json:"accessURL,omitempty"`
	AppIcon     string                 `json:"appIcon,omitempty"`
	CreateBy    string                 `json:"createBy,omitempty"`
	UpdateBy    string                 `json:"updateBy,omitempty"`
	CreateTime  int64                  `json:"createTime,omitempty"`
	UpdateTime  int64                  `json:"updateTime,omitempty"`
	UseStatus   int                    `json:"useStatus,omitempty"` //published:1，unpublished:-1
	Server      int                    `json:"server,omitempty"`
	DelFlag     int64                  `json:"delFlag,omitempty"`
	AppSign     string                 `json:"appSign,omitempty"`
	Extension   map[string]interface{} `json:"extension"`
	Description string                 `json:"description"`
	PerPoly     bool                   `json:"perPoly"`
}

// UserAppCenter UserAppCenter
type UserAppCenter struct {
	ID          string                 `json:"id,omitempty"`
	AppName     string                 `json:"appName"`
	AccessURL   string                 `json:"accessURL"`
	AppIcon     string                 `json:"appIcon"`
	Extension   map[string]interface{} `json:"extension"`
	Description string                 `json:"description"`
}

// GetAppsByIDsResp GetAppsByIDsResp
type GetAppsByIDsResp struct {
	Apps []*UserAppCenter `json:"apps"`
}

// CheckIsAdminResp CheckIsAdminResp
type CheckIsAdminResp struct {
	IsAdmin bool `json:"isAdmin"`
}

// AddAppScopeResp AddAppScopeResp
type AddAppScopeResp struct {
}

// GetOneResp GetOneResp
type GetOneResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	DelFlag int64  `json:"delFlag"`
	PerPoly bool   `json:"perPoly"`
}

// ExportAppResp ExportAppResp
type ExportAppResp struct {
	AppID   string `json:"appID"`
	AppName string `json:"appName"`
	Version string `json:"version"`
}

// ImportAppResp ImportAppResp
type ImportAppResp struct {
}

// CheckAppAccessResp CheckAppAccessResp
type CheckAppAccessResp struct {
	IsAuthority bool `json:"isAuthority"`
}

// FinishImportResp FinishImportResp
type FinishImportResp struct {
}

// ErrorImportResp ErrorImportResp
type ErrorImportResp struct {
}

// CheckImportVersionResp CheckImportVersionResp
type CheckImportVersionResp struct {
}

// InitCallBackResp InitCallBackResp
type InitCallBackResp struct {
}

// InitServerResp InitServerResp
type InitServerResp struct {
}

type ChangePerPolyResp struct {
}

type HomeAccessListResp struct {
	List  []models.Scope `json:"list"`
	Total int64          `json:"total"`
}
