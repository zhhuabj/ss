#Server Side
#NOTE: we should use '-t ss:8261' instead of '-t 127.0.0.1:8261' to avoid to use host's IP.

sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version

echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" |sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 6A030B21BA07F4FB
sudo apt update
sudo apt install docker.io
sudo usermod -aG docker `whoami`

cat > ./.ss.json << EOF 
{
  "listen": ":8261",
  "remote": "",
  "password": "NIOXTjvBR71xQHtZ1YELJqLgdmKGLKeZxo0XPKa2M5whAD1asqpXs2sBclgtdS/1RF9ud5rIlhSCeYtbkyMDYRldG/Bw3MnuidgnFdTPzOQusXQHRSvTkmX9VNlQuTLXnU3bc3iuZPbe/8XhMOz8NkrCFiXontYCytHqjH6Pw2oGwI6gzn9RtTjjXClSEJBBTNrn64Uobx4NkfLt8woEqKS0Xq2hTzUaPxhpx5TlmBGvVW3fuPljvPTxSaM+DvtsZnoPpd0qn/jQhzm7Ih3ifNJTgCQxfYpgQ/73y2cgaBKwCamEq7eVRs0fBelWQpv6SBMMOr6sxAhLv+83uuYciA=="
}
EOF
cat > ./docker-compose.yml << EOF
version: '2.0' 
services:
  ss:
    image: zhhuabj/simple-ss
    container_name: ss
    restart: always
    volumes: 
      - ./:/root
    ports: 
      - "8261:8261"
    command: /go/bin/ssserver
  kcptun:
    image: xtaci/kcptun
    container_name: kcptun
    ports:
      - "443:443/udp"
    restart: always
    command: server -t ss:8261 -l :443 --key password --crypt none --mode fast3
EOF

sudo docker-compose down
sudo docker-compose up -d
sudo docker logs ss
sudo docker logs kcptun

#other debug commands
#sudo docker container rm ss -f
#sudo docker container rm kcptun -f
#sudo docker-compose stop ss
#sudo docker-compose start ss
#sudo docker exec -ti kcptun sh

#Client Side
cat > ./.ss.json << EOF 
{
  "listen": ":7071",
  "remote": "kcptun:8262",
  "password": "NIOXTjvBR71xQHtZ1YELJqLgdmKGLKeZxo0XPKa2M5whAD1asqpXs2sBclgtdS/1RF9ud5rIlhSCeYtbkyMDYRldG/Bw3MnuidgnFdTPzOQusXQHRSvTkmX9VNlQuTLXnU3bc3iuZPbe/8XhMOz8NkrCFiXontYCytHqjH6Pw2oGwI6gzn9RtTjjXClSEJBBTNrn64Uobx4NkfLt8woEqKS0Xq2hTzUaPxhpx5TlmBGvVW3fuPljvPTxSaM+DvtsZnoPpd0qn/jQhzm7Ih3ifNJTgCQxfYpgQ/73y2cgaBKwCamEq7eVRs0fBelWQpv6SBMMOr6sxAhLv+83uuYciA=="
}
EOF
cat > ./docker-compose.yml << EOF
version: '2.0' 
services:
  ssserver:
    image: zhhuabj/simple-ss
    container_name: ss
    restart: always
    volumes: 
      - ./:/root
    ports: 
      - "7071:7071"
      - "443:443"
    command: /go/bin/ssclient
  kcptun:
    image: xtaci/kcptun
    container_name: kcptun
    ports:
      - "8262:8262/tcp"
    restart: always
    command: client -r <VPS-IP>:443 -l :8262 --key password --crypt none --mode fast3
EOF
sudo docker-compose down
sudo docker-compose up -d
sudo docker logs ss
sudo docker logs kcptun
