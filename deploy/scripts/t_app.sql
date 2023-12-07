-- public.t_app definition
-- DROP SEQUENCE public.t_app_seq;

CREATE SEQUENCE public.t_app_seq
	INCREMENT BY 1
	MINVALUE 10000
	MAXVALUE 9223372036854775807
	START 10000
	CACHE 1
	NO CYCLE;
-- Drop table

-- public.t_app definition

-- Drop table

-- DROP TABLE public.t_app;

CREATE TABLE public.t_app (
	id regclass NOT NULL DEFAULT nextval('t_app_seq'::regclass),
	secret varchar NULL, -- 应用密钥
	status int4 NULL, -- 应用状态
	remark varchar NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	account_id int4 NOT NULL DEFAULT 0,
	"name" varchar NULL, -- 应用名称
	CONSTRAINT t_app_pkey PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_app.secret IS '应用密钥';
COMMENT ON COLUMN public.t_app.status IS '应用状态';
COMMENT ON COLUMN public.t_app."name" IS '应用名称';

