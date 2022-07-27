CREATE TABLE doctor_paid_services
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    doctor_id uuid  REFERENCES doctors (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    paid_service_id uuid  REFERENCES paid_services (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);