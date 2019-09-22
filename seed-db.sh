#!/usr/bin/env bash
echo SEEDING THE DB
mysql --host=127.0.0.1 --port=3307 --user=esqimo --password=password  news < seed-db-with-sources
echo LOADING THE ARTICLES...
curl -X POST http://localhost:8000/load-news
echo DONE!