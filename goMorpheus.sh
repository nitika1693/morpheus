# Download,Setup and Test Go as described here: https://golang.org/doc/install. 

# Setup PATH and GOPATH environment variables.
export PATH=$PATH:/home/user/go1.12.5.linux-amd64/go/bin
export GOPATH=/home/user/goLangProjects

# Now we are ready to run Morpheus. Below command will skip the _test file and run Morpheus on port 8080 (port mentioned in env.properties). 
# It will also send the logs to file /tmp/morpheusLog.log
go run $(ls -t | grep -v _test | grep .go) env.properties > /tmp/morpheusLog.log 2>&1 &

# Test if Morpheus has run successfully.
curl -X GET -i http://localhost:8080


# If you face below error msg while running go test:
# go test exec: "gcc": executable file not found in $PATH
# Set below environment variable
export CGO_ENABLED=0

