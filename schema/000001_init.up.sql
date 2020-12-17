CREATE TABLE users
(
    user_id       serial   not null unique,
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
    email               varchar(255) not null unique,
    password_hash       varchar(255) not null,
    company_name        varchar(255) not null unique,
    director_full_name  varchar(255) not null,
    company_phone       varchar(255) not null unique,
    company_city        varchar(255) not null,
    company_street      varchar(255) not null,
    company_home_number varchar(10)  not null,
    company_flat        varchar(10)  not null
);

CREATE TABLE volume_data
(
    id            serial       not null unique,
    user_id       integer      not null,
    el_volume     varchar(255) default ('null'),
    gas_volume    varchar(255) default ('null'),
    hot_w_volume  varchar(255) default ('null'),
    cold_w_volume varchar(255) default ('null'),
    date_full     varchar(255) not null,
    date_year     varchar(255) not null,
    date_month    varchar(255) not null,
    date_day      varchar(255) not null
);
CREATE TABLE notifications
(
    id          serial      not null unique,
    company_id  integer     not null,
    article     varchar     not null,
    description varchar     not null,
    date_full   varchar(20) not null
)
