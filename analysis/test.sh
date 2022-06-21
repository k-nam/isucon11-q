set -e

export target=go

# pushd ../webapp/go
# golangci-lint run -c ../../.github/.golangci.yml
# popd

DIR=$(pwd)/$(date +%Y-%m-%d_%H-%M-%S)
mkdir $DIR

cd ../development
docker compose -f docker-compose-bench.yml exec nginx bash -c ': > /var/log/nginx/access.log'
docker compose -f docker-compose-bench.yml exec mysql-backend bash -c ': > /var/log/mysql/mariadb-slow.log'

make run-bench 1> $DIR/bench.log 2>&1
docker compose -f docker-compose-bench.yml cp nginx:/var/log/nginx/access.log $DIR/nginx.log
docker compose -f docker-compose-bench.yml cp mysql-backend:/var/log/mysql/mariadb-slow.log $DIR/slowquery.log

cat $DIR/nginx.log | alp ltsv --sort sum -r -m '^/api/condition/.*$','^/api/isu/.+/icon$','^/api/isu/.+/graph$','^/assets/.*$','^/api/isu$','^/api/trend$','^/api/isu/.+','^/isu/.*' > $DIR/alp_result.log