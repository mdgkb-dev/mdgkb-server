CREATE TABLE path_permissions (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    resource varchar unique,
    guest_allow boolean
);