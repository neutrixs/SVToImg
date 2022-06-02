package cli

func help() string {
	const text string = 
		"Usage: SVToIMG <source> -o output/file.jpeg [quality]\n"+
		"\n"+
		"sources:\n"+
		"	-p string\n"+
		"		Specify the street view's pano ID\n"+
		"	-u string\n"+
		"		Specify the street view URL\n"+
		"\n"+
		"quality:\n"+
		"	-q int\n"+
		"		Specify the quality of output jpeg between 0-100.\n"+
		"		Defaults to 80"

	return text
}