#!/usr/bin/env bash
docker run --name esqimo-news-app-db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=esqimo -e MYSQL_PASSWORD=password -e MYSQL_DATABASE=news -p 3307:3306 --tmpfs /var/lib/mysql mysql:8
