truncate search_elements;

insert into search_elements(key, label, description, search_column, value)
SELECT
    'doctor', full_name,description,
    setweight(to_tsvector('russian', full_name), 'A') ||
    setweight(to_tsvector('russian', position), 'B')
        AS search_column,
    slug as value
FROM
    doctors_view
UNION ALL
SELECT
    'division' ,name,info  ,
    setweight(to_tsvector('russian', name), 'A') ||
    setweight(to_tsvector('russian', info), 'B')
        AS search_column,
    slug as value
FROM
    divisions
        UNION ALL
SELECT
    'paidService', name, type,
    setweight(to_tsvector('russian', name), 'A') ||
    setweight(to_tsvector('russian', type), 'B')
        AS search_column,
    id::varchar as value
FROM
    paid_services ;