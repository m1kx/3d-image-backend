package gif

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"time"

	"github.com/m1kx/image/internal/images"
	"github.com/m1kx/image/util"
)

func CreateGif(imageInput *image.Image, points *[]util.Vec2) ([]byte, error) {

	start := time.Now()
	splittedImages := images.SplitImage(imageInput)
	durationSplit := time.Since(start)
	splitTime := float64(durationSplit.Milliseconds())
	fmt.Printf("Took %dms to split images\n", int(splitTime))
	offsets, maxOffset := util.GetOffsets(points)
	framesImg := images.GetOffsettedFrames(splittedImages, offsets, maxOffset)
	dithered := Dither(framesImg)
	buf, err := BufferDithered(dithered)
	if err != nil {
		return []byte{}, err
	}
	duration := time.Since(start)
	totalTime := float64(duration.Milliseconds())
	fmt.Printf("Ran %dms\n", int(totalTime))
	return buf.Bytes(), nil
}

func BufferDithered(dithered *gif.GIF) (bytes.Buffer, error) {
	var buf bytes.Buffer
	encodeStart := time.Now()
	err := gif.EncodeAll(&buf, dithered)
	if err != nil {
		return buf, err
	}
	durationEncode := time.Since(encodeStart)
	encodeTime := float64(durationEncode.Milliseconds())
	fmt.Printf("Took %dms to encode images\n", int(encodeTime))
	return buf, nil
}
