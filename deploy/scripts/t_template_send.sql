-- public.t_template definition
CREATE SEQUENCE public.t_template_send_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
	
CREATE TABLE public.t_template_send (
	id regclass NOT NULL DEFAULT nextval('t_template_send_seq'::regclass),
	queue_id varchar NULL,
	mobile text NULL,
	nationcode varchar NULL,
	app_id int4 NULL, -- 应用ID
	template_type varchar NULL, -- 0：普通短信  1 ：营销短信 （群发）
	template_id varchar NULL, -- 模版编号
	template_param varchar NULL, -- 模版参数
	send_at varchar NULL, -- 指定发送时间
	send_status int4 NULL DEFAULT '-1'::integer, -- -1：待发送 / 0：已发送  /1  : 发送失败
	sender_ip varchar NULL, -- 发送者IP
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT t_template_send_pkey PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_template_send.id IS '主键ID';
COMMENT ON COLUMN public.t_template_send.app_id IS '应用ID';
COMMENT ON COLUMN public.t_template_send.template_type IS '0：普通短信  1 ：营销短信 （群发）';
COMMENT ON COLUMN public.t_template_send.template_id IS '模版编号';
COMMENT ON COLUMN public.t_template_send.template_param IS '模版参数';
COMMENT ON COLUMN public.t_template_send.send_at IS '指定发送时间';
COMMENT ON COLUMN public.t_template_send.send_status IS '-1：待发送 / 0：已发送  /1  : 发送失败';
COMMENT ON COLUMN public.t_template_send.sender_ip IS '发送者IP';