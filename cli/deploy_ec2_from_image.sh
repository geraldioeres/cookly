# Use the bash terminal for window and make sure have login the AWS RDS
docker pull geraldiors/cookly:1.0.0
docker run -d -p 8080:8080 --name cookly geraldiors/cookly:1.0.0