package initialize

import (
	grpchandle "app/cmd/proxy-service/delivery/grpc"
	"sync"
)

func Run() {
	listSync := []func(){
		// runAppCommon,
	}

	for _, f := range listSync {
		f()
	}

	listAsync := []func(){
		grpchandle.Register,
	}

	var wg sync.WaitGroup
	for _, f := range listAsync {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(f)
	}

	wg.Wait()
}
