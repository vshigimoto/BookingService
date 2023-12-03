CREATE TABLE IF NOT EXISTS users(
                                    id           int                 PRIMARY KEY,
                                    first_name   VARCHAR(50)         NOT NULL,
    last_name    VARCHAR(255)        NOT NULL,
    phone        VARCHAR(255)        NOT NULL,
    login        VARCHAR(255) UNIQUE NOT NULL,
    password     VARCHAR(255)        NOT NULL,
    );

CREATE TABLE IF NOT EXISTS user_token (
                                          id SERIAL PRIMARY KEY,
                                          token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT user_id_unique UNIQUE (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
    );

CREATE TABLE IF NOT EXISTS card_transaction (
                                                id SERIAL PRIMARY KEY ,
                                                fromCard_id     int not null ,
                                                toCardRequisite varchar not null ,
                                                toCardName      varchar not null ,
                                                toCardSurname   varchar not null ,
                                                money           numeric not null ,
                                                time            timestamp not null,
                                                foreign key (fromCard_id) references card(id),
    );

CREATE TABLE IF NOT EXISTS user_cards(
                                         user_id int not null ,
                                         card_id serial primary key ,
                                         foreign key (user_id) references users(id)
    );


CREATE TABLE IF NOT EXISTS card(
                                   id serial primary key ,
                                   requisite varchar,
                                   exp varchar,
                                   cvc int not null ,
                                   full_name varchar,
                                   foreign key (id) references user_cards(card_id)
);

CREATE TABLE IF NOT EXISTS user_balance(
                                           id serial not null ,
                                           user_id int not null ,
                                           currency varchar,
                                           balance numeric,
                                           foreign key (user_id) references users(id)
    )

CREATE TABLE IF NOT EXISTS p2p_transaction (
                                               id SERIAL PRIMARY KEY ,
                                               from_id int not null ,
                                               to_id int not null ,
                                               money numeric not null ,
                                               message varchar,
                                               time timestamp not null,
                                               foreign key (from_id) references users(id),
    foreign key (to_id) references users(id)
    );





