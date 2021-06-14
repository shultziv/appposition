package grpc

import (
	pb "github.com/shultziv/appposition/internal/delivery/grpc/proto"
	"github.com/shultziv/appposition/internal/service/appposition"
	"google.golang.org/grpc"
)

type Handler struct {
	pb.UnimplementedAppPositionServer
	appPositionService *appposition.AppPosition
}

func NewHandler(appPositionService *appposition.AppPosition) *Handler {
	return &Handler{
		appPositionService: appPositionService,
	}
}

func (h *Handler) RegisterServer(server *grpc.Server) {
	pb.RegisterAppPositionServer(server, h)
}
