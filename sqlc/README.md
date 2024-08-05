migrate create -ext=sql -dir=sql/migrations -seq init => comando para criar a migration

migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up => comando para rodar a migration

migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down => comando para desfazer a migration

____SQLC____
sudo snap install sqlc => instalação do sqlc

sqlc generate => comando para gerar os arquivos sqlc