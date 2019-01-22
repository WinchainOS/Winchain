build:
	docker build -t winchainos/luckywin:v1.0.0 ../

rpc:
	docker run -d --name luckywin -p 9545:9545 -p 30303:30303 winchainos/luckywin:v1.0.0

bash:
    # 运行bash
	docker run -it --name luckywin winchainos/luckywin:v1.0.0 sh

kill:
	docker rm -f luckywin

init:
	docker exec -it luckywin /bin/entrypoint.sh init

start:
	docker exec -it luckywin /bin/entrypoint.sh start

stop:
	docker exec -it luckywin /bin/entrypoint.sh stop

test:
	curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x2831a3d42008a643eaa84d3547f6f77bfaa9a731", "latest"],"id":1}' localhost:9545
