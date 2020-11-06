CREATE TABLE users
(
    user_id       serial       not null unique,
    surname       varchar(255) not null,
    name          varchar(255) not null,
    patronymic    varchar(255) not null,
    full_name     varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    company_id    integer default 0
);
CREATE TABLE address
(
    address_id  serial       not null unique,
    user_id     integer      not null,
    city        varchar(255) not null,
    street      varchar(255) not null,
    home_number varchar(10)  not null,
    flat        varchar(10)  not null
);
CREATE TABLE company
(
    company_id          serial       not null unique,
    company_name        varchar(255) not null unique,
    director_full_name  varchar(255) not null,
    company_phone       varchar(255) not null unique,
    company_city        varchar(255) not null,
    company_street      varchar(255) not null,
    company_home_number varchar(10)  not null,
    company_flat        varchar(10)  not null
);

CREATE TABLE electricity_list
(
    id         serial       not null unique,
    volume     varchar(255) not null,
    date_full  varchar(255) not null,
    date_month varchar(255) not null,
    date_year  varchar(255) not null
);
CREATE TABLE users_electricity
(
    id             serial                                                 not null unique,
    user_id        int references users (user_id) on delete cascade       not null,
    electricity_id int references electricity_list (id) on delete cascade not null
);

CREATE TABLE cold_water_list
(
    id         serial       not null unique,
    volume     varchar(255) not null,
    date_full  varchar(255) not null,
    date_month varchar(255) not null,
    date_year  varchar(255) not null
);
CREATE TABLE users_cold
(
    id            serial                                                not null unique,
    user_id       int references users (user_id) on delete cascade      not null,
    cold_water_id int references cold_water_list (id) on delete cascade not null
);

CREATE TABLE hot_water_list
(
    id         serial       not null unique,
    volume     varchar(255) not null,
    date_full  varchar(255) not null,
    date_month varchar(255) not null,
    date_year  varchar(255) not null
);
CREATE TABLE users_hot
(
    id            serial                                               not null unique,
    user_id       int references users (user_id) on delete cascade     not null,
    cold_water_id int references hot_water_list (id) on delete cascade not null
);

CREATE TABLE gas_list
(
    id         serial       not null unique,
    volume     varchar(255) not null,
    date_full  varchar(255) not null,
    date_month varchar(255) not null,
    date_year  varchar(255) not null
);
CREATE TABLE users_gas
(
    id            serial                                           not null unique,
    user_id       int references users (user_id) on delete cascade not null,
    cold_water_id int references gas_list (id) on delete cascade   not null
);