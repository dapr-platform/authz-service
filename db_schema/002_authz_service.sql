-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists o_user(
                       id VARCHAR(25) NOT NULL,
                       tenant_id VARCHAR(25) NOT NULL,
                       mobile VARCHAR(15) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       identity VARCHAR(255) NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       gender integer NOT NULL default 0,
                       address VARCHAR(1024) NOT NULL default '',
                       password VARCHAR(255) NOT NULL,
                       type INTEGER NOT NULL,
                       org_id VARCHAR(255) NOT NULL,
                       avatar_url VARCHAR(255) NOT NULL,
                       create_at TIMESTAMP NOT NULL,
                       update_at TIMESTAMP NOT NULL,
                       status INTEGER NOT NULL,
                       PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IDX_o_user_NAME ON o_user(identity);

COMMENT ON COLUMN o_user.id IS 'Primary Key';
COMMENT ON COLUMN o_user.identity IS '用户标识';
COMMENT ON COLUMN o_user.type IS '用户类型';
COMMENT ON COLUMN o_user.org_id IS '组织ID';
COMMENT ON COLUMN o_user.avatar_url IS '头像';
COMMENT ON COLUMN o_user.create_at IS '创建时间';
COMMENT ON COLUMN o_user.update_at IS '更新时间';
COMMENT ON COLUMN o_user.status IS '状态(1正常，2:禁止登陆，3:删除';

create extension if not exists tablefunc;
create extension if not exists pgcrypto;

create or replace function public.nanoid(size integer default 21)
    returns text
    language plpgsql
    stable
as $function$
declare
    id text := '';
    i int := 0;
    urlalphabet char(64) := '_-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
    bytes bytea := gen_random_bytes(size);
    byte int;
    pos int;
begin
    while i < size loop
            byte := get_byte(bytes, i);
            pos := (byte & 63) + 1;
            id := id || substr(urlalphabet, pos, 1);
            i = i + 1;
        end loop;
    return id;
end
$function$;

insert into o_user(id, tenant_id,mobile, email, identity,name,gender,address, password, type, org_id,  avatar_url, create_at, update_at,  status)
values ((select nanoid()),'','+8613911111111', 'admin@business.com','admin','admin',1,'',(SELECT encode(digest('things2024'::bytea, 'sha1'), 'hex')),1,'','',now(),now(),1 );


CREATE TABLE o_resource(
                           id VARCHAR(255) NOT NULL,
                           name VARCHAR(255) NOT NULL,
                           parent_id varchar(255) not null default '',
                           module VARCHAR(255) NOT NULL,
                           service_name VARCHAR(255) NOT NULL,
                           service_name_cn VARCHAR(255) NOT NULL,
                           type INTEGER NOT NULL,
                           api_url VARCHAR(255) NOT NULL,
                           sort_index integer not null default 0,
                           data_conditions VARCHAR(10240) NOT NULL,
                           support_ops VARCHAR(255) NOT NULL,
                           PRIMARY KEY (id)
);
COMMENT ON TABLE o_resource IS '权限点资源';
COMMENT ON COLUMN o_resource.id IS '唯一标识,例如UI-001-001';
COMMENT ON COLUMN o_resource.name IS '名称';
COMMENT ON COLUMN o_resource.module IS '模块';
COMMENT ON COLUMN o_resource.service_name IS '服务名';
COMMENT ON COLUMN o_resource.service_name_cn IS '中文服务名';
COMMENT ON COLUMN o_resource.type IS '分类（1:菜单,2:api,3:功能点,4:数据)';
COMMENT ON COLUMN o_resource.api_url IS 'api鉴权用到';
COMMENT ON COLUMN o_resource.data_conditions IS '数据条件（数据鉴权用到，存放对象的json字符串)';
COMMENT ON COLUMN o_resource.support_ops IS '支持的操作(字典值,0:r,1:rw)';


CREATE TABLE o_role(
                       id VARCHAR(255) NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       sort_index integer not null default 0,
                       status integer not null default 1,
                       create_at TIMESTAMP NOT NULL default now(),
                       update_at TIMESTAMP NOT NULL default now(),
                       remark  VARCHAR(255) NOT NULL default '',
                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_role IS '角色';
COMMENT ON COLUMN o_role.id IS '唯一标识';
COMMENT ON COLUMN o_role.name IS '名称';

insert into o_role (id, name) values ('default_user','默认用户');

CREATE TABLE r_user_role(
                            id VARCHAR(255) NOT NULL,
                            user_id VARCHAR(255) NOT NULL,
                            role_id VARCHAR(255) NOT NULL,
                            PRIMARY KEY (id)
);

COMMENT ON TABLE r_user_role IS '用户角色关联';
COMMENT ON COLUMN r_user_role.id IS '唯一标识';
COMMENT ON COLUMN r_user_role.user_id IS '用户id';
COMMENT ON COLUMN r_user_role.role_id IS '角色id';
CREATE TABLE r_role_resource_operate(
                                        id VARCHAR(255) NOT NULL,
                                        role_id VARCHAR(255) NOT NULL,
                                        resource_id VARCHAR(255) NOT NULL,
                                        op VARCHAR(255) NOT NULL,
                                        filter_conditions text NOT NULL,
                                        PRIMARY KEY (id)
);

COMMENT ON TABLE r_role_resource_operate IS '角色资源操作关联';
COMMENT ON COLUMN r_role_resource_operate.id IS '唯一标识';
COMMENT ON COLUMN r_role_resource_operate.role_id IS '角色id';
COMMENT ON COLUMN r_role_resource_operate.resource_id IS '资源id';
COMMENT ON COLUMN r_role_resource_operate.op IS '操作（字典值）';
COMMENT ON COLUMN r_role_resource_operate.filter_conditions IS '过滤条件json字符串';


drop view if exists v_user_resource_op;
create view v_user_resource_op as
select t.*,p.type, p.api_url,p.data_conditions from
    (select r1.user_id,r2.resource_id,r2.op,r2.filter_conditions from r_user_role r1, r_role_resource_operate r2
     where r1.role_id=r2.role_id order by r2.op) t left join o_resource p on t.resource_id=p.id;

drop view if exists v_role_rel_detail;
create view v_role_rel_detail as
select r.*,pr.name role_name,
       pre.name resource_name,pre.module,pre.service_name,pre.service_name_cn,pre.type,pre.api_url,
       pre.data_conditions,pre.support_ops
from r_role_resource_operate r, o_role pr,o_resource pre where r.role_id=pr.id and r.resource_id=pre.id;

drop view if exists v_role_detail;
create view v_role_detail as
select *,
       (select array_to_json(array_agg(e)) from (select * from v_role_rel_detail where role_id=r.id and type=1) e) menu_resources,
       (select array_to_json(array_agg(e)) from (select * from v_role_rel_detail where role_id=r.id and type=3) e) func_resources,
       (select array_to_json(array_agg(e)) from (select * from v_role_rel_detail where role_id=r.id and type=2) e) api_resources,
       (select array_to_json(array_agg(e)) from (select * from v_role_rel_detail where role_id=r.id and type=4) e) data_resources
from o_role r;

drop view if exists v_role_with_resource_ids;
create view v_role_with_resource_ids as
select *,
       (select array_to_json(array_agg(resource_id)) from (select resource_id from v_role_rel_detail where role_id=r.id and type=1) e) menu_resource_ids,
       (select array_to_json(array_agg(resource_id)) from (select resource_id from v_role_rel_detail where role_id=r.id and type=3) e) func_resource_ids,
       (select array_to_json(array_agg(resource_id)) from (select resource_id from v_role_rel_detail where role_id=r.id and type=2) e) api_resource_ids,
       (select array_to_json(array_agg(resource_id)) from (select resource_id from v_role_rel_detail where role_id=r.id and type=4) e) data_resource_ids
from o_role r;

drop view if exists v_user_with_role;
create or replace view v_user_with_role as
select *,
       (select array_to_json(array_agg(e)) from
           (select * from (select r.id, r.user_id,r.role_id,pr.name from r_user_role r,o_role pr where r.role_id=pr.id) t where t.user_id=r.id) e) roles
from o_user r ;


drop view if exists v_user_with_menu;
create or replace view v_user_with_menu as
select *,
       (select array_to_json(array_agg(e)) from
           (select * from (select r.user_id,v.resource_id from r_user_role r,v_role_rel_detail v
                           where r.role_id=v.role_id and v.type=1) t where t.user_id=r.id) e) menu_ids
from o_user r ;

drop view if exists v_role_with_user;
create view v_role_with_user as
select r.id,pr.id role_id,pr.name role_name, pu.id user_id, pu.mobile,pu.email,pu.name,pu.create_at,pu.type,pu.org_id,pu.avatar_url,pu.status
from r_user_role r,o_role pr,o_user pu where r.role_id=pr.id and r.user_id =pu.id;



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists o_role cascade ;
drop table if exists o_resource cascade ;
drop table if exists o_user cascade ;

-- +goose StatementEnd
