CREATE TABLE IF NOT EXISTS app_rating(
   id SERIAL PRIMARY KEY,
   app_id INTEGER NOT NULL,
   country_id INTEGER NOT NULL,
   category_id INTEGER NOT NULL,
   sub_category_id INTEGER NOT NULL,
   date_rate date NOT NULL,
   pos  INTEGER NOT NULL,
   UNIQUE (app_id, country_id, category_id, sub_category_id, date_rate)
);