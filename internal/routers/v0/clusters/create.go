package clusters

import (
	"context"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
	"longhorn/proxy-admin/pkg/http"
)

func init() {
	Router.Register(courier.NewRouter(CreateCluster{}))
}

// 创建集群
type CreateCluster struct {
	httpx.MethodPost
	Body models.Cluster `name:"body" in:"body"`
}

func (req CreateCluster) Path() string {
	return ""
}

func (req CreateCluster) Output(ctx context.Context) (result interface{}, err error) {
	id, err := models.CreateCluster(&req.Body, storage.Database)
	if err != nil {
		return
	}

	result = &http.IDResponse{ID: id}
	return
}
