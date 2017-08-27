echo 'Starting!'
pwd
sudo cp -v bash_profile ~/.bash_profile
cat ~/.bash_profile
sudo cp -v path.sh /etc/profile.d/path.sh
cat /etc/profile.d/path.sh
sudo curl -LO https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.7.linux-amd64.tar.gz
sudo yum -y install git
mkdir -p ~/projects/{bin,pkg,src}

source /etc/profile && source ~/.bash_profile
echo 'Done!!'