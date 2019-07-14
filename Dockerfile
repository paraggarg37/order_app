FROM golang:latest
WORKDIR $GOPATH/src
RUN mkdir -p  $GOPATH/src/github.com/paraggarg37/order_app
ADD . $GOPATH/src/github.com/paraggarg37/order_app/
WORKDIR $GOPATH/src/github.com/paraggarg37/order_app/cmd
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -v
RUN go build -o logistics

EXPOSE 8080

CMD ./logistics
