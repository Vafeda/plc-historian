-- Таблица для точек, отсеянных deadband'ом. Хранит только последние часы.
-- Колонки и типы — как у history; id берётся из той же последовательности,
-- поэтому номера в history и history_raw не повторяются.
CREATE TABLE IF NOT EXISTS history_raw (
	id            BIGINT           NOT NULL DEFAULT nextval('history_id_seq'),
	itemid        BIGINT,
	itemtimestamp TIMESTAMP        NOT NULL,
	regtimestamp  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
	valint        BIGINT,
	valdouble     DOUBLE PRECISION,
	valbool       BOOLEAN,
	valstring     TEXT,
	quality       INTEGER
);

-- Гипертаблица кусками по часу: автоудаление снимает старые куски целиком.
SELECT create_hypertable('history_raw', by_range('itemtimestamp', INTERVAL '1 hour'), if_not_exists => TRUE);

-- Индекс для запросов «один сигнал за период».
CREATE INDEX IF NOT EXISTS history_raw_itemid_ts_idx ON history_raw (itemid, itemtimestamp DESC);

-- Хранить 2 часа (показываем последний час, час — запас).
SELECT add_retention_policy('history_raw', drop_after => INTERVAL '2 hours', if_not_exists => TRUE);

-- Всё вместе для Grafana: архив целиком + отсеянные точки за последний час.
CREATE OR REPLACE VIEW history_all AS
SELECT * FROM history_raw                  -- последний час: все точки
WHERE itemtimestamp >= (now() AT TIME ZONE 'UTC') - INTERVAL '1 hour'
UNION ALL
SELECT * FROM history                      -- всё, что старше: после deadband
WHERE itemtimestamp <  (now() AT TIME ZONE 'UTC') - INTERVAL '1 hour';