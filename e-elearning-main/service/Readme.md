# encoding:
	cmd := exec.Command("ffmpeg",
		"-f", "webm", // Định dạng đầu vào
		"-i", "pipe:0", // Nhận từ stdin
		"-vf", "scale=-2:360", // Giảm độ phân giải
		"-preset", "ultrafast", // Cấu hình preset nhanh
		"-vcodec", "libx264", // Bộ mã hóa video H.264
		"-acodec", "aac", // Bộ mã hóa âm thanh AAC
		"-fflags", "+genpts", // Đảm bảo timestamp chính xác
		"-movflags", "+frag_keyframe+empty_moov", // Đảm bảo phát trực tiếp
		"-f", "mpegts", // Định dạng đầu ra là MPEG-TS
		"pipe:1", // Ghi ra stdout
	)

  cmd := exec.Command("ffmpeg",
		"-f", "webm", // Định dạng đầu vào
		"-i", "pipe:0", // Nhận từ stdin
		"-c:v", "copy", // Copy codec video, không mã hóa lại
		"-c:a", "copy", // Copy codec âm thanh, không mã hóa lại
		"-fflags", "+genpts", // Đảm bảo timestamp chính xác
		"-f", "mpegts", // Định dạng đầu ra là MPEG-TS
		"pipe:1", // Ghi ra stdout
	)