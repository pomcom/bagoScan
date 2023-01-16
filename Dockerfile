FROM golang:latest

RUN apt-get update && apt-get upgrade -y

RUN apt-get update && apt-get install -y \
    golang \
    git \
    nmap \
    testssl.sh 

# set GOPATH and PATH
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# currently using alias to testssl.sh
# could be changed in the app code later
RUN echo "alias testssl='testssl.sh'" >> ~/.bashrc

WORKDIR /app

# copy the app files into container
COPY . .

RUN go build -o bagoScan main.go

RUN mv /app/bagoScan /usr/local/bin/

# run the app
CMD ["bagoScan"]

