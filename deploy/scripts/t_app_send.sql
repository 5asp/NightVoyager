CREATE SEQUENCE public.t_app_send_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;


CREATE TABLE public.t_app_send (
	id regclass NOT NULL DEFAULT nextval('t_app_send_seq'::regclass),
	account_id int4 NOT NULL,
	app_id int4 NOT NULL,
	send_status int4 NULL,
	created_at timestamptz NULL,
	extra jsonb NULL,
	reason text NULL,
	CONSTRAINT t_app_send_pk PRIMARY KEY (id)
);
-- Column comments

COMMENT ON COLUMN public.t_app_send.send_status IS '发送状态';