# GOUSH (Golang url shortener)

Routes 

- GET /    => View all links
- GET /links/create    => Page create a link
- POST /links/create    => path to create a link
- GET /link/:shortlink => get Full url
- PUT /links/edit    => Page change a links short form 
- DELETE /links/delete/:shortlink    => delete a shortlink


### Models 

Short url {
    Title string,
    url string,
    shortlink,
    expiration date, // optional
}

### Future 
1. Can create,retrieve and delete links
2. Users can create,retrieve and delete their own links
3. Add admin panel to manage users
