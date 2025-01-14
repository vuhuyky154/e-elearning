core:
	go run cmd/core-service/main.go
upload-mp4:
	go run cmd/upload-mp4-service/main.go
encoding:
	go run cmd/encoding-service/main.go
video-hls:
	go run cmd/video-hls-service/main.go
quizz:
	go run cmd/quizz-service/main.go
merge-blob:
	go run cmd/merge-blob/main.go
stream:
	go run cmd/stream-service/main.go
blob:
	go run cmd/blob-service/main.go
quantity:
	go run cmd/quantity-blob-service/main.go
proxy:
	go run cmd/proxy-service/main.go
# push mess to queue with quantity
# job delete video
all:
	go run cmd/core-service/main.go &
	go run cmd/upload-mp4-service/main.go &
	go run cmd/encoding-service/main.go &
	go run cmd/video-hls-service/main.go &
	go run cmd/quizz-service/main.go &
	go run cmd/merge-blob-service/main.go &
	go run cmd/stream-service/main.go &
	go run cmd/quizz-service/main.go &
	go run cmd/quantity-blob-service/main.go &
	wait
