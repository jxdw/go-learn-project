package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-framework-01/golang05_grpc_basic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
)


type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}
type toDoServiceServer struct {
	db *sql.DB
}
func NewToDoServiceServer(db *sql.DB) *toDoServiceServer {
	return &toDoServiceServer{db: db}
}
// connect returns SQL database connection from the pool
func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

type Greeter  struct{}

func (greeter *Greeter) Helloservice(ctx context.Context, req *protocol.RequestMessage) (*protocol.ResponseMessage, error) {
	log.Print("received ")
	//conn, err :=dbServer.connect(ctx)
	//if err!=nil {
	//
	//}
	//conn.QueryContext(ctx,"SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",)
	responseMessage:=protocol.ResponseMessage{Msg:"hello "+req.Name}
	return &responseMessage,nil
}
var dbServer *toDoServiceServer
func init(){
	var cfg Config
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()
	dbServer=NewToDoServiceServer(db)
}
func main() {

	lis,err:=net.Listen("tcp","10.3.20.223:8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer()
	greeter:=Greeter{}
	protocol.RegisterGreeterServer(s,&greeter)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("eddycjy: go-grpc-example"))
	//})
	//http.ListenAndServe("10.3.20.223:8083",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
	//		s.ServeHTTP(w, r)
	//	} else {
	//		mux.ServeHTTP(w, r)
	//	}
	//	return
	//}))
}
