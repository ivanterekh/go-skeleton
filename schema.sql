CREATE TABLE "USER"
(
  id       SERIAL PRIMARY KEY ,
  name     VARCHAR(100),
  role     VARCHAR(50) DEFAULT 'user',
  email    VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL
);
