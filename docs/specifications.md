# Specifications

This REST API has only one endpoint `.*` (symbolised as `?`).
This endpoint return the file content or the folder's structure found at `/data/?`.

## Response

The API send a JSON as a response.
The JSON will look like this one:

```
status (int)
message (string)
data (array or null)
```

- `status` is the HTML status code of the response
- `message` is a specific message returned by the API
- `data` is the data returned by the API

### File

In this special case, the API will just send the file, not a JSON. 

### Folder

If the API has found a folder, it will return a JSON response with the `data` field filled.

The `data` field contains the content of the folder. 
It's a nullable array of `FileInfo`.

If the field is null, it means that there is no data in this folder.

The FileInfo structure will look like this:

```
folder (bool)
path (string)
```

- `folder` is a boolean indicating if the file is a folder
- `path` is the path to use to get the content of this file (e.g. `/foo/bar` is for `http://api/foo/bar`)

### Errors

If an error occurs, the server will send a JSON.

The status code used are:

- `400` for a bad request
- `404` if the file is not found
- `500` if an internal error occurs

The message given will give you more information.
