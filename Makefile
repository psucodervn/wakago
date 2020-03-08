SHELL = /bin/bash

app:
	go build .
	appify -name "Wakago" -icon ./assets/logo.png wakago
	zip Wakago.zip Wakago.app

clean:
	rm wakago
	rm -rf Wakago.app
