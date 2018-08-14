FROM golang
ADD . /gocode/src/logmonitor
WORKDIR /gocode/bin
RUN export GOPATH=/gocode \
	&& GOBIN=/gocode/bin \
	&& cd /gocode/bin \
	&& mkdir config \
	&& cp /gocode/src/logmonitor/_DEPLOY/common/* ./config \
	&& cp /gocode/src/logmonitor/_DEPLOY/stg/* ./config \
	&& go install logmonitor 
CMD cd /gocode/bin && ./logmonitor 
