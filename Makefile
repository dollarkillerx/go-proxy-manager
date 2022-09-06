gen_agent_proto:
	protoc --proto_path=proto/agent --go_out=proto/agent --go_opt=paths=source_relative --go-grpc_out=proto/agent --go-grpc_opt=paths=source_relative proto/agent/*proto
	#protoc -I proto/agent/ proto/agent/*.proto --go-grpc_out=proto/agent/.