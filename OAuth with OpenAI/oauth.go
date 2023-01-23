package openai_oauth

import "golang.org/x/oauth2"


config := &oauth2.Config{
    ClientID:     "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET",
    Endpoint: oauth2.Endpoint{
        AuthURL:  "https://api.openai.com/v1/oauth/authorize",
        TokenURL: "https://api.openai.com/v1/oauth/token",
    },
}

token, err := config.Exchange(context.TODO(), "AUTH_CODE")
if err != nil {
    log.Fatal(err)
}

client := config.Client(context.TODO(), token)

// Use the client to make API requests
response, err := client.Get("https://api.openai.com/v1/engines/davinci-codex/completions")
if err != nil {
    log.Fatal(err)
}
