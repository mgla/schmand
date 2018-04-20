build:
		go-bindata data
		go build
zip:
		zip schmand.zip schmand
clean:
		rm -f schmand schmand.zip bindata.go
deploy: clean build zip updatelambda clean
updatelambda:
	aws lambda update-function-code --function-name schmandbod --zip-file fileb://schmand.zip
