# Update the APT package index
apt-get update
  
# Install packages to allow apt to use a repository over HTTPS
apt-get install -y apt-transport-https ca-certificates curl software-properties-common
  
# Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
  
# Set up the Docker repository
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  
# Update the APT package index again with the new repository
apt-get update
  
# Install the latest version of Docker CE
apt-get install -y docker-ce
  
# Add vagrant user to the docker group to run docker commands without sudo (optional)
usermod -aG docker vagrant


## Install Golang
sudo snap install go --classic