/// <reference path="../pb_data/types.d.ts" />
routerAdd("GET", "/proxy", (c) => {
  const url = c.queryParam("url")
  const res = $http.send({
    url:    url,
    method: "GET",
  })

  return c.string(200, res.raw)
})


