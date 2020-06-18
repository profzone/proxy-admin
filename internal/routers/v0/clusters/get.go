package clusters

import (
	"context"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
)

func init() {
	Router.Register(courier.NewRouter(GetCluster{}))
}

// 获取单个集群
type GetCluster struct {
	httpx.MethodGet

	ID uint64 `name:"id,string" in:"path"`
}

func (req GetCluster) Path() string {
	return "/:id"
}

func (req GetCluster) Output(ctx context.Context) (result interface{}, err error) {
	cluster, err := models.GetCluster(req.ID, storage.Database)
	if err != nil {
		return
	}

	result = cluster
	return
}
