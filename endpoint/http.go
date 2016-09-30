package endpoint

import (
	"github.com/sqdron/squad/endpoint/http"
)

func Http() *transport {
	return NewTransport(&http.HttpTransport{})
}

//
//type HttpOptions struct {
//	host string
//	port string
//	method string
//}
//
//type httpEndpoint struct{
//	m Middleware
//	ctx          context.Context
//	client *http.Client
//	options *HttpOptions
//	bufferedStream bool
//}

//
//func Http(cfg squad.IConfigurator) *EndPoint{
//	return &EndPoint{client.Default, cfg, &HttpOptions{"", "8080", http.MethodGet}, bool}
//}
//
//func (e *httpEndpoint) getAddres() string{
//	return ":" + e.options.port
//}
//
//func (e *httpEndpoint) getBody() string{
//	return "hello"
//}

//func (e *httpEndpoint) Publish() chan <- Message {
//	out := make(chan Message);
//	httpSend := func(ctx context.Context, message chan <- Message) (<- chan Message, error) {
//		ctx, cancel := context.WithCancel(ctx)
//		defer cancel()
//		defer close(out);
//
//		m := e.m(<-out)
//		req, _ := http.NewRequest(e.options.method, "", "")
//		resp, _ := ctxhttp.Do(ctx, e.client, req)
//		if !e.bufferedStream {
//			defer resp.Body.Close()
//		}
//		return resp, nil
//	}
//	go httpSend(nil, out)
//	return out;
//}
//
//func (e *httpEndpoint) Listen() EndpointHandler{
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		ctx, cancel := context.WithCancel(ctx)
//		defer cancel()
//
//		req, _ := http.NewRequest(e.options.method, e.getAddres(), e.getBody())
//		resp, _ := ctxhttp.Do(ctx, e.client, req)
//		if !e.bufferedStream {
//			defer resp.Body.Close()
//		}
//		return resp, nil
//	}
//}

//
//// ServeHTTP implements http.Handler.
//func (s httpEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//ctx := s.ctx

//for _, f := range s.before {
//	ctx = f(ctx, r)
//}

//request, err := s.dec(ctx, r)
//if err != nil {
//	s.logger.Log("err", err)
//	s.errorEncoder(ctx, Error{Domain: DomainDecode, Err: err}, w)
//	return
//}

//response, err := s.e(ctx, request)
//if err != nil {
//	s.logger.Log("err", err)
//	s.errorEncoder(ctx, Error{Domain: DomainDo, Err: err}, w)
//	return
//}
//
//for _, f := range s.after {
//	ctx = f(ctx, w)
//}
//
//if err := s.enc(ctx, w, response); err != nil {
//	s.logger.Log("err", err)
//	s.errorEncoder(ctx, Error{Domain: DomainEncode, Err: err}, w)
//	return
//}
//}
