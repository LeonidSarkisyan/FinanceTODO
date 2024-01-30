-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE sub_categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT REFERENCES categories(id)
);

INSERT INTO categories (title) VALUES ('Еда'), ('Транспорт'), ('Покупки'), ('Здоровье'), ('Личные');

INSERT INTO sub_categories (title, category_id) VALUES
    ('Продукты', 1),
    ('Ресторан, фаст-фуд', 1),
    ('Кафе, бар', 1),
    ('Дальние поездки', 2),
    ('Деловые поездки', 2),
    ('Общественный транспорт', 2),
    ('Такси', 2),
    ('Аптека', 3),
    ('Дети', 3),
    ('Дом и сад', 3),
    ('Домашние животные', 3),
    ('Красота и здоровье', 3),
    ('Одежда и обувь', 3),
    ('Отдых', 3),
    ('Подарки', 3),
    ('Электроника', 3),
    ('Ювелирные изделия', 3),
    ('Клиника', 4),
    ('Операции', 4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sub_categories DROP CONSTRAINT sub_categories_category_id_fkey;

DROP TABLE IF EXISTS sub_categories;

DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
