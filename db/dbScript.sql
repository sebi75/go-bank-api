create database if not exists banking;
use banking;

drop table if EXISTS customer;
create table `customers` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `date_of_birth` date not null,
    `city` varchar(100) not null,
    `zipcode` varchar(10) not null,
    `status` ENUM('1', '0') not null DEFAULT '1',
    PRIMARY KEY (`id`)
) engine=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

insert into `customers` VALUES
    (2000, "Sebastian", "1978-01-01", "Berlin", "10115", 1),
    (2001, "Klaus", "1978-01-01", "Berlin", "10115", 1),
    (2002, "Peter", "1978-01-01", "Berlin", "10115", 1),
    (2003, "Hans", "1978-01-01", "Berlin", "10115", 1),
    (2004, "Karl", "1978-01-01", "Berlin", "10115", 1),
    (2005, "Karl", "1978-01-01", "Berlin", "10115", 1);

drop table if EXISTS accounts;
create table `accounts` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) not null,
    `opening_date` datetime not null default current_timestamp,
    `account_type` varchar(10) not null,
    `pin` varchar(10) not null,
    `status` ENUM('1', '0') not null default '1'
    PRIMARY KEY (`id`),
    FOREIGN KEY (`customer_id`) REFERENCES `customers`(`id`) on delete cascade
) engine=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

insert into `accounts` VALUES
    (2000, 2000, "2019-01-01 00:00:00", "savings", "1234", 1),
    (2001, 2001, "2019-01-01 00:00:00", "savings", "1234", 1),
    (2002, 2002, "2019-01-01 00:00:00", "savings", "1234", 1),
    (2003, 2003, "2019-01-01 00:00:00", "savings", "1234", 1),
    (2004, 2004, "2019-01-01 00:00:00", "savings", "1234", 1),
    (2005, 2005, "2019-01-01 00:00:00", "savings", "1234", 1);

drop table if EXISTS transactions;
create table transactions (
    id int(11) not null auto_increment,
    account_id int(11) not null,
    amount decimal(10,2) not null,
    transaction_type varchar(10) not null,
    transaction_date datetime not null default current_timestamp,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`) on delete cascade
) engine=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
