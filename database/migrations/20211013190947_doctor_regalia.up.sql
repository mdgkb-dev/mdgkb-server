CREATE TABLE doctor_regalias (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    doctor_id uuid REFERENCES doctors (id)
);