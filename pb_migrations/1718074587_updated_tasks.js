/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rbu353iptlp3fc6")

  collection.name = "feed"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rbu353iptlp3fc6")

  collection.name = "tasks"

  return dao.saveCollection(collection)
})
