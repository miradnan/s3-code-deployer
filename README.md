# S3 Code Deployer

[ ![Download](null/packages/miradnanali/S3CodeDeployer/apt/images/download.svg?version=1.1) ](https://bintray.com/miradnanali/S3CodeDeployer/apt/1.1/link)

### Why I built this?
I was working with an On-Prem server and I had to install AWS CodeDeploy Agent to get the CI & CD Pipeline.
CodeDeploy uses ruby version 2.x. Ubuntu 16.04 ships with ruby 2.3 as default and CodeDeploy just doesn't work.
So, I thought why not write it in GoLang. Go is best suited for this job. No dependencies at all.



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
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian artful main" > /etc/apt/sources.list.d/s3-codedeployer.list
```
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian zesty main" > /etc/apt/sources.list.d/s3-codedeployer.list
```
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian xenial main" > /etc/apt/sources.list.d/s3-codedeployer.list
```
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian trusty main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

#### Debian
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian buster main" > /etc/apt/sources.list.d/s3-codedeployer.list
```
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian stretch main" > /etc/apt/sources.list.d/s3-codedeployer.list
```
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian jessie main" > /etc/apt/sources.list.d/s3-codedeployer.list
```
```
echo "deb [arch=amd64] https://dl.bintray.com/miradnanali/S3CodeDeployerDebian wheezy main" > /etc/apt/sources.list.d/s3-codedeployer.list
```

### Installation Command for Debian / Ubuntu
```
$ apt-get install s3-code-deployer
```

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
sudo echo "" > /etc/yum.repos.d/s3-code-deployer.repo

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