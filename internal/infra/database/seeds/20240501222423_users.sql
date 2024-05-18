--Seed files MUST be idempotent by design
INSERT INTO users (name, email, created_at)
VALUES
  ('John Doe', 'jdoe@gmail.com', '2024-04-20 21:15:00'),
  ('Mari Doe', 'mdoe@gmail.com', '2024-04-20 21:16:00'),
  ('Selena Hernandez', 'shernandez@gmail.com', '2024-04-21 13:16:00'),
  ('Hare Krishna', 'hkrishna@gmail.com', '2024-04-21 13:17:00'),
  ('Manuel Lopez', 'mlopez@gmail.com', '2024-04-21 13:18:00'),
  ('Deidre Esteban', 'desteban@gmail.com', '2024-04-21 13:19:00'),
  ('Xe Leong', 'xleong@gmail.com',  '2024-04-21 13:20:00'),
  ('Alice Mu', 'amu@gmail.com', '2024-04-21 13:21:00'),
  ('Bob Smith', 'bsmith@gmail.com', '2024-04-21 13:22:00')
ON CONFLICT (email) DO NOTHING
RETURNING *;


