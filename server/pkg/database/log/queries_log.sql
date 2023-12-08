CREATE TABLE users (
	id varchar(255) primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null unique,
    phone_number varchar(255) not null unique,
    title varchar(255),
    role varchar(255) not null,
	password varchar(255)
);

Create table registration (
	id int primary key auto_increment,
    registration_time varchar(255) not null,
    supervisor varchar(255) not null,
    user_id varchar(255) not null, 
    room_id varchar(255) not null,
    start_day timestamp not null,
    end_day timestamp not null
);

create table organization (
	id varchar(255) primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    head varchar(255) not null,
    phone_number varchar(255) not null,
    website varchar(255) not null
);

create table room (
	id varchar(255) primary key,
    office_id varchar(255) not null,
    name varchar(255) not null,
    description varchar(255) not null
);

create table office (
	id varchar(255) primary key,
    organization_id varchar(255) not null,
    name varchar(255) not null,
    address varchar(255) not null,
    manager varchar(255) not null,
    phone_number varchar(255) not null
);

create table history (
	id int primary key auto_increment,
    registration_id varchar(255) not null,
    activity_type varchar(255) not null,
    time timestamp not null,
    admin_id varchar(255) not null
);

use hpcc_checkin;

select * from office;

INSERT INTO office (id, organization_id, name, address, manager, phone_number) VALUES ("1", "1", "ist_lab", "ltk", "hdhd", "124");


insert into organization (id, name, email, head, phone_number, website) values ("1", "ist", "ist@hcmut.edu.vn", "dhd", "124124", "www.ist.com");
insert into room (id, office_id, name, description) values ("1", "1", "ist lab", "...");

select * from registration where id = 1;

SELECT registration.id as id,  registration_time, user_id as student_id, start_day, end_day, first_name, last_name, email, phone_number, org_name, office_name, room_name, supervisor
FROM registration 
	join users on users.id = registration.user_id
    join (
		select organization.name as org_name, office.name as office_name, room.name as room_name, office.id 
		from room 
			join office on room.office_id=office.id 
			join organization on organization.id=office.organization_id
	) as rooms on rooms.id = registration.room_id
where user_id = 123124
ORDER BY registration.id desc LIMIT 1;

select * from organization﻿;

select organization.name, office.name, room.name, office.id 
from room 
	join office on room.office_id=office.id 
    join organization on organization.id=office.organization_id;
    
select * from history where registration_id = 1 order by id limit 1 ;

update registration 
set registration_time = "11h - 12h", room_id = "1", start_day = "2013-08-05 18:19:03", end_day = "2013-08-05 18:19:03"
where user_id = "1231245" and room_id="1";

select history.id, activity_type, time, admin_id, registration_time, supervisor, 
	user_id, start_day, end_day, first_name, last_name, users.email, users.phone_number, 
	room.name as room_name, office.name as office_name, organization.name as org_name  
from history
	join registration on registration_id = registration.id
    join users on users.id = user_id
    join room on room.id = room_id
    join office on office.id = office_id
    join organization on organization.id = organization_id
order by history.id desc limit 10;

update users
set password = "123", role="admin"
where id = "123124";

delete from registration where id = "1";

UPDATE registration
SET registration_time = "11h-12h", supervisor = "DHD", start_day = "2013-08-05 18:19:03", end_day = "2013-08-05 18:19:03"
WHERE user_id = "12345678" AND room_id = "1";