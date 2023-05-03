FROM golang:1.16

ENV TZ=Asia/Jakarta
ENV HOST_IP=127.0.0.1
ENV PACKAGE_PATH=template
ENV DB_HOST=
ENV DB_PORT=
ENV DB_USER=
ENV DB_PASSWORD=
ENV DB_NAME=
ENV JWT_LIFE_TIME=
ENV JWT_SECRET=
ENV ENVIRONMENT=

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN date
RUN mkdir -p /go/src/
WORKDIR /go/src/$PACKAGE_PATH
COPY . /go/src/$PACKAGE_PATH/

RUN sed -i -E "s/(([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)|localhost):[0-9]+/${HOST_IP}/" docs/index.html
RUN sed -i -E "s/(([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)|localhost):[0-9]+/${HOST_IP}/" docs/swagger.yaml

RUN go get
RUN go build -o dist
RUN mkdir -p logs
ENTRYPOINT ./dist
EXPOSE 1323
