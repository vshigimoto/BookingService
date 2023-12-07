CREATE TABLE IF NOT EXISTS "hotel"
(
    "id"              serial PRIMARY KEY,
    "name"            VARCHAR(255)        NOT NULL,
    "phone"           VARCHAR(50)         NOT NULL,
    "address"         VARCHAR(255)        NOT NULL,
    "category"        VARCHAR(255)        NOT NULL,
    "rating"          float(25)           NOT NULL
);

insert into hotel (name, phone, address, category, rating) values ('Renner Inc', '+33 923 653 4484', '17 Victoria Way', 'comfort', 2.5);
insert into hotel (name, phone, address, category, rating) values ('Stehr Inc', '+33 588 406 2595', '3842 Gateway Junction', 'econom', 3.0);
insert into hotel (name, phone, address, category, rating) values ('Murphy LLC', '+63 750 673 4561', '83441 Northland Point', 'business', 2.9);
insert into hotel (name, phone, address, category, rating) values ('Zulauf Inc', '+86 389 584 0693', '5 Oxford Crossing', 'comfort+', 2.8);
insert into hotel (name, phone, address, category, rating) values ('Gutmann, Kub and Bosco', '+63 771 195 0517', '3110 Golden Leaf Park', 'luxury', 2.2);
insert into hotel (name, phone, address, category, rating) values ('Lindgren-Johns', '+86 561 609 8714', '54 Summit Junction', 'business', 2.9);
insert into hotel (name, phone, address, category, rating) values ('Effertz-Turcotte', '+420 645 964 0750', '1282 Maple Street', 'comfort+', 1.5);
insert into hotel (name, phone, address, category, rating) values ('Erdman and Sons', '+54 803 478 2629', '6402 Del Sol Crossing', 'comfort', 1.1);
insert into hotel (name, phone, address, category, rating) values ('Stamm Group', '+351 407 607 4084', '5796 Pepper Wood Road', 'econom', 3.3);
insert into hotel (name, phone, address, category, rating) values ('Barton, Corwin and Smith', '+504 993 827 0378', '757 Delaware Crossing', 'comfort', 3.7);
insert into hotel (name, phone, address, category, rating) values ('Daugherty, Yost and Wolf', '+30 672 763 3690', '65 Kedzie Center', 'business', 1.5);
insert into hotel (name, phone, address, category, rating) values ('Bartoletti Group', '+86 164 857 3996', '250 Oak Hill', 'comfort+', 4.7);
insert into hotel (name, phone, address, category, rating) values ('Spinka-Leuschke', '+591 721 188 4508', '4092 Golf View Park', 'econom', 2.4);
insert into hotel (name, phone, address, category, rating) values ('Ankunding, Friesen and Dicki', '+33 688 837 2972', '8 Lillian Plaza', 'luxury', 1.8);
insert into hotel (name, phone, address, category, rating) values ('Hammes, O''Connell and Weber', '+86 148 763 5981', '64 Holmberg Road', 'luxury', 3.4);
insert into hotel (name, phone, address, category, rating) values ('Mayer-Heller', '+385 923 561 0053', '8415 Golf Course Alley', 'econom', 3.3);
insert into hotel (name, phone, address, category, rating) values ('Kuhlman, Runolfsson and O''Keefe', '+55 529 969 4390', '8 Bobwhite Junction', 'business', 2.8);
insert into hotel (name, phone, address, category, rating) values ('Kling, Reichert and Hartmann', '+389 810 893 8097', '4 Hovde Place', 'business', 1.4);
insert into hotel (name, phone, address, category, rating) values ('Will-Toy', '+380 855 103 5832', '748 Mitchell Lane', 'comfort+', 2.2);
insert into hotel (name, phone, address, category, rating) values ('Schultz-Skiles', '+86 908 140 1601', '73510 Nancy Lane', 'econom', 3.4);


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

CREATE TABLE IF NOT EXISTS "bookrequest"
(
    "id"                      serial PRIMARY KEY,
    "hotel_id"                int   NOT NULL,
    "code"                    int   NOT NULL
);