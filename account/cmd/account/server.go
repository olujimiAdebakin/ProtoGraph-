package account

import (
	"context"
	"fmt"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    "github.com/olujimiAdebakin/ProtoGraphql/account/pb"
)

type grpcServer struct {
	service Service
}

func LstenGRPCServer(service Service, port int) error {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        return  err
    }

    grpcSrv := grpc.NewServer()
    pb.RegisterAccountServiceServer(grpcSrv, &grpcServer{service: service})
    reflection.Register(grpcSrv)
    return grpcSrv.Serve(lis)
}


func (s *grpcServer)PostAccount(ctx context.Context, req *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	    // Call business logic, passing all required params
    acc, err := s.service.PostAccount(ctx, req.Name, req.Email, req.Password)
    if err != nil {
        return nil, err // Return gRPC error to client
    }

    // Map internal Account to gRPC response message
    resp := &pb.PostAccountResponse{Account: &pb.Account{
        Id:    acc.ID,
        Name:  acc.Name,
        Email: acc.Email,
    }}

    return resp, nil
}

func (s *grpcServer)GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
    // Call business logic to get account by ID
	acc, err := s.service.GetAccount(ctx, req.Id)
    if err != nil {
	  return nil, err // Return gRPC error to client
    }

    return &pb.GetAccountResponse{Account: &pb.Account{
        Id:    acc.ID,
        Name:  acc.Name,
        Email: acc.Email,
    },

    }, nil
}



func (s *grpcServer)ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {

    acc, err := s.service.ListAccounts(ctx, req.Skip, req.Take)
    if err != nil {
        return nil, err // Return gRPC error to client
    }

    resp := &pb.ListAccountsResponse{ctx}
    for _, acc := range acc {
        resp.Accounts = append(resp.Accounts, &pb.Account{
            Id:    acc.ID,
            Name:  acc.Name,
            Email: acc.Email,
        })
    }
    return resp, nil
	
}

func (s *grpcServer)DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {

    deletedAccount, err := s.service.DeleteAccount(ctx, req.Id)
    if err != nil {
        return nil, err // Return gRPC error to client
    }
    return &pb.DeleteAccountResponse{      Success: true,
        Account: &pb.Account{  
            Id: deletedAccount.ID,
            Name: deletedAccount.Name,
            Email: deletedAccount.Email,
            Password: deletedAccount.Password,
        }, 
        }, nil
}


