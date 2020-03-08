app:
	go build .
	go install -v github.com/machinebox/appify
	ls ${GOPATH}/bin
	appify -name "Wakago" -icon ./assets/logo.png wakago
