DIR = $(shell pwd)
BINFILE = $(shell basename ${DIR})
TMPPATH = /tmp/deploy
DEPLOY_PATH = _DEPLOY
TMPDIR = ${TMPPATH}/${BINFILE}

build:
	@go build -ldflags "-X main.serverName=${BINFILE} -X main.gitCommit=`git rev-parse --short=7 HEAD` -X main.buildTime=`date +%Y-%m-%d_%H:%M:%S`" -o ${BINFILE}
	@echo "${BINFILE} is created."

.PHONY: stg prd clean
prd: build
	@mkdir -p ${TMPDIR}; \
	cp ${BINFILE} ${TMPDIR}; \
	cp -rf  ${DEPLOY_PATH}/common ${TMPDIR};\
	cp -rf  ${DEPLOY_PATH}/prd ${TMPDIR};\
	chmod u+x ${TMPDIR}/*.sh;\
	tar czf ${BINFILE}.prd.`git rev-parse --short=7 HEAD`.tgz -C ${TMPPATH} ${BINFILE}; \
	rm -rf ${TMPDIR}
	@echo "${BINFILE}.prd.`git rev-parse --short=7 HEAD`.tgz is created."

stg: build
	@mkdir -p ${TMPDIR}; \
	cp ${BINFILE} ${TMPDIR}; \
	cp -rf  ${DEPLOY_PATH}/common ${TMPDIR};\
	cp -rf  ${DEPLOY_PATH}/stg ${TMPDIR};\
	chmod u+x ${TMPDIR}/*.sh;\
	tar czf ${BINFILE}.prd.`git rev-parse --short=7 HEAD`.tgz -C ${TMPPATH} ${BINFILE}; \
	rm -rf ${TMPDIR}
	@echo "${BINFILE}.prd.`git rev-parse --short=7 HEAD`.tgz is created."

clean:
	@rm -rf ${BINFILE}; \
	rm -f *.tgz
	@echo "finish."
