CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    is_deleted BOOLEAN DEFAULT false,
    create_dt TIMESTAMPTZ DEFAULT NOW(),
    update_dt TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX todo_id_not_deleted ON todo (id) WHERE NOT is_deleted;
CREATE INDEX todo_status_not_deleted ON todo (status) WHERE NOT is_deleted;
