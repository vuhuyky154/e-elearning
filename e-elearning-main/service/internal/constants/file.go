package constant

type FOLDER string

var (
	THUMNAIL_COURSE FOLDER = "cmd/core-service/data/file/thumnail_course"

	UPLOAD_MP4_VIDEO FOLDER = "cmd/upload-mp4-service/data/video"

	ENCODING_HLS_ENDCODING FOLDER = "cmd/encoding-service/data/encoding"
	ENCODING_HLS_VIDEO     FOLDER = "cmd/encoding-service/data/video"

	VIDEO_HLS_VIDEO FOLDER = "cmd/video-hls-service/data/video"

	MERGE_BLOB FOLDER = "cmd/merge-blob/data/stream"
)
