CREATE TABLE IF NOT EXISTS vacancies (
  id SERIAL PRIMARY KEY,
  role VARCHAR,
  company VARCHAR,
  type VARCHAR,
  salary VARCHAR,
  location VARCHAR,
  email VARCHAR NOT NULL
);

