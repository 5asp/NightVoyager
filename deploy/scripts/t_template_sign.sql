CREATE SEQUENCE public.t_template_sign_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;

-- public.t_template_sign definition

-- Drop table

-- DROP TABLE public.t_template_sign;

CREATE TABLE public.t_template_sign (
	id regclass NOT NULL DEFAULT nextval('t_template_sign_seq'::regclass),
	sign_name varchar NOT NULL,
	is_default int4 NULL DEFAULT 0,
	status int4 NULL DEFAULT 1,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	account_id int4 NOT NULL DEFAULT 0,
	CONSTRAINT t_template_sign_pkey PRIMARY KEY (id)
);


-- public.t_template_sign foreign keys