-- public.t_template definition
CREATE TABLE public.t_template (
	id varchar NOT NULL,
	tpl_content varchar NULL,
	tpl_type varchar NULL DEFAULT 0, -- 0：验证码。\n1：短信通知。\n2：推广短信。\n3：国际/港澳台消息
	sign_id int4 NOT NULL DEFAULT 0,
	account_id int4 NOT NULL DEFAULT 0,
	status int4 NULL DEFAULT 1, -- 状态 0：有效 1：无效
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	remark varchar NULL,
	CONSTRAINT tpl_pkey PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.t_template.tpl_type IS '0：验证码。\n1：短信通知。\n2：推广短信。\n3：国际/港澳台消息';
COMMENT ON COLUMN public.t_template.status IS '状态 0：有效 1：无效';
