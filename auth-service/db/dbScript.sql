drop table if exists users;

create table users (
    id int not null auto_increment,
    username varchar(255) not null,
    customer_id int not null,
    password varchar(255) not null,
    primary key (id),
    foreign key (customer_id) references customers(id)
);

-- insert data in the users table

INSERT INTO users (username, customer_id, password) VALUES
('user2000', 2000, 'password2000'),
('user2001', 2001, 'password2001'),
('user2002', 2002, 'password2002'),
('user2003', 2003, 'password2003'),
('user2004', 2004, 'password2004');
