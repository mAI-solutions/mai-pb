/// <reference path="../pb_data/types.d.ts" />


routerAdd("POST", "/proxy", (c) => {
  const model = new DynamicModel({
    url: ""
  })

  c.bind(model)

  const urlMatcherStr = String.raw`/(?:([^\:]*)\:\/\/)?(?:([^\:\@]*)(?:\:([^\@]*))?\@)?(?:([^\/\:]*)\.(?=[^\.\/\:]*\.[^\.\/\:]*))?([^\.\/\:]*)(?:\.([^\/\.\:]*))?(?:\:([0-9]*))?(\/[^\?#]*(?=.*?\/)\/)?([^\?#]*)?(?:\?([^#]*))?(?:#(.*))?/`
  const urlMatcher = new RegExp(urlMatcherStr)
  if (!urlMatcher.test(model.url)) {
    c.noContent(400)
    return
  }

  const res = $http.send({
    url:    model.url,
    method: "GET",
  })

  return c.string(200, res.raw)
})


