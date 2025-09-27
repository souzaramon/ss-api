-- +goose Up
-- +goose StatementBegin
CREATE TABLE titles (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  id_author UUID REFERENCES authors(id),
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP      
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE titles;
-- +goose StatementEnd
