package appcommon

import (
	"app/generated/proto/servicegrpc"
	constant "app/internal/constants"
	logapp "app/pkg/log"
	"fmt"
	"log"
	"sync"

	"google.golang.org/grpc"
)

func GetListGrpcQuantityGrpc() []servicegrpc.QuantityServiceClient {
	log.Println("get proxy quantity")
	listIpServerGrpcQuantity := []string{
		"localhost:9008",
	}
	listQuantityGrpc := []servicegrpc.QuantityServiceClient{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, address := range listIpServerGrpcQuantity {
		wg.Add(1)
		go func(address string) {
			defer wg.Done()
			grpcConn, err := grpc.NewClient(address, grpc.WithInsecure())
			if err != nil {
				logapp.Logger(fmt.Sprintf("proxy connect grpc ip: %s", address), err.Error(), constant.ERROR_LOG)
				return
			}

			grpcClient := servicegrpc.NewQuantityServiceClient(grpcConn)

			mu.Lock()
			listQuantityGrpc = append(listQuantityGrpc, grpcClient)
			mu.Unlock()
		}(address)
	}

	wg.Wait()

	return listQuantityGrpc
}

func GetListGrpcMergeBlobGrpc() []servicegrpc.MergeBlobServiceClient {
	listIpServerGrpcMergeBlob := []string{
		"localhost:9007",
	}
	listMergeBlobGrpc := []servicegrpc.MergeBlobServiceClient{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, address := range listIpServerGrpcMergeBlob {
		wg.Add(1)
		go func(address string) {
			defer wg.Done()
			grpcConn, err := grpc.NewClient(address, grpc.WithInsecure())
			if err != nil {
				logapp.Logger(fmt.Sprintf("proxy connect grpc ip: %s", address), err.Error(), constant.ERROR_LOG)
				return
			}

			grpcClient := servicegrpc.NewMergeBlobServiceClient(grpcConn)

			mu.Lock()
			listMergeBlobGrpc = append(listMergeBlobGrpc, grpcClient)
			mu.Unlock()
		}(address)
	}

	wg.Wait()

	return listMergeBlobGrpc
}
