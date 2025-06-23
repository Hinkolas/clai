CREATE TABLE
    IF NOT EXISTS chats (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        model TEXT NOT NULL,
        status TEXT NOT NULL,
        is_pinned INTEGER NOT NULL,
        created_at INTEGER NOT NULL,
        updated_at INTEGER NOT NULL,
        last_message_at INTEGER NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS messages (
        id TEXT PRIMARY KEY,
        chat_id TEXT NOT NULL,
        stream_id TEXT NOT NULL,
        role TEXT NOT NULL,
        model TEXT NOT NULL,
        content TEXT NOT NULL,
        reasoning TEXT NOT NULL,
        status TEXT NOT NULL,
        created_at INTEGER NOT NULL,
        updated_at INTEGER NOT NULL,
        FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS attachments (
        id TEXT PRIMARY KEY,
        message_id TEXT NOT NULL,
        name TEXT NOT NULL,
        type TEXT NOT NULL,
        src TEXT NOT NULL,
        created_at INTEGER NOT NULL,
        updatet_at INTEGER NOT NULL,
        FOREIGN KEY (message_id) REFERENCES messages (id) ON DELETE CASCADE
    );

-- Create indexes for more efficient querying
CREATE INDEX IF NOT EXISTS idx_messages_chat_id ON messages (chat_id);

CREATE INDEX IF NOT EXISTS idx_attachments_message_id ON attachments (message_id);

CREATE INDEX IF NOT EXISTS idx_messages_chat_id_created_at ON messages (chat_id, created_at);

CREATE INDEX IF NOT EXISTS idx_chats_last_message_at ON chats (last_message_at);

-- Create triggers for automatic updates

CREATE TRIGGER update_chat_on_message_create AFTER INSERT ON messages FOR EACH ROW BEGIN
UPDATE chats
SET
    last_message_at = NEW.created_at,
    status = NEW.status
WHERE
    id = NEW.chat_id;

END;

CREATE TRIGGER update_chat_on_message_change AFTER
UPDATE OF status ON messages FOR EACH ROW WHEN OLD.status != NEW.status BEGIN
UPDATE chats
SET
    status = NEW.status
WHERE
    id = NEW.chat_id;

END;
