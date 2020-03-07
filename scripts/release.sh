#!/bin/bash

# NOTE: Be sure to replace all instances of 167.99.165.79
# with your own domain or IP address.

# Change to the directory with our code that we plan to work from
cd "$GOPATH/src/chefhub.pw"

echo "==== Releasing chefhub.pw ===="
echo "  Deleting the local binary if it exists (so it isn't uploaded)..."
rm chefhub.pw
echo "  Done!"

echo "  Deleting existing code..."
ssh root@167.99.165.79 "rm -rf /root/go/src/chefhub.pw"
echo "  Code deleted successfully!"

echo "  Uploading code..."
rsync -avr --exclude '.git/*' --exclude 'tmp/*' \
  --exclude 'images/*' ./ \
  root@167.99.165.79:/root/go/src/chefhub.pw/
echo "  Code uploaded successfully!"

echo "  Go getting deps..."
ssh root@167.99.165.79 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get golang.org/x/crypto/bcrypt"
ssh root@167.99.165.79 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/gorilla/mux"
ssh root@167.99.165.79 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/gorilla/schema"
ssh root@167.99.165.79 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/lib/pq"
ssh root@167.99.165.79 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/jinzhu/gorm"
ssh root@167.99.165.79 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/gorilla/csrf"

echo "  Building the code on remote server..."
ssh root@167.99.165.79 'export GOPATH=/root/go; \
  cd /root/app; \
  /usr/local/go/bin/go build -o ./server \
    $GOPATH/src/chefhub.pw/*.go'
echo "  Code built successfully!"

echo "  Moving assets..."
ssh root@167.99.165.79 "cd /root/app; \
  cp -R /root/go/src/chefhub.pw/assets ."
echo "  Assets moved successfully!"

echo "  Moving views..."
ssh root@167.99.165.79 "cd /root/app; \
  cp -R /root/go/src/chefhub.pw/views ."
echo "  Views moved successfully!"

echo "  Moving Caddyfile..."
ssh root@167.99.165.79 "cd /root/app; \
  cp /root/go/src/chefhub.pw/Caddyfile ."
echo "  Views moved successfully!"

echo "  Restarting the server..."
ssh root@167.99.165.79 "sudo service chefhub.pw restart"
echo "  Server restarted successfully!"

echo "  Restarting Caddy server..."
ssh root@167.99.165.79 "sudo service caddy restart"
echo "  Caddy restarted successfully!"

echo "==== Done releasing chefhub.pw ===="