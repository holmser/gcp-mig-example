sudo apt-get install wget

wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz

sudo rm tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin