# S3 Code Deployer

### Why I built this?
I was working with an On-Prem server and I had to install AWS CodeDeploy Agent to get the CI & CD Pipeline.
CodeDeploy uses ruby version 2.x. Ubuntu 16.04 ships with ruby 2.3 as default and CodeDeploy just doesn't work.
So, I thought why not write it in GoLang. Go is best suited for this job. No dependencies.

### Command to execute
```
$ ./s3deployer
```

### Sample YAML Config
```
name: S3CodeDeployer
revision_check_duration: 10
aws:
  accessKey: XASDASDASD
  secretKey: ASDQdijdfh892h2iu34n32
  bucket: deploymentsbucket
  region: ap-south-1
deployments:
- application: staging.example.com
  environment: staging
  destination: /Users/miradnan/workspace/go/src/github.com/miradnan/codeDeployer/appdir/staging
  s3_revision_file: myapp/staging.tar.gz
- application: www.example.com
  environment: production
  destination: /Users/miradnan/workspace/go/src/github.com/miradnan/codeDeployer/appdir/production
  s3_revision_file: myapp/prod.tar.gz
```