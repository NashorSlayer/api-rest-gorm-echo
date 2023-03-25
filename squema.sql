

--min 10 del video , hacer base de datos 

create table USERS(
id serial ,
email text not null,
name text not null,
password text not null,
primary key(id)
);

create table PRODUCTS(
    id serial,
    name text not null,
    description text not null,
    price float not null,
    created_by int not null,
    primary key(id),
    foreign key(created_by) references USERS(id)
);

create table ROLES(
    id serial,
    name text not null,
    primary key(id)
);

create table USER_ROLES(
    id serial,
    user_id int not null,
    role_id int not null,
    primary key(id),
    foreign key (user_id) references USERS(id),
    foreign key (role_id) references ROLES(id)
);



insert into ROLES (name) values 
('admin'),
('seller'),
('customer')


insert into USER_ROLES(user_id,role_id) values 
(7,1);