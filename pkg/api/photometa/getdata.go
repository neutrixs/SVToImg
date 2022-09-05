package photometa

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(panoid string) (Data, error) {
	url := GetURL(panoid)
	resp, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return Data{}, err
	}

	respStr := string(respByte)[5:]

	result := [][][]interface{}{}
	json.Unmarshal([]byte(respStr), &result)

	// what the hell is this
	organisedData := Data{
		Panoid: result[1][0][1].([]interface{})[1].(string),
		Resolution: struct{X float64; Y float64}{
			X: result[1][0][2].([]interface{})[2].([]interface{})[1].(float64),
			Y: result[1][0][2].([]interface{})[2].([]interface{})[0].(float64),
		},
		Location: struct{Name string; CountryCode string}{
			Name: result[1][0][3].([]interface{})[2].([]interface{})[0].([]interface{})[0].(string),
			CountryCode: result[1][0][3].([]interface{})[2].([]interface{})[0].([]interface{})[1].(string),
		},
		Position: Position{
			Lat: result[1][0][5].([]interface{})[0].([]interface{})[1].([]interface{})[0].([]interface{})[2].(float64),
			Long: result[1][0][5].([]interface{})[0].([]interface{})[1].([]interface{})[0].([]interface{})[3].(float64),
			Altitude: result[1][0][5].([]interface{})[0].([]interface{})[1].([]interface{})[1].([]interface{})[0].(float64),
			Pitch: result[1][0][5].([]interface{})[0].([]interface{})[1].([]interface{})[2].([]interface{})[1].(float64),
			Roll: result[1][0][5].([]interface{})[0].([]interface{})[1].([]interface{})[2].([]interface{})[2].(float64),
			Yaw: result[1][0][5].([]interface{})[0].([]interface{})[1].([]interface{})[2].([]interface{})[0].(float64),
		},
		Date: Date{
			Year: int(result[1][0][6].([]interface{})[7].([]interface{})[0].(float64)),
			Month: int(result[1][0][6].([]interface{})[7].([]interface{})[1].(float64)),
		},
		NearbyPanorama: getNearbyPanos(result),
		HistoricalCoverage: getHistoricalCoverage(result),
	}

	return organisedData, nil
}

func getNearbyPanos(data [][][]interface{}) []SimplePanoramaData {
	arr := []SimplePanoramaData{}
	panos := data[1][0][5].([]interface{})[0].([]interface{})[3].([]interface{})[0].([]interface{})

	for _, pano := range panos {
		panoData := SimplePanoramaData{
			Panoid: pano.([]interface{})[0].([]interface{})[1].(string),
			Position: Position{
				Lat: pano.([]interface{})[2].([]interface{})[0].([]interface{})[2].(float64),
				Long: pano.([]interface{})[2].([]interface{})[0].([]interface{})[3].(float64),
				Altitude: pano.([]interface{})[2].([]interface{})[1].([]interface{})[0].(float64),
				Pitch: pano.([]interface{})[2].([]interface{})[2].([]interface{})[1].(float64),
				Roll: pano.([]interface{})[2].([]interface{})[2].([]interface{})[2].(float64),
				Yaw: pano.([]interface{})[2].([]interface{})[2].([]interface{})[0].(float64),
			},
		}

		arr = append(arr, panoData)
	}

	return arr
}

func getHistoricalCoverage(data [][][]interface{}) []HistoricalCoverage {
	arr := []HistoricalCoverage{}
	panos := getNearbyPanos(data)

	historicals := data[1][0][5].([]interface{})[0].([]interface{})[8].([]interface{})
	for _, historical := range historicals {
		index := int(historical.([]interface{})[0].(float64))
		date := Date{
			Year: int(historical.([]interface{})[1].([]interface{})[0].(float64)),
			Month: int(historical.([]interface{})[1].([]interface{})[1].(float64)),
		}

		historicalData := HistoricalCoverage{
			Panoid: panos[index].Panoid,
			Position: panos[index].Position,
			Date: date,
		}

		arr = append(arr, historicalData)
	}

	return arr
}