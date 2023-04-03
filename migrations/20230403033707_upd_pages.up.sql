alter table
    pages
add
    role_id uuid REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL;