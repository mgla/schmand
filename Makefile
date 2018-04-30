NAME=schmandbot
LAMBDANAME=$(NAME)

LAMBDA_RUNTIME=go1.x

build:
		go-bindata data
		go build -o ${NAME} -v
zip:
		zip ${NAME}.zip ${NAME}
clean:
	rm -f ${NAME} ${NAME}.zip  bindata.go
deploy: clean build zip updatelambda clean
updatelambda:
	aws lambda update-function-code --function-name ${LAMBDANAME} --zip-file fileb://${NAME}.zip
