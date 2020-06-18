package servers

import (
	"context"
	"github.com/profzone/eden-framework/pkg/courier"
	"github.com/profzone/eden-framework/pkg/courier/httpx"
	"longhorn/proxy-admin/internal/models"
	"longhorn/proxy-admin/internal/storage"
)

func init() {
	Router.Register(courier.NewRouter(GetServers{}))
}

// 分页获取服务器
type GetServers struct {
	httpx.MethodGet
	// 分页偏移量
	Start uint64 `json:"start" name:"start" in:"query" default:"0" validate:"@uint64[0,]"`
	// 每页数据量
	Limit int64 `json:"limit" name:"limit" in:"query" default:"10" validate:"@int64[-1,100]"`
}

func (req GetServers) Path() string {
	return ""
}

type GetServersResult struct {
	NextID uint64                 `json:"nextID"`
	Data   []models.GeneralServer `json:"data"`
}

func (req GetServers) Output(ctx context.Context) (result interface{}, err error) {
	resp := &GetServersResult{
		NextID: 0,
		Data:   make([]models.GeneralServer, 0),
	}

	resp.NextID, err = models.WalkServers(req.Start, req.Limit, func(e storage.Element) error {
		resp.Data = append(resp.Data, *e.(*models.GeneralServer))
		return nil
	}, storage.Database)

	result = resp
	return
}
