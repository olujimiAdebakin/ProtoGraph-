
package account

import (
	"context"    // For context management (timeouts, cancellation)
	"fmt"        
	"net"  
    "time"    
      

	"google.golang.org/grpc"          
	"google.golang.org/grpc/reflection" 

	"github.com/olujimiAdebakin/ProtoGraphql/account/pb"
)

// grpcServer wraps the business logic service and implements gRPC methods
type grpcServer struct {
	service Service // Business logic layer interface
}

// ListenGRPCServer starts a gRPC server on the specified port
// service: Business logic implementation
// port: TCP port to listen on (e.g., 50051)
// Returns error if server fails to start
func ListenGRPCServer(service Service, port int) error {
	// Create TCP listener on specified port (e.g., ":50051")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
         // Return error if port binding fails
		return err
	}

	// Create new gRPC server instance
	grpcSrv := grpc.NewServer()
	
	// Register our gRPC server implementation with the gRPC framework
	// This connects our grpcServer methods to the AccountService protobuf definition
	pb.RegisterAccountServiceServer(grpcSrv, &grpcServer{service: service})
	
	// Enable gRPC reflection - allows tools like grpcurl to discover API
	reflection.Register(grpcSrv)
	
	// Start serving gRPC requests - blocks until server stops
	return grpcSrv.Serve(lis)
}

// PostAccount handles account creation requests via gRPC
// ctx: Request context (carries deadlines, cancellation signals)
// req: Incoming gRPC request with account details (name, email, password)
// Returns: gRPC response with created account or error
func (s *grpcServer) PostAccount(ctx context.Context, req *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	// Delegate to business logic layer to create account
	// Pass all required parameters from gRPC request
	account, err := s.service.PostAccount(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		// If business logic fails, propagate error to gRPC client
		return nil, err
	}

	// Map internal business account to gRPC response format
	// Note: Password is not included in response for security
	resp := &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:    account.ID,   
			Name:  account.Name,  
			Email: account.Email, 
		},
	}
// Return success response
	return resp, nil 
}

// GetAccount handles single account retrieval requests via gRPC
// ctx: Request context
// req: Incoming request with account ID to fetch
// Returns: gRPC response with account details or error
func (s *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	// Call business logic to fetch account by ID
	account, err := s.service.GetAccount(ctx, req.Id)
	if err != nil {
		// Return gRPC error if account not found or other failure
		return nil, err
	}

	// Convert internal account to gRPC response
	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:    account.ID,    // Map ID
			Name:  account.Name,  // Map name
			Email: account.Email, // Map email
			
		},
	}, nil
}

// ListAccounts handles paginated account listing requests via gRPC
// ctx: Request context
// req: Incoming request with pagination parameters (skip, take)
// Returns: gRPC response with list of accounts or error
func (s *grpcServer) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	// Call business logic with pagination parameters
	// skip: Number of records to skip (offset)
	// take: Number of records to fetch (limit)
	accounts, err := s.service.ListAccounts(ctx, req.Skip, req.Take)
	if err != nil {
        // Propagate business logic errors
		return nil, err 
	}

	// Initialize empty response
	resp := &pb.ListAccountsResponse{}
	
	// Convert each internal account to gRPC format
	for _, account := range accounts {
		// Append converted account to response list
		resp.Accounts = append(resp.Accounts, &pb.Account{
			Id:    account.ID,    // Map ID
			Name:  account.Name,  // Map name
			Email: account.Email, // Map email
		})
	}
	// Return paginated account list
	return resp, nil 
}

// DeleteAccount handles account deletion requests via gRPC
// ctx: Request context
// req: Incoming request with account ID to delete
// Returns: gRPC response confirming deletion or error
func (s *grpcServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {

    // Set timeout for deletion operation
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Call business logic to delete account
	// Returns deleted account info (optional) and error
	deletedAt, err := s.service.DeleteAccount(ctx, req.Id)
	if err != nil {
        // Return error if deletion fails
		return nil, err 
	}
	
	// Return success response with optional deleted account info
	return &pb.DeleteAccountResponse{
		Success: true, // Confirmation flag
         Message: "Account deleted successfully",
         DeletedAccountId: req.Id,
         deletedAt.Format(time.RFC3339),
		//  Account: &pb.Account{
		// 	Id:       deletedAccount.ID,       // Map ID of deleted account
		// 	Name:     deletedAccount.Name,     // Map name of deleted account
		// 	Email:    deletedAccount.Email,    // Map email of deleted account
		// },
	}, nil
}