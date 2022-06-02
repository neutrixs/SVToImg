package api

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"log"
	"net/http"
	"sync"

	"github.com/nfnt/resize"
)

var EmptyImage *image.RGBA = image.NewRGBA(
	image.Rectangle{
		image.Point{0, 0},
		image.Point{0, 0},
	},
)

type EachgenData struct {
	zoomLevel int
	tileResolution int
	xAmount int
	yAmount int
}

var GenData = map[int]EachgenData {
	1: {
		zoomLevel: 3,
		tileResolution: 512,
		xAmount: 6 + 1,
		yAmount: 3 + 1,
	},
	2: {
		zoomLevel: 5,
		tileResolution: 512,
		xAmount: 25 + 1,
		yAmount: 12 + 1,
	},
	// no gen 3 because it's exactly the same as gen 2
	4: {
		zoomLevel: 5,
		tileResolution: 512,
		xAmount: 31 + 1,
		yAmount: 15 + 1,
	},
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetImages(panoid string) (*image.RGBA, error) {
	var wg sync.WaitGroup
	imagesYX := make(map[int]map[int]image.Image)
	errorChannel := make(chan error)

	var setImagesYX = func(y, x int, content image.Image) {
		imagesYX[y][x] = content
	}

	gen, err := GetGeneration(panoid)

	if err != nil {
		return EmptyImage, err
	}

	config := GenData[gen]

	downloaded := 0
	total := config.xAmount * config.yAmount

	fmt.Print("\n")

	for y := 0; y < config.yAmount; y++ {
		imagesYX[y] = map[int]image.Image{}
		for x := 0; x < config.xAmount; x++ {
			wg.Add(1)
			go getImage(setImagesYX, &wg, config, panoid, y, x, total, &downloaded, errorChannel)
		}
	}

	wg.Wait()
	close(errorChannel)
	returnedError := <- errorChannel

	// combines the images

	combinedImage := image.NewRGBA(image.Rectangle {
		Min: image.Point{0, 0},
		Max: image.Point{
			config.xAmount * config.tileResolution,
			config.yAmount * config.tileResolution,
		},
	})

	totalIMG := config.yAmount * config.xAmount
	imgHaveBeenCombined := 0

	fmt.Printf("\n\rCombining images: %d/%d", imgHaveBeenCombined, totalIMG)

	for y := 0; y < config.yAmount; y++ {
		for x := 0; x < config.xAmount; x++ {
			rectangle := image.Rectangle {
				Min: image.Point {
					x * config.tileResolution,
					y * config.tileResolution,
				},
				Max: image.Point {
					x * config.tileResolution + config.tileResolution,
					y * config.tileResolution + config.tileResolution,
				},
			}

			draw.Draw(combinedImage, rectangle, imagesYX[y][x], image.Point{0, 0}, draw.Src)
			imgHaveBeenCombined++
			fmt.Printf("\rCombining images: %d/%d", imgHaveBeenCombined, totalIMG)
		}
	}
	fmt.Print("\n")
	
	return combinedImage, returnedError
}

func getImage(setImagesYX func(y, x int, content image.Image), wg *sync.WaitGroup, config EachgenData, panoid string, y, x, total int, downloaded *int, c chan error) {
	defer wg.Done()
	defer func() {
		*downloaded++
		fmt.Printf("\rDownloaded %d/%d images", *downloaded, total)
	}()

	baseURL := "https://streetviewpixels-pa.googleapis.com/v1/tile?cb_client=maps_sv.tactile&panoid=%v&x=%d&y=%d&zoom=%d&nbt=1&fover=2"
	URL := fmt.Sprintf(baseURL, panoid, x, y, config.zoomLevel)

	res, err := http.Get(URL)

	if err != nil {
		c <- err
		return
	}

	resImage, _, err := image.Decode(res.Body)

	if err != nil {
		c <- err
		return
	}

	imageWidth := resImage.Bounds().Dx()

	if imageWidth != config.tileResolution {
		resImage = resize.Resize(uint(config.tileResolution), 0, resImage, resize.Lanczos3)
	}

	setImagesYX(y, x, resImage)
}