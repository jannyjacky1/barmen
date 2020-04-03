-- +goose Up
INSERT INTO tbl_complication_levels(name, time) VALUES('Легко', '1 мин');
INSERT INTO tbl_complication_levels(name, time) VALUES('Средне', '2-3 мин');
INSERT INTO tbl_complication_levels(name, time) VALUES('Сложно', '5-10 мин');

INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Безалкогольный', 0, 0);
INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Легкий', 1, 12);
INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Средний', 12, 25);
INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Крепкий', 25, 100);

INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('до 60 мл', 0, 60);
INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('60-120 мл', 60, 120);
INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('120-250 мл', 120, 250);
INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('более 250 мл', 250, 1000);

INSERT INTO tbl_settings(name, alias, value) VALUES('Коктель дня', 'day_cocktail', 10);

-- +goose Down
DELETE FROM tbl_complication_levels;
DELETE FROM tbl_fortress_levels;
DELETE FROM tbl_volumes;
DELETE FROM tbl_settings WHERE alias = 'day_cocktail';