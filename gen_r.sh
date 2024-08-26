dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=r_user_role,r_role_resource_operate --model_naming "{{ toUpperCamelCase ( replace . \"r_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"r_\" \"\") }}" \
--module authz-service --api true

