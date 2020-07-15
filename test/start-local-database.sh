db_response=""

while [[ ${db_response} != *"users"* ]]
do
    db_response=$(docker exec rest-go-users_db_1 /usr/bin/mysql --user=root --password=rootpasswd --execute "USE app; SHOW TABLES LIKE 'users';" 2> /dev/null)
done
