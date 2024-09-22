SELECT EXISTS (
    SELECT FROM information_schema.tables 
    WHERE table_schema = 'urlshortener' 
    AND table_name = 'keys'
);

CREATE TABLE IF NOT EXISTS urlshortener.keys (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP WITH TIME ZONE,
    ip VARCHAR(16),
    url VARCHAR(2048),
    key VARCHAR(32),
    CONSTRAINT unique_key UNIQUE (key)
);

CREATE INDEX idx_keys_key ON urlshortener.keys(key)

INSERT INTO urlshortener.keys (timestamp, ip, url, key) 
VALUES ('2024-02-23 3:32:32', '192.168.0.1', 'http://google.com', 'abc')

SELECT * FROM urlshortener.keys
WHERE key = 'abc'