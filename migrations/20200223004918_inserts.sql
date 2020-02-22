-- +goose Up
INSERT INTO tbl_complication_levels(name, time) VALUES('Легко', '5 мин');
INSERT INTO tbl_complication_levels(name, time) VALUES('Средне', '10 мин');
INSERT INTO tbl_complication_levels(name, time) VALUES('Сложно', '15-20 мин');

INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Безалкогольный', 0, 0);
INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Легкий', 1, 15);
INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Средний', 15, 30);
INSERT INTO tbl_fortress_levels(name, fortress_from, fortress_to) VALUES('Крепкий', 30, 100);

INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('до 60 мл', 0, 60);
INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('60-120 мл', 60, 120);
INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('120-250 мл', 120, 250);
INSERT INTO tbl_volumes(name, volume_from, volume_to) VALUES('более 250 мл', 250, 1000);

-- +goose Down
DELETE FROM tbl_complication_levels;
DELETE FROM tbl_fortress_levels;
DELETE FROM tbl_volumes;