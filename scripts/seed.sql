-- scripts/seed.sql  (PostgreSQL)
-- Optional: create simple check constraints (if not using AutoMigrate constraints)
ALTER TABLE IF EXISTS tasks
    ADD CONSTRAINT IF NOT EXISTS chk_task_status CHECK (status IN ('pending','completed')),
    ADD CONSTRAINT IF NOT EXISTS chk_task_priority CHECK (priority IN ('low','medium','high'));

-- OPTIONAL: seed demo user (ISI HASH DENGAN YG KAMU GENERATE)
-- INSERT INTO users (username, email, password_hash, created_at)
-- VALUES ('demo', 'demo@example.com', '<PASTE_BCRYPT_HASH_DI_SINI>', NOW())
-- ON CONFLICT DO NOTHING;

-- Seed contoh task buat user id tertentu (misal 1) — jalankan setelah register atau setelah insert user di atas
INSERT INTO tasks (user_id, title, description, priority, status, created_at)
VALUES
    (1, 'Belajar Go', 'Membuat API Todo', 'high', 'pending', NOW()),
    (1, 'Rapihin README', 'Tambahkan instruksi Docker & Test', 'medium', 'pending', NOW())
    ON CONFLICT DO NOTHING;
