package handlers

import "html/template"

var templates = template.Must(template.New("").Funcs(template.FuncMap{
    "FormatRupiah": FormatRupiah,
}).ParseGlob("templates/*.html"))
