create table paid_programs_packages_options
(
    id uuid default uuid_generate_v4() not null primary key ,
    paid_program_package_id uuid references paid_program_packages on delete cascade,
    paid_program_option_id uuid references paid_program_options on delete cascade
);

