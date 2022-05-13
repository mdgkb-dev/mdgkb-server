CREATE TABLE regalias (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    doctor_id uuid REFERENCES doctors (id),
    head_id uuid REFERENCES heads (id)
);