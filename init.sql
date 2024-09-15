CREATE DATABASE poc;

CREATE TABLE IF NOT EXISTS games (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  metadata JSON
);

INSERT INTO games (title, metadata) VALUES ('Kindom Hearts', '{"tags": ["RPG"]}');

INSERT INTO games (title, description, metadata) VALUES ('God of War', 'Franchise reboot', '{"platforms": ["PS4"]}');

INSERT INTO games (title) VALUES ('The Legend of Zelda');