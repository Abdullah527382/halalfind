package views

type Template struct {
  template *template.Template
}

func ParseFiles(files ...string) (*Template, error) {
  tpl, err := template.New("").Funcs(template.FuncMap{
    "flash": func() string { return "" },
  }).ParseFiles(files...)
  if err != nil {
    return nil, fmt.Errorf("parsing view files: %w", err)
  }
  return &Template{
    template: tpl,
  }, nil
}

func (t *Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
  // clone the template BEFORE adding user-specifics!
  clone := t.clone()

  // Get the flash cookie
	cookie, err := r.Cookie("flash")
	switch err {
  case nil:
    // override the flash msg
    clone.template = clone.template.Funcs(template.FuncMap{
      "flash": func() string { return cookie.Value },
    })
		// Delete the flash cookie so we don't repeat it.
		expired := http.Cookie{ ... }
		http.SetCookie(w, &expired)
	default:
		// noop
	}

  err = clone.template.Execute(w, data)
	if err != nil { ... }
}