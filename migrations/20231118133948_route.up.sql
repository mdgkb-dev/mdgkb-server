create table map_nodes (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar(255),
  is_entry boolean
);

create table map_edges (
  previous_node_id uuid references map_nodes(id),
  next_node_id uuid references map_nodes(id),
  primary key (previous_node_id, next_node_id)
);

create table map_routes (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  start_node_id uuid references map_nodes(id),
  end_node_id uuid references map_nodes(id)
);

create table map_route_nodes (
  map_route_id uuid references map_routes(id),
  map_node_id uuid references map_nodes(id),
  item_order int default 0
);

-- insert into nodes(name,isentry) values ('корпус 17', true);
-- insert into nodes(name,isentry) values ('Корпус 18', true);
-- insert into nodes(name,isentry) values ('Корпус 21', true);
-- insert into nodes(name,isentry) values ('Перекресток1', false);
-- insert into nodes(name,isentry) values ('Корпус 13', true);
-- insert into nodes(name,isentry) values ('Корпус 10', true);
-- insert into nodes(name,isentry) values ('Корпус 9', true);
-- insert into nodes(name,isentry) values ('Корпус 7', true);
-- insert into nodes(name,isentry) values ('Корпус 15', true);
-- insert into nodes(name,isentry) values ('Корпус 22а', true);
-- insert into nodes(name,isentry) values ('Корпус 1', true);
-- insert into nodes(name,isentry) values ('Корпус 1а', true);

-- insert into edges(previous_node,next_node) values (1,2);
-- insert into edges(previous_node,next_node) values (1,4);
-- insert into edges(previous_node,next_node) values (1,9);
-- insert into edges(previous_node,next_node) values (1,10);

-- insert into edges(previous_node,next_node) values (2,1);
-- insert into edges(previous_node,next_node) values (2,3);
-- insert into edges(previous_node,next_node) values (2,9);

-- insert into edges(previous_node,next_node) values (3,2);
-- insert into edges(previous_node,next_node) values (3,10);
-- insert into edges(previous_node,next_node) values (3,12);

-- insert into edges(previous_node,next_node) values (4,1);
-- insert into edges(previous_node,next_node) values (4,5);
-- insert into edges(previous_node,next_node) values (4,6);
-- insert into edges(previous_node,next_node) values (4,9);

-- insert into edges(previous_node,next_node) values (5,4);
-- insert into edges(previous_node,next_node) values (5,6);
-- insert into edges(previous_node,next_node) values (5,7);

-- insert into edges(previous_node,next_node) values (6,4);
-- insert into edges(previous_node,next_node) values (6,5);
-- insert into edges(previous_node,next_node) values (6,8);
-- insert into edges(previous_node,next_node) values (6,9);

-- insert into edges(previous_node,next_node) values (7,5);
-- insert into edges(previous_node,next_node) values (7,8);

-- insert into edges(previous_node,next_node) values (8,6);
-- insert into edges(previous_node,next_node) values (8,7);

-- insert into edges(previous_node,next_node) values (9,1);
-- insert into edges(previous_node,next_node) values (9,2);
-- insert into edges(previous_node,next_node) values (9,4);
-- insert into edges(previous_node,next_node) values (9,6);

-- insert into edges(previous_node,next_node) values (10,1);
-- insert into edges(previous_node,next_node) values (10,3);
-- insert into edges(previous_node,next_node) values (10,11);
-- insert into edges(previous_node,next_node) values (10,12);

-- insert into edges(previous_node,next_node) values (11,10);
-- insert into edges(previous_node,next_node) values (11,12);

-- insert into edges(previous_node,next_node) values (12,3);
-- insert into edges(previous_node,next_node) values (12,10);
-- insert into edges(previous_node,next_node) values (12,11);

-- select n.id,e.next_node from nodes n inner join edges e on n.id = e.previous_node order by n.id