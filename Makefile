SHELL = /bin/bash

app:
	go build .
	appify -name "Wakago" -icon ./assets/logo.png wakago

clean:
	rm wakago
	rm -rf Wakago.app
