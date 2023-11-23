DROP TABLE IF EXISTS event_belongs_to;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    password VARCHAR(60) NOT NULL,
    email VARCHAR(200) NOT NULL
);

CREATE TABLE events(
    event_id SERIAL PRIMARY KEY,
    event_name VARCHAR(200),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE TABLE event_belongs_to(
    user_id INT,
    event_id INT,
    FOREIGN KEY (event_id) REFERENCES events ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE
);
