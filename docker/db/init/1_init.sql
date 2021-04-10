CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE memos (
  id SERIAL PRIMARY KEY,
  user_id int REFERENCES users(id),
  content TEXT,
  done BOOLEAN,
  delete BOOLEAN,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE profits (
  id SERIAL PRIMARY KEY,
  user_id int REFERENCES users(id),
  day INT,
  week INT,
  month INT,
  total INT,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE histories (
  id SERIAL PRIMARY KEY,
  user_id int REFERENCES users(id),
  name TEXT,
  side TEXT,
  amount INT,
  ta_date DATE,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE tickers (
  id serial PRIMARY KEY,
  day INT,
  week INT,
  month INT,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

INSERT INTO users (email, password) VALUES ('hoge@sample.co.jp', 'golang');
INSERT INTO memos (content, done, user_id) VALUES ('test', true, 1);
INSERT INTO profits (day, week, month, total, user_id) VALUES (1, 2, 3, 4, 1);
INSERT INTO histories (name, side, amount, ta_date, user_id) VALUES ('BTC', 'SELL', 1, '2021-12-31', 1);
INSERT INTO tickers (day, week, month) VALUES (1, 2, 3);
