app:
	go build .
	go install -v github.com/machinebox/appify
	ls ${HOME}/go/bin
	${HOME}/go/bin/appify -name "Wakago" -icon ./assets/logo.png wakago
