package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// codeToLevel redirects OK to DEBUG level logging instead of INFO
func codeToLevel(code codes.Code) logrus.Level {
	if code == codes.OK {
		// It is DEBUG
		return logrus.DebugLevel
	}
	return grpc_logrus.DefaultCodeToLevel(code)
}

// AddLogging returns grpc.Server config option that turn on logging.
func AddLogging(logger *logrus.Entry, opts []grpc.ServerOption) []grpc.ServerOption {
	o := []grpc_logrus.Option{
		grpc_logrus.WithLevels(codeToLevel),
	}
	grpc_logrus.ReplaceGrpcLogger(logger)

	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_logrus.UnaryServerInterceptor(logger, o...),
	))

	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_logrus.StreamServerInterceptor(logger, o...),
	))

	return opts
}
