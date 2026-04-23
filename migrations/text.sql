-- ALL necessary creations 
CREATE DATABASE IF NOT EXISTS unibook;
USE unibook;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS posts;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key unique,
    name varchar(50) not null, 
    nick varchar(50) not null unique, 
    email varchar(50) not null unique,
    image_url varchar(255),
    password varchar(100) not null, 
    created_at timestamp default current_timestamp()
) 


CREATE TABLE followers(
	user_id int not null, 
    foreign key (user_id)
    references users(id)
	ON DELETE CASCADE, 
    
    follower_id int not null, 
    foreign key (follower_id)
    references users(id)
    on delete cascade,
    
    primary key (user_id, follower_id)
)
    
insert into users(name, nick, email, password)
    values 
    () -- users 

CREATE TABLE post(
    id uuid not null, 
    title varchar(50) not null,
    body varchar(300) not null,

    user_id uuid not null, 
    foreign key (user_id)
    references users(id)
	ON DELETE CASCADE, 

    community_id uuid, 
    foreign key (community_id)
    references communities(id)
    ON DELETE CASCADE, 

    likes int default,  
    created_at timestamp default current_timestamp()   
    )