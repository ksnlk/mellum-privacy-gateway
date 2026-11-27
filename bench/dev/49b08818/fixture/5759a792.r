require(httr)

files = list(
  `username` = 'davidwalsh',
  `password` = 'jrtdcnznp'
)

res <- httr::POST(url = 'http://domain.tld/post-to-me.php', body = files)
