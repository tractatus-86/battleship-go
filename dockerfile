# Container image that runs your code
FROM golang:1.16
RUN mkdir /app
ADD . /app
WORKDIR /app/cmd/battleship/
RUN go build 
CMD [ "/bin/sh" ]
