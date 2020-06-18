package servers

import (
	"context"
	"fmt"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
)

func init() {
	Router.Register(courier.NewRouter(UpdateServer{}))
}

// 更新服务器
type UpdateServer struct {
	httpx.MethodPatch
	// 编号
	ID   uint64                   `name:"id,string" in:"path"`
	Body models.ServerRequestBody `name:"body" in:"body"`
}

func (req UpdateServer) Path() string {
	return ""
}

func (req UpdateServer) Output(ctx context.Context) (result interface{}, err error) {
	server := req.Body.ToServer()
	if server == nil {
		return nil, fmt.Errorf("unsupport server type %v", req.Body.ServerType)
	}
	server.SetIdentity(req.ID)
	err = models.UpdateServer(server, storage.Database)
	return
}
