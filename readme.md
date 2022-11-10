# get
## url
/url/get?id=
### axios
axios.get("http://127.0.0.1:4000/url/get?id=" + id)
# post
## url
/url/create
### axios
axios.post("http://127.0.0.1:4000/url/create", {
    headers: { "Content-Type": "application/json;"}
    ,body: JSON.stringify({"url":url})
})