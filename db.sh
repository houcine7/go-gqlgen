printf "Creating mysql container...\n"
docker run --name mysql1 -e MYSQL_ROOT_PASSWORD=123456 -d -p 3307:3306 mysql:5.7


sleep 20

# create a db 
printf "Creating database...\n"
docker exec -it mysql1 mysql -uroot -p123456 -e "create database go_graphql;"

# create User table
printf "Creating User table...\n"
docker exec -it mysql1 mysql -uroot -p123456 -e "use go_graphql; create table if not exists User (ID int not null auto_increment, Username varchar(20) not null, Password varchar(255), primary key(ID));"

# create Link table
printf "Creating User table...\n"
docker exec -it mysql1 mysql -uroot -p123456 -e "use go_graphql; create table if not exists Link (ID int not null auto_increment, Title varchar(255) not null, Address varchar(255),UserID int,primary key(ID), foreign key (UserID) references User(ID))"

