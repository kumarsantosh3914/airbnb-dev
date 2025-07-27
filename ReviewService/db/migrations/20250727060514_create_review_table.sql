-- +goose Up
-- +goose StatementBegin
CREATE TABLE review (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    hotel_id INT NOT NULL,
    booking_id INT NOT NULL,
    comment  TEXT,
    rating INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_synced BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS review;
-- +goose StatementEnd
