CREATE TABLE hospitalization_type_analyzes (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  hospitalization_type_id uuid REFERENCES hospitalizations_types(id) ON UPDATE CASCADE ON DELETE CASCADE,
  name varchar,
  children boolean,
  duration_days integer
);

CREATE TABLE hospitalization_type_documents (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  hospitalization_type_id uuid REFERENCES hospitalizations_types(id) ON UPDATE CASCADE ON DELETE CASCADE,
  name varchar,
  children boolean
);


CREATE TABLE hospitalization_type_stages (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  hospitalization_type_id uuid REFERENCES hospitalizations_types(id) ON UPDATE CASCADE ON DELETE CASCADE,
  name varchar
);
