package apis

import (
	"context"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
	"longhorn/proxy-admin/pkg"
	"longhorn/proxy-admin/pkg/http"
)

func init() {
	Router.Register(courier.NewRouter(CreateApi{}))
}

// 创建API
type CreateApi struct {
	httpx.MethodPost
	Body models.API `name:"body" in:"body"`
}

func (req CreateApi) Path() string {
	return ""
}

func (req CreateApi) Output(ctx context.Context) (result interface{}, err error) {
	id, err := pkg.Generator.GenerateUniqueID()
	if err != nil {
		return
	}

	req.Body.ID = id
	id, err = models.CreateAPI(&req.Body, storage.Database)
	if err != nil {
		return
	}

	result = http.IDResponse{
		ID: id,
	}

	return
}
