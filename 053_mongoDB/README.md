# MongoDB

## install mongodb

segui tutte le istruzioni sul sito

## import drivers

- `go get gopgk.in/mgo.v2`
- `go get gopkg.in/mgo.v2/bson`

## collega il programma a mongodb

we need a mongo sessions

- the user controller must be able to connect to that session

```go
package user

type UserController struct {
    session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
    return &UserController{s}
}
```

```go
package main

func Session() (*mgo.Session, error) {
    s, err := mgo.Dial(`mongodb://localhost`)

    if err != nil {
        return nil, err
    }

    return s, nil
}
```

## CRUD with mongoDB

Come ID per i dati del model usa il tipo `bson.ObjectId`

```Go
func (uc UserController) CreateUser(w http.ReponseWriter, r *http.Request, _ httprouter.Params) {
    json.NewDecoder(r.Body).Decode(&u)

    u.Id = bson.NewObjectId()
    uc.session.DB("go-web-dev-db").C("users").Insert(u)
    uj, _ := json.Marshal(u)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
}
```
