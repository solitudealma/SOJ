########## api ##########
api-gen:
	goctl api go -api desc/*.api -dir . --home ../../../../common/goctl --style=go_zero

########## rpc ##########
rpc-gen:
	goctl rpc protoc *.protoc --go_out=. --go-grpc_out=.  --zrpc_out=. --home ../../../../common/goctl --style=go_zero

########## model ##########
model-gen:
	/bin/bash ${PWD}/scripts/model_gen.sh $d $t
