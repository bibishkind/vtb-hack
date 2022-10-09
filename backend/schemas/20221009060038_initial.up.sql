CREATE TABLE users (
	id serial PRIMARY KEY,
	username varchar(255) UNIQUE NOT NULL,
	password varchar(255) NOT NULL,
	first_name varchar(255) NOT NULL DEFAULT '',
	middle_name varchar(255) NOT NULL DEFAULT '',
	last_name varchar(255) NOT NULL DEFAULT '',
	email varchar(255) NOT NULL DEFAULT '',
	role int NOT NULL DEFAULT 0,
	public_key varchar(255) NOT NULL,
	private_key varchar(255) NOT NULL
);

CREATE TABLE teams (
	id serial PRIMARY KEY,
	name varchar(255) UNIQUE NOT NULL
);

CREATE TABLE users_teams (
	user_id int,
	team_id int,

	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE tasks (
	id serial PRIMARY KEY,
	title varchar(255) NOT NULL,
	description varchar(255) NOT NULL DEFAULT '',
	body text NOT NULL DEFAULT '',
	revenue real NOT NULL DEFAULT 0,
	type int NOT NULL DEFAULT 0,
	priority int NOT NULL DEFAULT 0,
	status int NOT NULL DEFAULT 0,
	thumbnail varchar(255) NOT NULL DEFAULT ''
);

CREATE TABLE teams_tasks (
	team_id int,
	task_id int,

	FOREIGN KEY (team_id) REFERENCES teams(id),
	FOREIGN KEY (task_id) REFERENCES tasks(id)
);

CREATE TABLE users_tasks (
	user_id int,
	task_id int,

	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (task_id) REFERENCES tasks(id)
);

CREATE TABLE cards (
	id serial PRIMARY KEY,
	title varchar(255) NOT NULL,
	description varchar(255) NOT NULL DEFAULT '',
	body text NOT NULL DEFAULT '',
	price real NOT NULL DEFAULT 0,
	thumbnail varchar(255) NOT NULL DEFAULT ''
);

CREATE TABLE scores (
    user_id int UNIQUE,
    score int NOT NULL DEFAULT 0,

    FOREIGN KEY (user_id) REFERENCES users(id)
);