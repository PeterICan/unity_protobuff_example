set out_put_dir=generated_go

if not exist %out_put_dir% (
    mkdir %out_put_dir%
)

protoc --proto_path=./proto --go_out=%out_put_dir% ./proto/*.proto