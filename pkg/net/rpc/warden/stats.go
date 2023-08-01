package warden

import (
	"context"
	"strconv"

	nmd "git.huoys.com/middle-end/kratos/pkg/naming/metadata"
	"git.huoys.com/middle-end/kratos/pkg/stat/sys/cpu"

	"google.golang.org/grpc"
	gmd "google.golang.org/grpc/metadata"
)

func (s *Server) stats() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		var cpustat cpu.Stat
		cpu.ReadStat(&cpustat)
		if cpustat.Usage != 0 {
			trailer := gmd.Pairs([]string{nmd.CPUUsage, strconv.FormatInt(int64(cpustat.Usage), 10)}...)
			grpc.SetTrailer(ctx, trailer)
		}
		return
	}
}