#docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
#docker-compose up --build

#SELECT 'CREATE DATABASE offiworks' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'offiworks')
#psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'offiworks'" | grep -q 1 | psql -U postgres -c "CREATE DATABASE offiworks"
#eval $(minikube docker-env)