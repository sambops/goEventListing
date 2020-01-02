create table events (
    id  integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL,
    details text,
    price numeric NOT NULL DEFAULT 0,
    image varchar(255),

    city varchar(255),
    country varchar(255),
    town varchar(255),
    coordinates varchar(255),
    rating numeric NOT NULL DEFAULT 0,
    users_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    tag_id integer REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE,
    ispassed boolean NOT NULL DEFAULT false
); 



create table rating (
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    rating integer NOT NULL DEFAULT 0,
  event_id integer REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE,
  user_id integer REFERENCES users(id) ON UPDATE CASCADE,
  
);


create table tags (
    
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(255) NOT NULL,    
    description text,
    icon varchar(255)

);


create table event_tag (

    event_id integer REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE,
    tag_id integer REFERENCES categories(id) ON UPDATE CASCADE,
    CONSTRAINT event_tags_pky PRIMARY KEY (event_id, tag_id) 

);


create table user_tag (
  user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  tag_id integer REFERENCES categories(id) ON UPDATE CASCADE,
  CONSTRAINT event_tags_pky PRIMARY KEY (event_id, tag_id)
);


create table users (
  id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  user_name varchar(64) NOT NULL UNIQUE,
  full_name varchar(255),
  email varchar(255) NOT NULL UNIQUE,
  phone varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL
);


create table notifications (
  user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  event_id integer REFERENCES events(id) ON UPDATE CASCADE,
  CONSTRAINT user_events_pkey PRIMARY KEY (user_id, event_id)
);






