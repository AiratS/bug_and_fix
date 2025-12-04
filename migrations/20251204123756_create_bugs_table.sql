-- +goose Up
-- +goose StatementBegin
CREATE TABLE bugs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(1024) NOT NULL,
    description TEXT NOT NULL,
    project VARCHAR(255) NOT NULL,
    error_message TEXT,
    device VARCHAR(127) NOT NULL,
    env INTEGER NOT NULL,
    solution TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bugs;
-- +goose StatementEnd
