use banking;
alter table `banking`.`users`
add column role ENUM("ADMIN", "USER") default "USER";