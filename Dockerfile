FROM golang:1.18-stretch

ENV BIN_FILE /usr/src/app/dist/api
ENV GO111MODULE on

WORKDIR /usr/src/app
COPY . /usr/src/app

# make build
RUN make build-api ; chmod a+x ./dist/api
ENV TZ Europe/Moscow

ENV CONFIG_FILE /usr/src/app/dist/api.json
COPY ./configs/api.json ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}