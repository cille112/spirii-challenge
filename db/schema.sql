CREATE TABLE data (
    meterId TEXT,
    consumerId Text,
    timestamp DATETIME,
    meterReading INTEGER,
    UNIQUE(timestamp, meterId)
);