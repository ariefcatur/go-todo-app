INSERT INTO users (id, username, email, password_hash, created_at)
VALUES (1, 'demo', 'demo@example.com', '$2a$10$E1...hash_bcrypt...', NOW())
    ON CONFLICT DO NOTHING;
