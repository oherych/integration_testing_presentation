CREATE TABLE book (
  book_id varchar(100) COLLATE pg_catalog."default" NOT NULL,
  name varchar(100) COLLATE pg_catalog."default" NOT NULL,
  CONSTRAINT file_pkey PRIMARY KEY (book_id)
)
WITH (
OIDS = FALSE
)
TABLESPACE pg_default;