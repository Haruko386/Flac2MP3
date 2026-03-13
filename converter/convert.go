package converter

import (
	"Flac2MP3/utils"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Convert(inputFile string) error {
	// 检查文件是否存在
	if err := utils.CheckFileExists(inputFile); err != nil {
		return err
	}
	// 检查是否为flac格式
	if filepath.Ext(inputFile) != ".flac" {
		return errors.New("not a flac file")
	}
	// 生成文件输出名
	outputFile := utils.GetOutputPath(inputFile)
	fmt.Printf("Converting %s to %s\n", inputFile, outputFile)

	// 构建FFmpeg命令  -i: 输入文件 -q:a 0: 使用LAME编码器的最高质量VBR压缩 -map a: 忽略可能存在的视频或图片封面
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:a", "0", "-map", "a", outputFile)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %s", err)
	}

	fmt.Printf("Converting %s to %s\n", inputFile, outputFile)
	return nil
}
