# ENDPOINTS
### Backend hosts
* Golang backend_host = localhost:8081
### Common
1. *Create news* POST{backend_host}/news
  * body - "count":{count_news}
  * count_news=500000 is a max value per request
1. *Delete news* DELETE{backend_host}/news
### DB side aggregation
1. *Get min, max, avg values for the rating* GET{backend_host}/dbsa/news/min-max-avg-rating
1. *Get grouped per month JSON data* GET{backend_host}/dbsa/news/per-month-json-data
### Server side aggregation
1. *Get min, max, avg values for the rating* GET{backend_host}/ssa/news/min-max-avg-rating?use_rows={bool}
  * use_rows = true|false or ignore it, will work with models.News entities unless use_rows=true
