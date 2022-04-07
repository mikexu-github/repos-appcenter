package exec

import (
	"context"
	"fmt"
	"net/http"

	"github.com/quanxiang-cloud/appcenter/pkg/chaos/define"
	"github.com/quanxiang-cloud/cabin/tailormade/client"
)

// Key
const (
	FormCreateRole = "form-role"
	FormAssignRole = "form-assign"

// FormHost   = "form"
// CreateRole = "/api/v1/form/%s/internal/apiRole/create"
// AssignRole = "/api/v1/form/%s/internal/apiRole/grant/assign/%s"
)

const (
	perInitTypes = 1
	name         = "全部权限"
	description  = "系统默认角色"
)

// FormExecutor FormExecutor
type FormExecutor struct {
	Client     http.Client
	CreateRole string
	AssignRole string
}

type createRoleReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type createRoleResp struct {
	RoleID string `json:"roleID"`
}

type assignReq struct {
	Add []*user `json:"add"`
}

type assignResp struct{}

type user struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
}

// Exec Exec
func (s *FormExecutor) Exec(ctx context.Context, m define.Msg) error {
	roleReq := &createRoleReq{
		Name:        name,
		Description: description,
	}
	roleResp := &createRoleResp{}
	if err := client.POST(ctx, &s.Client, fmt.Sprintf(s.CreateRole, m.AppID), roleReq, roleResp); err != nil {
		return err
	}

	assignReq := &assignReq{
		Add: []*user{
			{
				ID:   m.CreateBy,
				Type: perInitTypes,
			},
		},
	}
	assignResp := &assignResp{}
	if err := client.POST(ctx, &s.Client, fmt.Sprintf(s.AssignRole, m.AppID, roleResp.RoleID), assignReq, assignResp); err != nil {
		return err
	}

	return nil
}

// Bit Bit
func (*FormExecutor) Bit() int {
	return define.BitFormAPI
}
