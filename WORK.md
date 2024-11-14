Working notes

Bookmarks location: ~/Library/Application Support/Google/Chrome/[Default or UserProfileDir]/Bookmarks

Bookmarks structure
```json
{
  "checksum": "0bdb7d9af402b276bfe985699600410e",
  "roots": {
    "bookmark_bar": {
      "children": [ {
        "date_added": "13371505572849829",
        "date_last_used": "13371590325473788",
        "guid": "1c8a55b3-207d-4092-abe6-1edb898f9d25",
        "id": "2",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/jit/search3333",
        "type": "url",
        "url": "https://www.google.com/search?q=%2Fjit%2Fsearch"
      }, {
        "children": [ {
          "children": [  ],
          "date_added": "13376002887414186",
          "date_last_used": "0",
          "date_modified": "13376002887414186",
          "guid": "e6a9bccf-0b7e-422d-ae94-c544088b184a",
          "id": "13",
          "name": "test222",
          "type": "folder"
        } ],
        "date_added": "13376002873765072",
        "date_last_used": "0",
        "date_modified": "13376002887414273",
        "guid": "2a5f8c2d-a6dd-4b64-85f7-d7ef87e34a10",
        "id": "12",
        "name": "test",
        "type": "folder"
      } ],
      "date_added": "13314909169875218",
      "date_last_used": "0",
      "date_modified": "13376002873765446",
      "guid": "0bc5d13f-2cba-5d74-951f-3f233fe6c908",
      "id": "1",
      "name": "Bookmarks.json Bar",
      "type": "folder"
    },
    "other": {
      "children": [ {
        "date_added": "13371505572849829",
        "date_last_used": "13371590325473788",
        "guid": "1c8a55b3-207d-4092-abe6-1edb898f9d24",
        "id": "4",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/jit/search",
        "type": "url",
        "url": "https://www.google.com/search?q=%2Fjit%2Fsearch"
      }, {
        "date_added": "13371505590230167",
        "date_last_used": "0",
        "guid": "8cf732e6-2a77-41fa-b778-495c9dab974e",
        "id": "5",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/jit/board",
        "type": "url",
        "url": "https://www.google.com/search?q=%2Fjit%2Fboard"
      }, {
        "date_added": "13371505621682441",
        "date_last_used": "0",
        "guid": "d7c3b7df-285a-4ce8-8325-37c8725e3c07",
        "id": "6",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/grafana/postgres",
        "type": "url",
        "url": "https://www.google.com/search?q=%2Fgrafana%2Fpostgres"
      }, {
        "date_added": "13371505631551643",
        "date_last_used": "13371590323798905",
        "guid": "5ecc5822-a6e4-4e6d-bb98-78f8f36b847b",
        "id": "7",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/grafana/service",
        "type": "url",
        "url": "https://www.google.com/search?q=%2Fgrafana%2Fservice"
      }, {
        "date_added": "13371505653040825",
        "date_last_used": "0",
        "guid": "96e8fc7e-8a30-4c2b-bb81-833e9c23cde0",
        "id": "8",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/outlook",
        "type": "url",
        "url": "https://www.google.com/search?q=%2Foutlook"
      }, {
        "date_added": "13371506275713026",
        "date_last_used": "0",
        "guid": "97501b5c-dfb9-4681-91d2-814611bfe8d1",
        "id": "9",
        "meta_info": {
          "power_bookmark_meta": ""
        },
        "name": "/testupdated",
        "type": "url",
        "url": "https://www.youtube.com/"
      } ],
      "date_added": "13314909169875218",
      "date_last_used": "0",
      "date_modified": "13371506275713026",
      "guid": "82b081ec-3dd3-529c-8475-ab6c344590dd",
      "id": "3",
      "name": "Other Bookmarks.json",
      "type": "folder"
    },
    "synced": {
      "children": [  ],
      "date_added": "13314909169875219",
      "date_last_used": "0",
      "date_modified": "0",
      "guid": "4cf2e351-0e85-532b-bb37-df045d8f8d0f",
      "id": "10",
      "name": "Mobile Bookmarks.json",
      "type": "folder"
    }
  },
  "version": 1
}
```

Algorithm:
1. Read bookmarks
2. Add new entry into roots.other.imported
3. Update file