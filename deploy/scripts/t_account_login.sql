-- public.t_account_login definition

-- Drop table

-- DROP SEQUENCE public.t_account_login_seq;

CREATE SEQUENCE public.t_account_login_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP TABLE public.t_account_login;

CREATE TABLE public.t_account_login (
	id regclass NOT NULL DEFAULT nextval('t_account_login_seq'::regclass),
	account_id int4 NOT NULL,
	ip_addr text NOT NULL,
	created_at timestamptz NOT NULL,
	device text NULL,
    CONSTRAINT t_account_login_pkey PRIMARY KEY (id)
);

-- public.t_account_login_seq definition
