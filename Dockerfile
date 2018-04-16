FROM golang
COPY . /go/src/morningo
WORKDIR /go/src/morningo
RUN make deps
EXPOSE 3000
CMD cd /go/src/morningo
