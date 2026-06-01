CREATE SCHEMA todolist;

CREATE TABLE todolist.users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL,

    CONSTRAINT check_username CHECK(length(username) > 8) 
);

CREATE TABLE todolist.sessions(
    session_token VARCHAR(255) NOT NULL PRIMARY KEY,
    csrf_token VARCHAR(255) NOT NULL,
    user_id INT NOT NULL REFERENCES todolist.users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,

    CONSTRAINT check_expires_at CHECK(expires_at > created_at)
);

CREATE TABLE todolist.tasks(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES todolist.users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP,

    CONSTRAINT check_completed_at CHECK(completed_at >= created_at)
);