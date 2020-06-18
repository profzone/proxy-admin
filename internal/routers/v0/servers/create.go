package servers

import (
	"context"
	"fmt"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
	"longhorn/proxy-admin/pkg/http"
)

func init() {
	Router.Register(courier.NewRouter(CreateServer{}))
}

// 创建服务器
type CreateServer struct {
	httpx.MethodPost
	Body models.ServerRequestBody `in:"body" name:"body"`
}

func (req CreateServer) Path() string {
	return ""
}

func (req CreateServer) Output(ctx context.Context) (result interface{}, err error) {
	server := req.Body.ToServer()
	if server == nil {
		return nil, fmt.Errorf("unsupport server type %v", req.Body.ServerType)
	}
	id, err := models.CreateServer(server, storage.Database)
	if err != nil {
		return
	}

	result = &http.IDResponse{ID: id}
	return
}
