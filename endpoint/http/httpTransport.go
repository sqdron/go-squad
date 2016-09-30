package http
//
//import (
//	"context"
//	//"github.com/sqdron/squad/endpoint"
//	//"net"
//	"net/http"
//)
//
//type HttpTransport struct {
//	Options        *Options
//	ctx            context.Context
//	client         *http.Client
//	bufferedStream bool
//}
//
//type serveHandler func(http.ResponseWriter, *http.Request) (int, error)

//
//func (t *HttpTransport) Listen() <-chan endpoint.Message {
//	server :=
//		http.Server{
//			Addr:    ":" + m.options.Port,
//			Handler: m,
//		}
//
//	addr := server.Addr
//	if addr == "" {
//		addr = ":http"
//	}
//	ln, err := net.Listen("tcp", addr)
//	if err != nil {
//		//return err
//	}
//	server.Serve(newStoppableListener(ln.(IListenerAcceptor), stop))
//	//return func(ctx context.Context, request interface{}) (interface{}, error) {
//	//	ctx, cancel := context.WithCancel(t.ctx)
//	//	defer cancel()
//	//
//	//	req, _ := http.NewRequest(e.options.method, e.getAddres(), e.getBody())
//	//	resp, _ := ctxhttp.Do(ctx, e.client, req)
//	//	if !e.bufferedStream {
//	//		defer resp.Body.Close()
//	//	}
//	//	return resp, nil
//	//}
//}
//
////func NewHttpTransport() *endpoint.Transport{
////	return endpoint.NewTransport(&HttpTransport{})
////}
//
//func (s *HttpTransport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//	//for _, f := range s.before {
//	//	ctx = f(ctx, r)
//	//}
//	//
//	//request, err := s.dec(ctx, r)
//	//if err != nil {
//	//	s.logger.Log("err", err)
//	//	s.errorEncoder(ctx, Error{Domain: DomainDecode, Err: err}, w)
//	//	return
//	//}
//	//
//	//response, err := s.e(ctx, request)
//	//if err != nil {
//	//	s.logger.Log("err", err)
//	//	s.errorEncoder(ctx, Error{Domain: DomainDo, Err: err}, w)
//	//	return
//	//}
//	//
//	//for _, f := range s.after {
//	//	ctx = f(ctx, w)
//	//}
//	//
//	//if err := s.enc(ctx, w, response); err != nil {
//	//	s.logger.Log("err", err)
//	//	s.errorEncoder(ctx, Error{Domain: DomainEncode, Err: err}, w)
//	//	return
//	//}
//}
