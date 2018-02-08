CREATE TABLE book (
  book_id char(36) COLLATE pg_catalog."default" NOT NULL,
  name varchar(100) COLLATE pg_catalog."default" NOT NULL,
  CONSTRAINT book_id_pkey PRIMARY KEY (book_id)
);