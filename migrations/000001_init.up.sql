CREATE SCHEMA todolist;

CREATE TABLE todolist.users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    encrypted_password VARCHAR(255) NOT NULL,

    CONSTRAINT check_username CHECK(char_length(username) BETWEEN 3 AND 20) 
);

CREATE TABLE todolist.sessions(
    session_token VARCHAR(255) NOT NULL PRIMARY KEY,
    csrf_token VARCHAR(255) NOT NULL,
    user_id INT NOT NULL REFERENCES todolist.users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT check_expires_at CHECK(expires_at > created_at)
);

CREATE TABLE todolist.tasks(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES todolist.users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ,

    CONSTRAINT check_completed_at CHECK(
        (completed_at >= created_at AND completed IS TRUE)
        OR
        (completed_at IS NULL AND completed IS FALSE)
    )
);