# grpc服务端编程步骤
## 第1步net.dail生成一个listen对象

## 第2步grpc.NewServer()
生成一个grpc.Server对象

## 第3步pb.Register*(s, &server{})
将实现的service服务端注册到grpc.Server对象。

## 第4步:grpc.Server.serve(listen对象)方法

### 4.1 执行for循环：rawConn, err := lis.Accept()
````
rawConn, err := lis.Accept()
````

### 4.2 单独开一个协程处理rawConn
````
//google.golang.org/grpc/server.go 
func (s *Server) handleRawConn(rawConn net.Conn) {
    if s.quit.HasFired() {
		rawConn.Close()
		return
	}
	rawConn.SetDeadline(time.Now().Add(s.opts.connectionTimeout))
	conn, authInfo, err := s.useTransportAuthenticator(rawConn)
	if err != nil {
		// ErrConnDispatched means that the connection was dispatched away from
		// gRPC; those connections should be left open.
		if err != credentials.ErrConnDispatched {
			s.mu.Lock()
			s.errorf("ServerHandshake(%q) failed: %v", rawConn.RemoteAddr(), err)
			s.mu.Unlock()
			grpclog.Warningf("grpc: Server.Serve failed to complete security handshake from %q: %v", rawConn.RemoteAddr(), err)
			rawConn.Close()
		}
		rawConn.SetDeadline(time.Time{})
		return
	}

	// Finish handshaking (HTTP2)
	st := s.newHTTP2Transport(conn, authInfo)
	if st == nil {
		return
	}

	rawConn.SetDeadline(time.Time{})
	if !s.addConn(st) {
		return
	}
	go func() {
		s.serveStreams(st)
		s.removeConn(st)
	}()
}
````
### 4.3 最核心的代码

`````
func (s *Server) serveStreams(st transport.ServerTransport) {
    defer st.Close()
    var wg sync.WaitGroup
    // HandleStreams 是注册 grpc server处理 http2 stream 数据的处理函数
    st.HandleStreams(func(stream *transport.Stream) { 
        wg.Add(1)
        //每次有新request时会调用这个方法， 这个方法就是开新的协程处理请求
        go func() {
            defer wg.Done()
            s.handleStream(st, stream, s.traceInfo(st, stream))
        }()  
    }, func(ctx context.Context, method string) context.Context {
        if !EnableTracing {
            return ctx
        }    
        tr := trace.New("grpc.Recv."+methodFamily(method), method)
        return trace.NewContext(ctx, tr)
    })   
    wg.Wait()
}
`````