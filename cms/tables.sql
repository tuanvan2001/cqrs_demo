create table locations (
    id integer primary key auto_increment,
    name varchar(255) not null,
    address varchar(255) not null,
    status integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create table games (
    id integer primary key auto_increment,
    location_id integer not null,
    name varchar(255) not null,
    description text,
    status integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create table news (
    id integer primary key auto_increment,
    title varchar(255) not null,
    content text not null,
    status integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create table settings (
    id integer primary key auto_increment,
    location_id integer not null,
    `key` varchar(255) not null,
    value varchar(5000) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create table banners (
    id integer primary key auto_increment,
    location_id integer not null,
    image varchar(255) not null,
    link varchar(255) not null,
    status integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create index game_location_id on games (location_id);
create index banner_location_id on banners (location_id);
create index settings_location_id on settings (location_id);

-- ERD
alter table games add constraint fk_games_location_id foreign key (location_id) references locations (id) on delete cascade;
alter table settings add constraint fk_settings_location_id foreign key (location_id) references locations (id) on delete cascade;
alter table banners add constraint fk_banners_location_id foreign key (location_id) references locations (id) on delete cascade;