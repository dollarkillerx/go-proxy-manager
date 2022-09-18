gen_proto_common:
	protoc --proto_path=proto/ --go_out=proto/ --go_opt=paths=source_relative --go-grpc_out=proto/ --go-grpc_opt=paths=source_relative proto/common/*proto


gen_proto_agent:
	protoc --proto_path=proto/ --go_out=proto/ --go_opt=paths=source_relative --go-grpc_out=proto/ --go-grpc_opt=paths=source_relative proto/agent/*proto


gen_proto_backend:
	protoc --proto_path=proto/ --go_out=proto/ --go_opt=paths=source_relative --go-grpc_out=proto/ --go-grpc_opt=paths=source_relative proto/backend/*proto


# --proto_path=proto/ 项目根路径 --go_out=proto/ --go-grpc_out=proto/ 指定生成文件目录