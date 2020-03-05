--创建用户表
CREATE TABLE "public"."K_User" (
"user_id" int4 DEFAULT nextval('test_c_id_seq'::regclass) NOT NULL,
"user_name" varchar(255) COLLATE "default" NOT NULL,
"user_cc" float8 DEFAULT 0,
"register_time" timestamp(6) NOT NULL,
"head_shot_url" varchar(255) COLLATE "default" NOT NULL,
"is_delete" bool DEFAULT false,
CONSTRAINT "K_User_pkey" PRIMARY KEY ("user_id")
)
WITH (OIDS=FALSE)


--创建项目表
CREATE TABLE "public"."K_Project" (
"project_id" int4 DEFAULT nextval('"K_Project1_project_id_seq"'::regclass) NOT NULL,
"project_name" varchar(255) COLLATE "default" NOT NULL,
"project_url" varchar(255) COLLATE "default" NOT NULL,
"project_cover_url" varchar(255) COLLATE "default" NOT NULL,
"project_institution" int4,
"project_description" text COLLATE "default",
CONSTRAINT "k_project_pkey" PRIMARY KEY ("project_id")
)
WITH (OIDS=FALSE)
;

ALTER TABLE "public"."K_Project" OWNER TO "sspkukcoin";

COMMENT ON TABLE "public"."K_Project" IS '项目表';

COMMENT ON COLUMN "public"."K_Project"."project_id" IS '项目id（主键）';

COMMENT ON COLUMN "public"."K_Project"."project_name" IS '项目名称';

COMMENT ON COLUMN "public"."K_Project"."project_url" IS '项目url';

COMMENT ON COLUMN "public"."K_Project"."project_cover_url" IS '项目封面urll';

COMMENT ON COLUMN "public"."K_Project"."project_institution" IS '项目制度（1是独裁制，2是委员会占比制，3是委员会单票制，4是全体占比制，5是全体单票制）';

--创建用户 in项目表
CREATE TABLE "public"."K_User_in_Project" (
"item_id" int4 DEFAULT nextval('"K_User_in_Project1_item_id_seq"'::regclass) NOT NULL,
"project_id" int4 NOT NULL,
"user_id" int4 NOT NULL,
"user_cs" float8 DEFAULT 0 NOT NULL,
"user_cc" float8 DEFAULT 0 NOT NULL,
CONSTRAINT "k_user_in_project_pkey" PRIMARY KEY ("user_id", "project_id"),
CONSTRAINT "K_User_in_Project_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."K_User" ("user_id") ON DELETE NO ACTION ON UPDATE NO ACTION,
CONSTRAINT "K_User_in_Project_project_id_fkey" FOREIGN KEY ("project_id") REFERENCES "public"."K_Project" ("project_id") ON DELETE NO ACTION ON UPDATE NO ACTION
)
WITH (OIDS=FALSE)
;

--创建触发器
CREATE TRIGGER "cc_trigger" AFTER UPDATE ON "public"."K_User_in_Project"
FOR EACH ROW
EXECUTE PROCEDURE "sumcc"();


--创建权限表
CREATE TABLE "public"."K_Permission" (
"id" int4 DEFAULT nextval('"K_Permission_id_seq"'::regclass) NOT NULL,
"Permission_id" int4 NOT NULL,
"Permission" varchar(100) COLLATE "default" NOT NULL,
CONSTRAINT "K_Permission_pkey" PRIMARY KEY ("id")
)
WITH (OIDS=FALSE)
;

ALTER TABLE "public"."K_Permission" OWNER TO "sspkukcoin";

COMMENT ON TABLE "public"."K_Permission" IS '权限表';

COMMENT ON COLUMN "public"."K_Permission"."id" IS '主键ID';

COMMENT ON COLUMN "public"."K_Permission"."Permission_id" IS '权限ID';

COMMENT ON COLUMN "public"."K_Permission"."Permission" IS '权限名';

--创建角色表
CREATE TABLE "public"."k_role" (
"id" int4 DEFAULT nextval('k_role_id_seq'::regclass) NOT NULL,
"role_id" int4 NOT NULL,
"role_name" varchar(100) COLLATE "default" NOT NULL,
"role_permissions" int4 NOT NULL,
"is_able" bool DEFAULT false NOT NULL,
CONSTRAINT "k_role_pkey" PRIMARY KEY ("id"),
CONSTRAINT "k_role_role_permissions_fkey" FOREIGN KEY ("role_permissions") REFERENCES "public"."K_Permission" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
)
WITH (OIDS=FALSE)
;

ALTER TABLE "public"."k_role" OWNER TO "sspkukcoin";

COMMENT ON TABLE "public"."k_role" IS '角色表';

COMMENT ON COLUMN "public"."k_role"."id" IS '主键ID';

COMMENT ON COLUMN "public"."k_role"."role_id" IS '用户ID';

COMMENT ON COLUMN "public"."k_role"."role_name" IS '用户名';

COMMENT ON COLUMN "public"."k_role"."role_permissions" IS '用户权限';

COMMENT ON COLUMN "public"."k_role"."is_able" IS '用户是否使用';

--创建项目中，用户的角色表
CREATE TABLE "public"."k_user_role_in_project" (
"id" int4 DEFAULT nextval('k_user_role_in_project_id_seq'::regclass) NOT NULL,
"project_id" int4 NOT NULL,
"user_id" int4 NOT NULL,
"role_id" int4 NOT NULL,
CONSTRAINT "k_user_role_in_project_pkey" PRIMARY KEY ("id"),
CONSTRAINT "k_user_role_in_project_project_id_fkey" FOREIGN KEY ("project_id") REFERENCES "public"."K_Project" ("project_id") ON DELETE NO ACTION ON UPDATE NO ACTION,
CONSTRAINT "k_user_role_in_project_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."K_User" ("user_id") ON DELETE NO ACTION ON UPDATE NO ACTION
)
WITH (OIDS=FALSE)
;

