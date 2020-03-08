app:
	go build .
	go get github.com/machinebox/appify
	appify -name "Wakago" -icon ./assets/logo.png wakago
