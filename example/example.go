package main

import (
	auth "github.com/peiqi/goauth"
	provider "github.com/peiqi/goauth/provider"

	"fmt"
	"net/http"
)

var homepage = `
<html>
	<head>
		<title>Demo</title>
	</head>
	<body>
		<div>
			Fast Login:
		</div>
		<div style="display:block">
			<a href="/redict/github"><span style="background-image:url(icon/icon.svg);width:28px;height:28px;background-repeat:no-repeat;display:inline-block;background-position:-28px -28px;"></span></a>
		</div>
	</body>
</html>
`

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, homepage)
}

func GetRes(w http.ResponseWriter, r *http.Request, res provider.Result) {
	if res.Error != nil {
		fmt.Fprintf(w, "error:"+res.Error.GetError())
	} else {
		fmt.Fprintf(w, "Id:"+res.User.Id())
		fmt.Fprintf(w, ",Provider:"+res.User.Provider())
		fmt.Fprintf(w, ",Name:"+res.User.Name())
		fmt.Fprintf(w, ",Email:"+res.User.Email())
		fmt.Fprintf(w, ",Org:"+res.User.Org())
		fmt.Fprintf(w, ",Picture:"+res.User.Picture())
		fmt.Fprintf(w, ",Link:"+res.User.Link())
		fmt.Fprintf(w, ",Bio:"+res.User.Bio())
		fmt.Fprintf(w, ",Location:"+res.User.Location())
	}
}

func main() {
	auth.Auth(`{
		"github":{
			"id|key": "65c96526e1c87099ea6c",
			"secret": "16f60d4e69b7397cd12342cd6304f8ac4846dbc3",
			"callback": "/redict/github",
			"scope" : "user",
			"authres": "/redict/login"
		}
	}`)
	fmt.Println("goauth demo starting on port 8080")
	http.Handle("/icon/", http.StripPrefix("/icon/", http.FileServer(http.Dir("./icon"))))
	http.HandleFunc("/home", Home)
	http.HandleFunc("/redict/login", auth.ResHandlerFunc(GetRes))
	http.ListenAndServe(":8080", nil)
}

// ,
// 		"google":{
// 			"id|key": "842371556982-ke4vsrhh5se7p3900466ingqgm5lc13b.apps.googleusercontent.com",
// 			"secret": "ahM0zZf2KPFIo3nDDM40KnyU",
// 			"redict": "redict/google",
// 			"scope" : "user"
// 		}
