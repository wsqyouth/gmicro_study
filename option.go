package gmicro

import (
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Option is service functional option
// See this post about the "functional options" pattern:
// http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type Option func(s *Service)

// WithRecovery service recover func.
func WithRecovery(f func()) Option {
	return func(s *Service) {
		s.recovery = f
	}
}

// WithHTTPHandler returns an Option to set the httpHandler
func WithHTTPHandler(h HTTPHandlerFunc) Option {
	return func(s *Service) {
		s.httpHandler = h
	}
}

// WithAnnotator returns an Option to append some annotator
func WithAnnotator(annotator ...AnnotatorFunc) Option {
	return func(s *Service) {
		s.annotators = append(s.annotators, annotator...)
	}
}

// WithErrorHandler returns an Option to set the errorHandler
func WithErrorHandler(errorHandler runtime.ProtoErrorHandlerFunc) Option {
	return func(s *Service) {
		s.errorHandler = errorHandler
	}
}

// WithUnaryInterceptor returns an Option to append some unaryInterceptor
func WithUnaryInterceptor(unaryInterceptor ...grpc.UnaryServerInterceptor) Option {
	return func(s *Service) {
		s.unaryInterceptors = append(s.unaryInterceptors, unaryInterceptor...)
	}
}

// WithStreamInterceptor returns an Option to append some streamInterceptor
func WithStreamInterceptor(streamInterceptor ...grpc.StreamServerInterceptor) Option {
	return func(s *Service) {
		s.streamInterceptors = append(s.streamInterceptors, streamInterceptor...)
	}
}

// WithShutdownFunc returns an Option to register a function which will be called when server shutdown
func WithShutdownFunc(f func()) Option {
	return func(s *Service) {
		s.shutdownFunc = f
	}
}

// WithShutdownTimeout returns an Option to set the timeout before the server shutdown abruptly
func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *Service) {
		s.shutdownTimeout = timeout
	}
}

// WithPreShutdownDelay returns an Option to set the time waiting for running goroutines
// to finish their jobs before the shutdown starts
func WithPreShutdownDelay(timeout time.Duration) Option {
	return func(s *Service) {
		s.preShutdownDelay = timeout
	}
}

// WithInterruptSignal returns an Option to append a interrupt signal
func WithInterruptSignal(signal os.Signal) Option {
	return func(s *Service) {
		s.interruptSignals = append(s.interruptSignals, signal)
	}
}

// WithStaticDir returns an Option to set the staticDir
func WithStaticDir(dir string) Option {
	return func(s *Service) {
		s.staticDir = dir
	}
}

// WithGRPCServerOption returns an Option to append a gRPC server option
func WithGRPCServerOption(serverOption ...grpc.ServerOption) Option {
	return func(s *Service) {
		s.grpcServerOptions = append(s.grpcServerOptions, serverOption...)
	}
}

// WithGRPCDialOption returns an Option to append a gRPC dial option
func WithGRPCDialOption(dialOption ...grpc.DialOption) Option {
	return func(s *Service) {
		s.grpcDialOptions = append(s.grpcDialOptions, dialOption...)
	}
}

// WithGRPCMuxOption returns an Option to append a mux option
func WithGRPCMuxOption(muxOption ...runtime.ServeMuxOption) Option {
	return func(s *Service) {
		s.muxOptions = append(s.muxOptions, muxOption...)
	}
}

// WithHTTPServer returns an Option to set the http server, note that the Addr and Handler will be
// reset in startGRPCGateway(), so you are not able to specify them
func WithHTTPServer(server *http.Server) Option {
	return func(s *Service) {
		s.HTTPServer = server
	}
}

// WithLogger uses the provided logger
func WithLogger(logger Logger) Option {
	return func(s *Service) {
		s.logger = logger
	}
}

// WithPrometheus enble prometheus config.
func WithPrometheus(b bool) Option {
	return func(s *Service) {
		s.enablePrometheus = b
	}
}

// WithReverseProxyFunc add handlerFromEndpoint to http gw endPoint
func WithReverseProxyFunc(reverseProxyFunc ...ReverseProxyFunc) Option {
	return func(s *Service) {
		s.reverseProxyFuncs = append(s.reverseProxyFuncs, reverseProxyFunc...)
	}
}

// WithRoutes adds additional routes.
func (s *Service) WithRoutes(routes ...Route) Option {
	return func(s *Service) {
		s.routes = append(s.routes, routes...)
	}
}

func (s *Service) apply(opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}
