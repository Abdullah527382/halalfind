package controllers

// // Take note: our UserHandler has a UserStore injected in!
// // Global data stores can be problematic long term.
// type UserHandler struct {
// 	Users    models.UserStore
// 	Sessions models.SessionStore
//   }
  
//   func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil { ... }
  
// 	var form struct {
// 	  Email    string `schema:"email"`
// 	  Password string `schema:"password"`
// 	}
// 	schema.NewDecoder().Decode(&form, r.PostForm)
// 	user := sql.User{
// 	  Email: form.Email,
// 	  Password: form.Password,
// 	}
// 	err = uh.Users.Create(&user)
// 	if err != nil { ... }
  
// 	token, err := uh.Session.Create(&user)
// 	if err != nil { ... }
// 	// create session cookie using token
  
// 	http.Redirect(w, r, "/dashboard", http.StatusFound)
//   }