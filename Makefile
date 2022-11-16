run:
	docker compose down && docker compose up -d --build

deploy:
	cd deployment &&  kubectl apply -f api-gateway-service.yaml,auth-service-deployment.yaml,auth-service-service.yaml,cache-deployment.yaml,cache-persistentvolumeclaim.yaml,cache-service.yaml,db-networkpolicy.yaml,mongodb-data-container-persistentvolumeclaim.yaml,mongodb-deployment.yaml,mongodb-service.yaml,user-service-deployment.yaml,user-service-service.yaml