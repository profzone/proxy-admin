package clusters

import (
	"context"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
)

func init() {
	Router.Register(courier.NewRouter(DeleteCluster{}))
}

// 删除集群
type DeleteCluster struct {
	httpx.MethodDelete
	// 编号
	ID uint64 `name:"id,string" in:"path"`
}

func (req DeleteCluster) Path() string {
	return "/:id"
}

func (req DeleteCluster) Output(ctx context.Context) (result interface{}, err error) {
	err = models.DeleteCluster(req.ID, storage.Database)
	return
}
