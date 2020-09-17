#!/bin/bash

DATABASE_PARAMS=''

case $DATABASE_TYPE in
    mysql)
        echo "Configured with mysql database"
        DATABASE_PARAMS="${DATABASE_PARAMS}-Dspring.datasource.url=jdbc:mysql://$DATABASE_URL "
        ;;
    h2)
        echo "Configured with his h2 embedded database"
        DATABASE_PARAMS="${DATABASE_PARAMS}-Dspring.datasource.url=jdbc:h2:$DATABASE_URL;DB_CLOSE_DELAY=-1 "
        ;;
    *)
        echo "ERROR: please specify a DATABASE_TYPE env var with the target database type"
        exit 1
        ;;
esac

if [[ ! -z $DATABASE_USER ]]
then
    DATABASE_PARAMS="${DATABASE_PARAMS}-Dspring.datasource.username=$DATABASE_USER "
fi

if [[ ! -z $DATABASE_PASSWORD ]]
then
    DATABASE_PARAMS="${DATABASE_PARAMS}-Dspring.datasource.password=$DATABASE_PASSWORD "
fi

if [[ ! -z $DATA_DB_MODE ]]
then
    DATABASE_PARAMS="${DATABASE_PARAMS}-Dspring.jpa.properties.hibernate.hbm2ddl.auto=$DATA_DB_MODE "
fi

java $JAVAOPS \
        -Dspring.profiles.active=$DATABASE_TYPE \
        $DATABASE_PARAMS \
        -jar app.jar
