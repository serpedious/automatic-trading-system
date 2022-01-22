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

CREATE TABLE signal_events (
  id serial PRIMARY KEY,
  time TIMESTAMP NOT NULL, 
  product_code TEXT,
  side TEXT,
  price FLOAT,
  size FLOAT,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE BTC_JPY_1s (
  id serial PRIMARY KEY,
  time TIMESTAMP NOT NULL,
  open FLOAT,
  close FLOAT,
  high FLOAT,
  low FLOAT,
  volume FLOAT,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE BTC_JPY_1m0s (
  id serial PRIMARY KEY,
  time TIMESTAMP NOT NULL,
  open FLOAT,
  close FLOAT,
  high FLOAT,
  low FLOAT,
  volume FLOAT,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);

CREATE TABLE BTC_JPY_1h0m0s (
  id serial PRIMARY KEY,
  time TIMESTAMP NOT NULL,
  open FLOAT,
  close FLOAT,
  high FLOAT,
  low FLOAT,
  volume FLOAT,
  created_at DATE NOT NULL DEFAULT now(),
  updated_at DATE NOT NULL DEFAULT now()
);


-- INSERT INTO users (email, password) VALUES ('guest@gmail.com', 'password');
-- INSERT INTO memos (content, done, delete, user_id) VALUES ('test', false, false, 1);
-- INSERT INTO signal_events (time, product_code, side, price, size) VALUES ('2006-01-02T15:04:05Z07:00', 'BTC_JPY', 'BUY', 555000.03, 1.0);
-- INSERT INTO BTC_JPY_1s (time, open, close, high, low, volume) VALUES ('2006-01-02T15:04:05Z07:00', 449.34820443, 449.34820443, 555000.03, 449.34820443, 449.34820443);
-- INSERT INTO BTC_JPY_1m0s (time, open, close, high, low, volume) VALUES ('2006-01-02T15:04:05Z07:00', 449.34820443, 449.34820443, 555000.03, 449.34820443, 449.34820443);
-- INSERT INTO BTC_JPY_1h0m0s (time, open, close, high, low, volume) VALUES ('2006-01-02T15:04:05Z07:00', 449.34820443, 449.34820443, 555000.03, 449.34820443, 449.34820443);
