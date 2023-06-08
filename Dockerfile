FROM golang:1.19-alpine3.16

RUN mkdir /restaurant-management

COPY . /restaurant-management

WORKDIR /restaurant-management

LABEL Name=restaurant-management Version=0.0.1

RUN go build -o restaurant-management-api

EXPOSE  80

CMD [ "./restaurant-management-api", "--m=true", "--c=true" ]