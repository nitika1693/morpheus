export PATH=$PATH:/home/chiragjain/mobikwik/softwares/go1.12.5.linux-amd64/go/bin
export GOPATH=/home/chiragjain/mobikwik/codebase/goLangProjects
export CGO_ENABLED=0
go run $(ls -t | grep -v _test | grep .go) env.properties > /tmp/morpheusLog.log 2>&1 &

curl -X GET -i http://localhost:8080
