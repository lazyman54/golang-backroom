package design

type Server struct {
	Addr     string
	Port     int64
	Protocol string
	Timeout  int64
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(server *Server) {
		server.Protocol = p
	}
}
func Timeout(p int64) Option {
	return func(server *Server) {
		server.Timeout = p
	}
}

func newServer(addr string, port int64, options ...func(*Server)) *Server {
	ser := &Server{Addr: "1213", Port: 344}

	for _, option := range options {
		option(ser)
	}
	return ser
}

func OptionalSolution() {
	newServer("1213", 344, Timeout(12), Protocol("tcp"))
}
