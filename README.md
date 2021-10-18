# ss

mv vendor vendor_bak<br/>
GOARCH=arm GOARM=5 go get golang.org/x/crypto/ssh/terminal<br/>
...<br/>
GOARCH=arm GOARM=5 go build -v -o ./bin/ ...<br/>
