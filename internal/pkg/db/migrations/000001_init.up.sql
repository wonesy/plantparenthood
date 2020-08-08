CREATE TABLE IF NOT EXISTS plant(
    id VARCHAR(127) PRIMARY KEY NOT NULL,
    botanical_name VARCHAR(255) UNIQUE NOT NULL,
    common_name VARCHAR(255) NOT NULL,
    sun_preference VARCHAR(63),
    water_preference VARCHAR(63),
    soil_preference VARCHAR(63)
);

CREATE TABLE IF NOT EXISTS member(
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_on TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

CREATE TABLE IF NOT EXISTS care_regimen(
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    water_ml INT NOT NULL,
    water_hr INT NOT NULL
);

CREATE TABLE IF NOT EXISTS plant_baby(
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    owner VARCHAR(128) NOT NULL,
    plant VARCHAR(128) NOT NULL,
    care_regimen VARCHAR(128) NOT NULL,
    nickname VARCHAR(128),
    location VARCHAR(128),
    added_on TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner) REFERENCES member (id),
    FOREIGN KEY (plant) REFERENCES plant (id),
    FOREIGN KEY (care_regimen) REFERENCES care_regimen (id)
);

CREATE TABLE IF NOT EXISTS watering(
    id VARCHAR(128) PRIMARY KEY NOT NULL,
    watered_on TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    amount_ml INT,
    plant_baby_id VARCHAR(128) NOT NULL,
    FOREIGN KEY (plant_baby_id) REFERENCES plant_baby (id)
);