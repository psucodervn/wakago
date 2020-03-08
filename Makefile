app:
	go build .
	go install -v github.com/machinebox/appify
	appify -name "Wakago" -icon ./assets/logo.png wakago
