CREATE TABLE IF NOT EXISTS workflows (
    id uuid PRIMARY KEY,
    nodes jsonb,
    edges jsonb
);