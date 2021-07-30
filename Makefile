logo:
	go run . img/logo.png
	mogrify -gravity center -pointsize 96 -fill yellow \
		-font DejaVu-Sans-Book -annotate +0+0 "fakespace" \
		img/logo.png
