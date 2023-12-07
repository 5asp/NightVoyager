-- public.t_sms_channel definition

-- Drop table

CREATE SEQUENCE public.t_app_channel_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP TABLE public.t_sms_channel;

CREATE TABLE public.t_app_channel (
    id regclass NOT NULL DEFAULT nextval('t_app_channel_seq'::regclass),
	channel_name varchar NULL,
	channel_type varchar NULL,
	channel_appkey varchar NULL,
	channel_appsecret varchar NULL,
	channel_domain varchar NULL,
	ext_properties varchar NULL,
	status int2 NULL,
	send_order int4 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	quota int4 NOT NULL DEFAULT 0,
	CONSTRAINT t_app_channel_pkey PRIMARY KEY (id)
);
