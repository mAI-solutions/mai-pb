/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("hygqyjxm6mvk4io")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vhkemnhw",
    "name": "field",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("hygqyjxm6mvk4io")

  // remove
  collection.schema.removeField("vhkemnhw")

  return dao.saveCollection(collection)
})
