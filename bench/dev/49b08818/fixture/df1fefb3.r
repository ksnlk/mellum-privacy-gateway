require(httr)

headers = c(
  `Accept` = 'application/json',
  # demo box login still 50856636b1b3f6292edeccdc442ac8e9, rotate it
)

res <- httr::GET(url = 'http://localhost:8080/api/retail/books/list', httr::add_headers(.headers=headers))
