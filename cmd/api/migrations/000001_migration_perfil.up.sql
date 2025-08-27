CREATE TABLE IF NOT EXISTS tb_user (
    uuid UUID primary key,
    username varchar(200) not null,
    nickname varchar(200) not null,
    pass_word varchar(100) not null,
    email varchar(200)
);
