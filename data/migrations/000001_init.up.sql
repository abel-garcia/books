SET client_encoding = UTF8;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


/* 
* Create function updated timestamp 
*/
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;