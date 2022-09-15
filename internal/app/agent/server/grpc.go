package server

import (
	"context"

	"github.com/dollarkillerx/go-proxy-manager/proto/agent"
	"github.com/dollarkillerx/go-proxy-manager/proto/common"
)

func (s *Server) Info(ctx context.Context, req *agent.AgentInfoReq) (*agent.AgentInfoResp, error) {
	return &agent.AgentInfoResp{}, nil
}

func (s *Server) AddTask(ctx context.Context, req *agent.AddTaskReq) (*common.Empty, error) {
	s.cache.AddOrUpdateTask(req)
	return &common.Empty{}, nil
}

func (s *Server) ModifyTask(ctx context.Context, req *agent.ModifyTaskReq) (*common.Empty, error) {
	return &common.Empty{}, nil
}

func (s *Server) ApplyCertificate(ctx context.Context, req *agent.ApplyCertificateReq) (*common.Certificate, error) {
	return &common.Certificate{}, nil
}
