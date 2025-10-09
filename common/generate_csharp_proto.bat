set out_put_dir=generated_csharp

if not exist %out_put_dir% (
    mkdir %out_put_dir%
)

protoc --proto_path=./proto --csharp_out=%out_put_dir% ./proto/*.proto