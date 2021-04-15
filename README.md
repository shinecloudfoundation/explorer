# explorer
Blockchain explorer for Shinecloudnet

# 部署Shinecloudnet主网区块链浏览器

## 环境准备

### 安装nginx

```shell script
sudo apt-get update
sudo apt-get install nginx make
```
添加nginx的配置:
```shell script
sudo cp default.conf /etc/nginx/conf.d
```

### 安装docker

```shell script
sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo usermod -aG docker $USER
sudo service docker restart
```

### 下载mongo镜像

```shell script
sudo docker pull mongo
```

### 安装golang

```shell script
wget https://golang.org/dl/go1.13.14.linux-amd64.tar.gz
tar -xf go1.13.14.linux-amd64.tar.gz
sudo mv go /usr/local
mkdir go/bin -p
```

### 安装nodejs

```shell script
wget https://nodejs.org/dist/v12.18.3/node-v12.18.3-linux-x64.tar.xz
tar -xf node-v12.18.3-linux-x64.tar.xz
mv node-v12.18.3-linux-x64 node
sudo mv node /usr/local
```

### 添加环境变量

编辑`$HOME/.bashrc`，添加如下配置

```shell script
export GOPATH=/$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/node/bin:$GOPATH/bin:$GOROOT/bin
```

重新加载环境变量
```shell script
source $HOME/.bashrc
```

### 安装yarn

```shell script
npm install -g yarn
```

## 部署服务

### 启动mongo数据库

1. 启动mongo容器

    ```shell script
    sudo docker run -p 0.0.0.0:27017:27017/tcp -d --name shinecloudnet-mongo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo
    ```

2. 登陆mongo shell

    ```shell script
    mongo -u mongoadmin -p secret --authenticationDatabase admin shinecloudnet-explorer
    ```

3. 在mongo shell里面创建mongo用户

    ```
    use sync-shinecloudnet
    db.createUser({
        user:"shinecloudnet",
        pwd:"shinecloudnetpassword",
        roles:[
            {role:"readWrite",db:"shinecloudnet-explorer"}
        ],
        mechanisms:[
        "SCRAM-SHA-1"
        ]
    })
    ```

4. 在mongo shell中执行`database.sql`里面的脚本

### 启动LCD

下载shinescloudnet的binary:

```shell script
git clone https://github.com/shinecloudfoundation/shinecloudnet-binary.git
```

执行`startlcd.sh`来启动lcd

### 启动sync服务

下载sync服务的代码
```shell script
git clone https://github.com/shinecloudfoundation/shinecloudnet-sync.git
```

编译并启动
```shell script
cd $HOME/shinecloudnet-sync
make build
./start.sh
```

### 启动浏览器

下载浏览器代码:
```shell script
git clone https://github.com/shinecloudfoundation/explorer.git
```

#### 启动后端

```shell script
cd $HOME/explorer/backend
make build
./start.sh
```

#### 启动前端

```shell script
cd $HOME/explorer/frontend
yarn install
yarn build
./refresh.sh
```

