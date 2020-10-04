FROM golang:1.14

RUN mkdir /app
COPY . /app
WORKDIR /app

EXPOSE 3000
RUN go build -o main . 
CMD [ "/app/main" ]