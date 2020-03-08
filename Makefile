SHELL = /bin/bash

app:
	go build .
	appify -name "Wakago" -icon ./assets/logo.png wakago
	zip -r Wakago Wakago.app

clean:
	rm -rf wakago Wakago.app Wakago.zip
