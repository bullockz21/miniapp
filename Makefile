all: run

clean: rm_img

rm_con:
	sudo docker rm miniapp_pg miniapp

rm_img:
	sudo docker rmi miniapp-miniapp

run:
	# cd ./backend/ && make build && cd ..
	sudo docker-compose up -d

stop:
	sudo docker-compose down -v

.PHONY: all clean rm_con rm_img run stop