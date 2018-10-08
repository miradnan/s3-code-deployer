# S3 Code Deployer

[![CircleCI](https://circleci.com/gh/miradnan/s3-code-deployer/tree/master.svg?style=svg)](https://circleci.com/gh/miradnan/s3-code-deployer/tree/master)

### Why I built this?
I was working with an On-Prem server and I had to install AWS CodeDeploy Agent to get the CI & CD Pipeline.
CodeDeploy uses ruby version 2.x. Ubuntu 16.04 and above, ships with ruby 2.3 as default and CodeDeploy just doesn't work.
So, I thought why not write it in GoLang. Go is best suited for this job. No dependencies at all. Just a binary file that runs as a service just.



```
################################################################
####################### DEBIAN PACKAGE #########################
################################################################
```

##### Import GPG Public Key
```
wget - no-cache -O - https://api.bintray.com/users/miradnanali/keys/gpg/public.key | sudo apt-key add -
```

### Add Source List

#### Ubuntu

##### Bionic Beaver (18.04)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian bionic main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Artful (17.10)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian artful main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Zesty (17.04)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian zesty main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Xenial (16.04)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian xenial main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Trusty (14.04)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian trusty main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

#### Debian

##### Buster (10)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian buster main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Stretch (9)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian stretch main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Jessie (8)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian jessie main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

##### Wheezy (7)
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian wheezy main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

### Installation Command for Debian / Ubuntu
```
$ apt-get install s3-code-deployer
```
#### See create config.yml file below




```
################################################################
########################## RPM PACKAGE #########################
################################################################
```

##### Import GPG Public Key
```
rpm --import https://api.bintray.com/users/miradnanali/keys/gpg/public.key
```

#### Add Repo
```
sudo echo "
[s3-code-deployer]
name=s3-code-deployer
baseurl=https://api.bintray.com/content/miradnanali/S3CodeDeployerRPM
gpgkey=https://api.bintray.com/users/miradnanali/keys/gpg/public.key
gpgcheck=0
enabled=1
repo_gpgcheck=1
" > /etc/yum.repos.d/s3-code-deployer.repo
```

#### Update Yum cache:
```
sudo yum update
```

#### Install
```
sudo yum install -y s3-code-deployer
```


### Create a `config.yml` file in `/etc/s3-code-deployer/config.yml`. Below is a Sample config.yml.
```
revision_check_duration: 10 // in minutes

aws:
  accessKey: YOUR_AWS_ACCESS_KEY
  secretKey: YOUR_AWS_SECRET_KEY
  bucket: YOUR_AWS_DEPLOYMENT_BUCKET
  region: YOUR_AWS_REGION

deployments:

- application: staging.example.com
  environment: staging
  destination: /var/www/html/staging.example.com
  s3_revision_file: example.com/staging.tar.gz

- application: www.example.com
  environment: production
  destination: /var/www/html/www.example.com
  s3_revision_file: example.com/prod.tar.gz
```