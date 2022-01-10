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


INSERT INTO users (email, password) VALUES ('guest@gmail.com', 'password');
INSERT INTO memos (content, done, delete, user_id) VALUES ('test', false, false, 1);
INSERT INTO profits (day, week, month, total, user_id) VALUES (1, 2, 3, 4, 1);
INSERT INTO histories (name, side, amount, ta_date, user_id) VALUES ('BTC', 'SELL', 1, '2021-12-31', 1);
INSERT INTO tickers (day, week, month) VALUES (1, 2, 3);
INSERT INTO signal_events (time, product_code, side, price, size) VALUES ('2006-01-02T15:04:05Z07:00', 'BTC_JPY', 'BUY', 555000.03, 1.0);
INSERT INTO BTC_JPY_1s (time, open, close, high, low, volume) VALUES ('2006-01-02T15:04:05Z07:00', 449.34820443, 449.34820443, 555000.03, 449.34820443, 449.34820443);
INSERT INTO BTC_JPY_1m0s (time, open, close, high, low, volume) VALUES ('2006-01-02T15:04:05Z07:00', 449.34820443, 449.34820443, 555000.03, 449.34820443, 449.34820443);
INSERT INTO BTC_JPY_1h0m0s (time, open, close, high, low, volume) VALUES ('2006-01-02T15:04:05Z07:00', 449.34820443, 449.34820443, 555000.03, 449.34820443, 449.34820443);


-- get signalevent
-- SELECT * FROM 
-- (SELECT time, product_code, side, price, size 
-- FROM signal_events 
-- WHERE product_code = 'BTC_JPY' 
-- ORDER BY time DESC LIMIT 10) 
-- AS signal ORDER BY time ASC;

-- to_timestamp('2022-01-09 16:30:00 +0000 +0000', 'DD/MM/YYYY hh24:mi:ss')::timestamp
-- SELECT time, product_code, side, price, size FROM signal_events WHERE time >= to_timestamp('2022-01-09 16:30:00 +0000 +0000', 'DD/MM/YYYY hh24:mi:ss')::timestamp ORDER BY time DESC;
-- delete from BTC_JPY_1s;
-- delete from BTC_JPY_1m0s;
-- delete from BTC_JPY_1h0m0s; 