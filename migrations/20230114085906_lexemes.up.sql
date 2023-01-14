create extension rum;

create table if not exists lexemes
(
    lexeme text constraint lexemes_uniq_idx unique
);

insert into lexemes (lexeme)
select word from ts_stat($$
    select search_column from search_elements
$$);