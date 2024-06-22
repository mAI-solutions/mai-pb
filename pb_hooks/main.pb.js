/// <reference path="../pb_data/types.d.ts" />
routerAdd("POST", "/proxy", (c) => {
  const model = new DynamicModel({
    url: ""
  })

  c.bind(model)

  try {
    new URL(model.url);
  } catch (e) {
    c.noContent(400)
    return
  }

  const res = $http.send({
    url:    model.url,
    method: "GET",
  })

  return c.string(200, res.raw)
})


