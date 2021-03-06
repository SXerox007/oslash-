-- uuid extenstion 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- https://stackoverflow.com/questions/22446478/extension-exists-but-uuid-generate-v4-fails


-- uuid generate
Select uuid_generate_v1mc();
SELECT uuid_generate_v5();
SELECT uuid_generate_v1();
SELECT uuid_generate_v4();
SELECT uuid_in(md5(random()::text || clock_timestamp()::text)::cstring);
SELECT md5('Store hash for long string, maybe for index?')::uuid AS md5_hash;
