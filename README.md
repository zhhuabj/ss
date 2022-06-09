# ss
go env -w GO111MODULE=off </br>
#unset GO111MODULE && go env -w GO111MODULE=on </br>
proxychains go get -u -d -v github.com/zhhuabj/ss/... </br>
proxychains go build -v -o ../bin/ github.com/zhhuabj/ss/... </br>
#proxychains go install -v github.com/zhhuabj/ss/... </br>j

GOARCH=arm GOARM=5 go build xxx
