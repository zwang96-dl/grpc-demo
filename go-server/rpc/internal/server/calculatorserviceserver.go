// Code generated by goctl. DO NOT EDIT!
// Source: calculator.proto

package server

import (
	"context"

	"go-server/rpc/calculator"
	"go-server/rpc/internal/logic"
	"go-server/rpc/internal/svc"
)

type CalculatorServiceServer struct {
	svcCtx *svc.ServiceContext
}

func NewCalculatorServiceServer(svcCtx *svc.ServiceContext) *CalculatorServiceServer {
	return &CalculatorServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CalculatorServiceServer) Add(ctx context.Context, in *calculator.AddRequest) (*calculator.AddResponse, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}
