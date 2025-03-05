CREATE TABLE users
(
    id SERIAL NOT NULL UNIQUE,
    userlogin varchar(255) NOT NULL UNIQUE,
    userpass varchar(255) NOT NULL
    userRole varchar(6) NOT NULL
)

CREATE TABLE tasks
(
    id SERIAL NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    description TEXT,
    done BOOLEAN NOT NULL
    created timestamp NOT NULL
)

CREATE TABLE userstasks
(
    id SERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES user(id) ON DELETE CASCADE NOT NULL,
    task_id INT REFERENCES task(id) ON DELETE CASCADE NOT NULL
)