CREATE TABLE path_permissions_roles
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    path_permission_id  uuid  REFERENCES path_permissions(id) ON UPDATE CASCADE ON DELETE CASCADE,
    role_id  uuid  REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE
);
