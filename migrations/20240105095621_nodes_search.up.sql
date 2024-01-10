alter table divisions add column map_node_name varchar;
alter table buildings add column map_node_name varchar;
alter table gates add column map_node_name varchar;
alter table entrances add column map_node_name varchar;

create view map_objects as
select name, map_node_name, 'division' as type from divisions
union
select name, map_node_name, 'building' as type from buildings
union
select name, map_node_name, 'gate' as type from gates
order by type;

insert into public.search_groups (id, key, search_group_order, route, search_group_table, search_column, label, label_column, value_column, description_column)
values  ('ee959430-76c7-4cdd-ad19-8e795f92e48c', 'mapObject', null, null, 'map_objects', 'name', null, 'name', 'map_node_name', null);
