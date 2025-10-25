# urB00ks -- 线上图书推荐

## 这个项目本质是一个考核用的项目，但是还是先完整的把他的前后端搭建出来，并上传至Github，以便之后学习

### 文档会在后续进行完善 ...

## Deploy / Development Preparation
**硬件要求**
- 保证至少 **5GB** 的存储空间

**软件要求**
- 使用数据库（Dev）  `mariadb`
- 虚拟环境    `docker`
- 部署系统    `Linux`

> [!WARNING]
> 当您想要对本项目进行再开发的时候，请务必确保您的系统上有 `mariadb` 以方便您的调试；相应数据已经在 `/backend/books.csv` 中给出

### 配置开发环境步骤（仅对数据库方面进行说明）

#### 1，下载mariadb并初始化

先运行下面的命令来下载
```bash
sudo pacman -S mariadb
```

当下载完成之后，按照下面的方式对数据库进行初始化
```bash
sudo mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
```

之后，启动服务
```bash
sudo systemctl enable mariadb
```

之后，运行安全脚本，确保对您的数据库账户设置对应密码，根据脚本提示操作即可：

```bash
sudo mariadb-secure-installation
```

完成后，您可以输入下面的指令进入 `mariadb`
```bash
mariadb -u root -p
```

之后，您将看到这样的界面
```
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 5
Server version: 12.0.2-MariaDB Arch Linux

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]>
```
若您的输出是这样的，说明您成功的配置好了您的 `mariadb`

#### 2，创建对应数据库并修改 `/backend/dao/db.go` 文件的内容

mariadb是mysql在linux上的分支版本，所以您可以按照mysql的操作进行创建数据库的操作

当您创建完您的表之后，为了使您的后端可以连接到数据库中，请跳转到：`/backend/dao/db.go`，您将看到以下常量：
```go
var OfficialDB *gorm.DB
var DSN string
var DevDSN = "fzsgball:123456@tcp(localhost:3306)/official_book_library?charset=utf8mb4&parseTime=True&loc=Local"
```
**下面进行解释**：

**OfficialDB**
- `gorm.DB` 指针，在这个板块中我们用不到他

**DSN**
- 部署后的DSN，针对docker内环境，不用修改

**DevDSN**
- 开发环境的DSN，在您修改了 `isDev` 常量后，将会使用这个DSN，您可以按照配置DSN的方法对其进行配置，下面对其进行简要说明：

|fzsgball|123456|localhost|3306|official_book_library|
|---|---|---|---|---|
|您数据库的用户名称，这个用户应该是专门给这个项目使用的|用户密码|代表本机的标识符，您使用127.0.0.1亦可|端口，mysql或者说mariadb的默认端口都是3306，不必修改|对应数据库的名称|

执行到这里，恭喜您，已经完成了对于数据库相关的配置了！

### 部署步骤

#### 1，确保docker
首先，确保您的设备上有 `docker`，如果没有，请下载

对于ArchLinux系统：
```bash
sudo pacman -S docker
```

之后启用服务：
```bash
sudo systemctl enable docker
```
即可

对于WSL 2：
请先确保您的Windows系统下载了 Docker Desktop
之后在 `Settings > Resources > WSL integration` 中启用对应的WSL即可

#### 2，Composer building
运行下面的命令来进行Composer Building（请注意，在本项目的根目录下进行）：
```bash
docker compose build
```
当生成完成之后，运行下面的命令来运行您的Containers：
```bash
docker compose up
```

之后，如果您没有改变配置的话，您可以在：`localhost:8088` 处访问这个项目

## Other Documents
- [api文档](frontend/api.md)
- [后端文档](backend/README.md)

## Development
### Frontend
在启动urB00ks的前端服务之前，您需要确保您的设备上有pnpm，如果没有，请尝试运行以下命令：
```
npm install -g pnpm
```
如果要启动web前端服务（开发模式），您只需要切换到当前目录并：
```
pnpm dev
```
如果没有安装依赖，请：
```
npm install
```
### Backend
运行之前，请确保您的系统上存在 **mariadb**

因为项目在docker上运行，所以在进行开发时，需要对 `main.go` 内的 `isDev` 进行更改：
```go
var isDev = true
```
当 `isDev` 为true的时候就说明打开了开发模式，此时再运行 `go run main.go` 就会出现正常的输出