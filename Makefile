app:
	go env
	echo ${PATH}
	go build .
	go get -v github.com/machinebox/appify
	ls ${PATH}
	appify -name "Wakago" -icon ./assets/logo.png wakago
