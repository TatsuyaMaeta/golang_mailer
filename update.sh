# コンテナ停止 コンテナ削除 再ビルド
container=go_docker_ec2-db

# ソースコードの反映をさせたいときはこの内容で反映される
docker compose down -v
# docker stop $container
# docker container rm $container
docker compose up -d --build