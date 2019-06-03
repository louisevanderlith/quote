# quote
Quote.API holds Invoices. Thats it.

## Run with Docker
* $ docker build -t avosa/quote:dev .
* $ docker rm QuoteDEV
* $ docker run -d -p 8099:8099 --network mango_net --name QuoteDEV avosa/quote:dev
* $ docker logs QuoteDEV