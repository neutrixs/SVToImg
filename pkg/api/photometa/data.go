package photometa

type Position struct {
	Lat 		float64
	Long 		float64
	Altitude 	float64
	Pitch 		float64
	Roll 		float64
	Yaw 		float64
}

type Date struct {
	Year 	int
	Month 	int
}

type SimplePanoramaData struct {
	Panoid 		string
	Position 	Position
}

type HistoricalCoverage struct {
	Panoid 		string
	Position 	Position
	Date
}

type Data struct {
	Panoid 				string
	Resolution 			struct {
		X 				float64
		Y 				float64
	}
	Location 			struct {
		Name			string
		CountryCode 	string
	}
	Position 			Position
	NearbyPanorama 		[]SimplePanoramaData
	Date 				struct {
		Year 			int
		Month 			int
	}
	HistoricalCoverage 	[]HistoricalCoverage
}