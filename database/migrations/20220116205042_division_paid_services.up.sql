CREATE TABLE division_paid_services
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    paid_service_id uuid  REFERENCES paid_services (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    division_id uuid  REFERENCES divisions (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL
);