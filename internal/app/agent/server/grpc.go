package server

import (
	"context"

	"github.com/dollarkillerx/go-proxy-manager/proto/agent"
	"github.com/dollarkillerx/go-proxy-manager/proto/common"
)

func (s *Server) AddTask(ctx context.Context, req *common.TaskInfo) (*common.Empty, error) {
	s.cache.AddOrUpdateTask(req)
	return &common.Empty{}, nil
}

func (s *Server) ModifyTask(ctx context.Context, req *common.TaskInfo) (*common.Empty, error) {
	s.cache.AddOrUpdateTask(req)
	return &common.Empty{}, nil
}

func (s *Server) ApplyCertificate(ctx context.Context, req *agent.ApplyCertificateReq) (*common.Certificate, error) {
	return &common.Certificate{}, nil
}
