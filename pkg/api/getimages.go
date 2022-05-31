package api

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"sync"
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
	tileLowerResolution int
	resolutionLowersAtY [2]int
	xAmount int
	yAmount int
}

var GenData = map[int]EachgenData {
	1: {
		zoomLevel: 3,
		tileResolution: 512,
		tileLowerResolution: 512,
		resolutionLowersAtY: [2]int {0, 0},
		xAmount: 6 + 1,
		yAmount: 3 + 1,
	},
	2: {
		zoomLevel: 5,
		tileResolution: 512,
		tileLowerResolution: 256,
		resolutionLowersAtY: [2]int {3, 9},
		xAmount: 25 + 1,
		yAmount: 12 + 1,
	},
	// no gen 3 because it's exactly the same as gen 2
	4: {
		zoomLevel: 5,
		tileResolution: 512,
		tileLowerResolution: 256,
		resolutionLowersAtY: [2]int {4, 11},
		xAmount: 31 + 1,
		yAmount: 15 + 1,
	},
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetImages(panoid string) (*image.RGBA, error) {
	var imagesYX [][]image.Image
	var wg sync.WaitGroup

	gen, err := GetGeneration(panoid)

	if err != nil {
		return EmptyImage, err
	}

	config := GenData[gen]

	downloaded := 0
	total := config.xAmount * config.yAmount

	fmt.Print("\n")

	for y := 0; y < config.yAmount; y++ {
		imagesYX = append(imagesYX, []image.Image{})
		for x := 0; x < config.xAmount; x++ {
			wg.Add(1)
			getImage(&imagesYX, &wg, panoid, y, x, config.zoomLevel, total, &downloaded)
		}
	}

	wg.Wait()
	fmt.Println(imagesYX)

	// FIXME: temporary
	return EmptyImage, nil
}

func getImage(imagesYX *[][]image.Image, wg *sync.WaitGroup, panoid string, y, x, zoom, total int, downloaded *int) {
	defer wg.Done()
	defer func() {
		*downloaded++
		fmt.Printf("\rDownloading images %d/%d", *downloaded, total)
	}()

	baseURL := "https://streetviewpixels-pa.googleapis.com/v1/tile?cb_client=maps_sv.tactile&panoid=%v&x=%d&y=%d&zoom=%d&nbt=1&fover=2"
	URL := fmt.Sprintf(baseURL, panoid, x, y, zoom)

	res, err := http.Get(URL)

	if err != nil {
		log.Fatalln(err)
	}

	resImage, _, err := image.Decode(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	(*imagesYX)[y][x] = resImage
}