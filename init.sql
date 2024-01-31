CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE messages (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    key varchar NOT NULL,
    value jsonb NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    scheduled_at timestamp NOT NULL,
    is_published bool NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    INDEX idx_scheduled_at ("scheduled_at")
);

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_messages_updated_at
BEFORE UPDATE ON messages
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
