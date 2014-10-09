curl http://localhost:8983/solr/resource/foods.reviews_by_day/solrconfig.xml --data-binary @solrconfig.xml -H 'Content-type:text/xml;'
curl http://localhost:8983/solr/resource/foods.reviews_by_day/schema.xml --data-binary @schema.xml -H 'Content-type:text/xml;'
curl "http://localhost:8983/solr/admin/cores?action=CREATE&name=foods.reviews_by_day"
