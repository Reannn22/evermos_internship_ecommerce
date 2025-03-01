-- Count all users
SELECT COUNT(*) FROM users;

-- Show all users
SELECT id, nama, email, notelp FROM users;

-- Check for soft deleted records
SELECT id, nama, email, notelp, deleted_at 
FROM users 
WHERE deleted_at IS NOT NULL;

-- Check for any constraints that might affect insertion
SHOW CREATE TABLE users;
