package metrics

//ServerRequestMetrics metrics fro grpc and http request
var ServerRequestMetrics = MustRegisterCounterVec("request", "default",
	"server request metrics", []string{"status", "rpc", "method", "path", "handler"})
