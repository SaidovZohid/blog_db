CREATE TABLE "email_ver" (
    id serial primary key not null,
    username varchar(50) not null,
    email varchar(255) not null,
    ver_code int not null
);