create or replace procedure update_search_elements()
    language sql
as
$$
truncate search_elements;
insert into search_elements(key, label, description, search_column, value)
SELECT
    'doctor', full_name,description,
    setweight(to_tsvector('english', surname), 'A') ||
    setweight(to_tsvector('russian', full_name), 'B') ||
    setweight(to_tsvector('russian', position), 'C')
                AS search_column,
    id::varchar as value
FROM
    doctors_view
union all
SELECT
    'division' ,name,info  ,
    setweight(to_tsvector('russian', name), 'A') ||
    setweight(to_tsvector('russian', info), 'B')
                AS search_column,
    id::varchar as value
FROM
    divisions
UNION ALL
SELECT
    'paidService', name, type,
    setweight(to_tsvector('russian', name), 'D') ||
    setweight(to_tsvector('russian', type), 'D')
                AS search_column,
    id::varchar as value
FROM
    paid_services ;
$$;

create or replace procedure update_lexemes()
    language sql
as
$_$
truncate lexemes;

insert into lexemes (lexeme)
select word from ts_stat($$select search_column from search_elements$$);
$_$



