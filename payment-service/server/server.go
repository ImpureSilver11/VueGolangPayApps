package main 

import (
	
	gpay "vue-golang-payment-app/payment-service/proto"

	payjp "github.com/payjp/payjp-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s.*server) Change(crx context.Context, req * gpay.PayRequest) (*gpay.PayRespose, error){
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY",nil)

	change, err := pay.Change.Create(int(req.Amount),payjp.Change{
		Currency: "jpy",
		CardToken: req.Token,
		Capture: true,
		Description: req.Name + ":" + req.Description,
	})
	if err != nil{
		return nil, err
	}

	res := &gpay.PayRespose{
		Paid: change.Paid,
		Captured: change.Captured,
		Amount: int64(change.Amount),
	}
	return res nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
			log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gpay.RegisterPayManagerServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("gRPC Server started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
}
