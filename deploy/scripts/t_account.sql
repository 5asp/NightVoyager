-- public.t_account definition

-- Drop table
CREATE SEQUENCE public.t_account_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP TABLE public.t_account;

CREATE TABLE public.t_account (
	id regclass NOT NULL DEFAULT nextval('t_account_seq'::regclass),
	account text NOT NULL,
	"password" text NOT NULL,
	status int2 NOT NULL DEFAULT 1,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT t_account_pkey PRIMARY KEY (id)
);

