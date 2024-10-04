package nmodule

import (
	"errors"
	"google.golang.org/grpc"
	"strings"
)

const MaxMessageSize = 50

func ExtractRPCErrorMessage(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "desc = ") {
		parts := strings.Split(err.Error(), "desc = ")
		if len(parts) == 2 {
			return errors.New(strings.TrimSpace(parts[1]))
		}
	}
	return err
}

func DefaultGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	opts = append(opts,
		grpc.MaxRecvMsgSize(MaxMessageSize*1048*1048),
		grpc.MaxSendMsgSize(MaxMessageSize*1048*1048))
	return grpc.NewServer(opts...)
}
