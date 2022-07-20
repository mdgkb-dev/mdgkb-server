# mdgkb-server
Сервер портала Морозовской больницы
## Installation
Для запуска проекта необходимо установить ряд пакетов.
Ниже приведен пример порядка установки для Ubuntu 20.04

### 1. Установите Node Version Manager (nvm): 
```
	sudo apt update 
	sudo apt install curl
	curl https://raw.githubusercontent.com/creationix/nvm/master/install.sh | bash
	source ~/.bashrc
```
	и далее, установите node v16.11.0:
```
	nvm install v16.11.0
```
### 2. Установите Visual Studio Code:
```
	sudo apt install software-properties-common apt-transport-https wget
	wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add -
	sudo add-apt-repository "deb [arch=amd64]https://packages.microsoft.com/repos/vscode stable main"

	sudo apt update
	sudo apt install code
```
### 3. Установите и настройте Git:
```
	sudo apt update
	sudo apt install git
	git --version
	git config --global user.name "ваше имя в git"
	git config --global user.email "ваш email в git"
	git config --list
```
### 4. Если вы пользуйтесь программой контроля версий Sublime Merge, то установите ее:
```
	sudo apt update && sudo apt install apt-transport-https -y
	wget -qO - https://download.sublimetext.com/sublimehq-pub.gpg | sudo apt-key add -

	echo "deb https://download.sublimetext.com/ apt/stable/" | sudo tee /etc/apt/sources.list.d/sublime-text.list

	sudo apt update && sudo apt install sublime-merge -y
```
### 5. Установите базу данных PostgreSQL и проверьте ее статус:
```
	sudo apt update
	sudo apt -y install postgresql
	systemctl status postgresql
```
	По умолчанию в Ubuntu не задан пароль для пользователя «postgres». Нужно задать пароль (например, 'postgres'), иначе будут проблемы с подключением к базе данных:
```
	sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"
```
### 6. Установите redis-server:
```
	sudo apt update
	sudo apt install redis-server
```
	Откройте файл redis.conf:
```
	sudo nano /etc/redis/redis.conf
```
	Внутри файла найдите директиву supervised. Она задает систему инициализации для управления сервисом Redis, предоставляя вам больше контроля над ним. По умолчанию директива supervised имеет значение no. Поскольку вы используете Ubuntu, а она использует систему инициализации systemd, укажите здесь ее: 
```    
	supervised systemd
```
	далее перезапустите redis и проверте его статус
```
	sudo systemctl restart redis.service
	sudo systemctl status redis
```
### 7. Установите make:
```
	sudo apt install make
```
### 8. Если вы используете pgadmin4 для контроля базы данных, установите его: 
```
	sudo wget https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add packages_pgadmin_org.pub

	sudo apt-get -y install apt-transport-https lsb-release ca-certificates curl
	sudo curl https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add
	sudo sh -c 'echo "deb https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list'

	sudo apt-get update
	sudo apt install pgadmin4
```
### 9. Установите Goland и проверьте его версию:
```
	sudo apt install golang
	export PATH=$PATH:/opt/go/bin
	go version
```
### 10. Установите Reflex:
```
	go install github.com/cespare/reflex@latest  
```
### 11.Установите  wkhtmltopdf:
```
	sudo apt install ./wkhtmltox_0.12.6-1.focal_amd64.deb
```