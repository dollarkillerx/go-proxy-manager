package server

import (
	"context"

	"github.com/dollarkillerx/go-proxy-manager/proto/backend"
)

func (s *Server) Register(ctx context.Context, info *backend.RgInfo) (*backend.RgResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SyncData(ctx context.Context, req *backend.SyncDataReq) (*backend.SyncDataResp, error) {
	//TODO implement me
	panic("implement me")
}
