package initialize

import (
	appcommon "app/cmd/stream-service/app_common"
	"log"
)

func initProcessStream() {
	log.Println("listen event create process")
	for l := range appcommon.GetChanListenAddProcessStream() {
		go func(l string) {
			process := appcommon.GetProcessStream(l)
			for blob := range process {
				log.Printf("process-%s: %s", l, blob)
			}
		}(l)
	}
}
