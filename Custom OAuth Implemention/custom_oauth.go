package custom_oauth

import (
    "context"
    "fmt"
    "net/http"
    "golang.org/x/oauth2"
)

config := &oauth2.Config{
    ClientID:     "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET",
    RedirectURL:  "YOUR_REDIRECT_URI",
    Endpoint: oauth2.Endpoint{
        AuthURL:  "YOUR_AUTHORIZATION_URL",
        TokenURL: "YOUR_TOKEN_URL",
    },
    Scopes: []string{"YOUR_SCOPES"},
}

//generates the OAuth2 URL where users can grant access to application

func generateAuthURL(w http.ResponseWriter, r *http.Request) {
    url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusFound)
}


// handles the callback from the OAuth provider

func handleCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    token, err := config.Exchange(context.TODO(), code)
    if err != nil {
        http.Error(w, "Error exchanging token", http.StatusInternalServerError)
        return
    }
    // Use the token to authenticate with the OAuth provider
    client := config.Client(context.TODO(), token)
    response, err := client.Get("YOUR_RESOURCE_URL")
    if err != nil {
        http.Error(w, "Error fetching resource", http.StatusInternalServerError)
        return
    }
    // Handle the response
    fmt.Fprintf(w, "Response: %s", response)
}

// ToDo functions: handle token expiration, refresh token and error handling.

//handles the corresponding routes in server
http.HandleFunc("/login", generateAuthURL)
http.HandleFunc("/callback", handleCallback)

// http.ListenAndServe(":8000", nil)

