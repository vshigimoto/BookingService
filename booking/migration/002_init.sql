CREATE TABLE IF NOT EXISTS "hotel"
(
    "id"             serial PRIMARY KEY,
    "phone"           VARCHAR(50)         NOT NULL,
    "address"         VARCHAR(255)        NOT NULL,
    "category"        VARCHAR(255)        NOT NULL,
    "rating"          float(25)           NOT NULL
);

insert into hotel (id, phone, address, category, rating) values (1, '753-754-7192', '08 Novick Place', 'comfort+', 3.1);
insert into hotel (id, phone, address, category, rating) values (2, '284-935-4557', '8479 Del Sol Court', 'comfort', 1.6);
insert into hotel (id, phone, address, category, rating) values (3, '956-186-1051', '682 Novick Street', 'comfort+', 1.5);
insert into hotel (id, phone, address, category, rating) values (4, '200-163-5377', '3316 Southridge Circle', 'comfort', 3.1);
insert into hotel (id, phone, address, category, rating) values (5, '159-682-3895', '44009 Prairieview Circle', 'business', 1.5);
insert into hotel (id, phone, address, category, rating) values (6, '794-486-3415', '73 Charing Cross Place', 'business', 2.5);
insert into hotel (id, phone, address, category, rating) values (7, '658-990-7948', '7 Eggendart Alley', 'comfort+', 2.5);
insert into hotel (id, phone, address, category, rating) values (8, '960-363-0836', '69001 Blaine Drive', 'econom', 1.5);
insert into hotel (id, phone, address, category, rating) values (9, '770-421-6826', '1470 Jay Parkway', 'comfort+', 3.5);
insert into hotel (id, phone, address, category, rating) values (10, '875-982-6733', '614 Upham Drive', 'comfort+', 2.7);
insert into hotel (id, phone, address, category, rating) values (11, '492-461-7671', '1417 Kingsford Pass', 'business', 1.5);
insert into hotel (id, phone, address, category, rating) values (12, '798-174-7196', '481 Nelson Junction', 'comfort+', 3.5);
insert into hotel (id, phone, address, category, rating) values (13, '121-472-9330', '37481 Kropf Avenue', 'business', 4.2);
insert into hotel (id, phone, address, category, rating) values (14, '214-839-0965', '57 Raven Drive', 'luxury', 2.0);
insert into hotel (id, phone, address, category, rating) values (15, '772-917-6597', '09 Jackson Hill', 'comfort', 1.6);
insert into hotel (id, phone, address, category, rating) values (16, '527-245-3755', '3848 Brickson Park Center', 'business', 3.9);
insert into hotel (id, phone, address, category, rating) values (17, '202-496-6963', '24904 Everett Avenue', 'econom', 4.9);
insert into hotel (id, phone, address, category, rating) values (18, '827-736-8574', '6862 Pawling Plaza', 'comfort', 1.6);
insert into hotel (id, phone, address, category, rating) values (19, '834-562-0795', '55643 Bay Trail', 'business', 4.0);
insert into hotel (id, phone, address, category, rating) values (20, '237-332-5896', '3 Hoard Hill', 'comfort+', 4.0);


CREATE TABLE IF NOT EXISTS "bookcalendar"
(
    "id"                      serial PRIMARY KEY,
    "hotel_id"                int   NOT NULL,
    "room_count"              int   NOT NULL
);
insert into bookcalendar (id, hotel_id, room_count) values (1, 1, 216);
insert into bookcalendar (id, hotel_id, room_count) values (2, 2, 282);
insert into bookcalendar (id, hotel_id, room_count) values (3, 3, 419);
insert into bookcalendar (id, hotel_id, room_count) values (4, 4, 44);
insert into bookcalendar (id, hotel_id, room_count) values (5, 5, 152);
insert into bookcalendar (id, hotel_id, room_count) values (6, 6, 378);
insert into bookcalendar (id, hotel_id, room_count) values (7, 7, 221);
insert into bookcalendar (id, hotel_id, room_count) values (8, 8, 445);
insert into bookcalendar (id, hotel_id, room_count) values (9, 9, 144);
insert into bookcalendar (id, hotel_id, room_count) values (10, 10, 470);
insert into bookcalendar (id, hotel_id, room_count) values (11, 11, 257);
insert into bookcalendar (id, hotel_id, room_count) values (12, 12, 250);
insert into bookcalendar (id, hotel_id, room_count) values (13, 13, 295);
insert into bookcalendar (id, hotel_id, room_count) values (14, 14, 34);
insert into bookcalendar (id, hotel_id, room_count) values (15, 15, 57);
insert into bookcalendar (id, hotel_id, room_count) values (16, 16, 395);
insert into bookcalendar (id, hotel_id, room_count) values (17, 17, 446);
insert into bookcalendar (id, hotel_id, room_count) values (18, 18, 213);
insert into bookcalendar (id, hotel_id, room_count) values (19, 19, 234);
insert into bookcalendar (id, hotel_id, room_count) values (20, 20, 74);
