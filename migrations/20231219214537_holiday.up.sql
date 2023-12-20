create table holiday_forms(

    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    surname varchar,
    patronymic varchar,

parent_name varchar,
parent_surname varchar,
parent_patronymic varchar,
created_at timestamp DEFAULT current_timestamp,

representative varchar,
email varchar,
phone varchar,

dance varchar,
song varchar,
music varchar,
custom_show varchar,
needing varchar,


color varchar,
hobby varchar,
happy varchar,
place varchar


)

