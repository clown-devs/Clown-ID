FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app

RUN make check
RUN make build

CMD ["/app/bin/server"]