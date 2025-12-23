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
    solution TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    solved_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bugs;
-- +goose StatementEnd
