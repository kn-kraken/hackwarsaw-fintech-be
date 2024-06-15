CREATE TABLE IF NOT EXISTS real_estates (
   real_estate_id INT GENERATED ALWAYS AS IDENTITY,
   address        VARCHAR(200),
   occurance_type VARCHAR(200),
   area           DECIMAL,
   initial_price  DECIMAL,
   district       VARCHAR(200),
   PRIMARY KEY(real_estate_id)
);

CREATE TABLE IF NOT EXISTS real_estate_destinations (
   real_estate_id INT,
   num INT,
   destination VARCHAR(200),
   PRIMARY KEY(real_estate_id, num),
   CONSTRAINT fk_real_estate
      FOREIGN KEY(real_estate_id) 
      REFERENCES real_estates(real_estate_id)
      ON DELETE CASCADE
);
