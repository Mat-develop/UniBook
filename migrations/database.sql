CREATE DATABASE IF NOT EXISTS unibook;
USE unibook;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key unique,
    name varchar(50) not null, 
    nick varchar(50) not null unique, 
    email varchar(50) not null unique,
    image_url varchar(255),
    password varchar(100) not null, 
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

{
    "name": "Matheus",
    "nick": "Matheusiznho",
    "email": "something@gmail.com",
    "password": "123456",
    "image_url": "none"
}

CREATE TABLE IF NOT EXISTS community(
    id int auto_increment primary key unique,
    name varchar(50) not null, 
    image_url varchar(255),
    created_at timestamp default current_timestamp()
)

CREATE TABLE IF NOT EXISTS community_followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE, 

    community_id int not null,
    FOREIGN KEY (community_id)
    REFERENCES community(id)
    ON DELETE CASCADE,

    primary key(user_id, community_id)
) ENGINE=INNODB;

// 1, 2
// 1, 3
// 1, 4 