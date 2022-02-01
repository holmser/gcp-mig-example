# Download Go
curl -OL https://go.dev/dl/go1.17.6.linux-amd64.tar.gz

# Unpack Go binary
sudo rm tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz

# Add Go directory to Path
echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile
source $HOME/.profile

