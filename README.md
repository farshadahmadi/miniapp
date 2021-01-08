# Mini app to investigate split sdk synchronization issue

## To run the project:

1. Add your split key to the `config/local/local.yml` file

```
  FeatureFlag:
    APIKey: <put-your-key>
```

2. Then run `go run cmd/main.go`

3. Run the following command to get a treatment for a specific split. Replace {split-name} with the split name of your choice.

```
curl -X POST http://localhost:8050/v1/feature/{split-name} -H 'Content-Type: application/json' -d '{"trafficId":"B7F6E54F-31F2-442B-AD74-EE044971B6BJ"}' 
```

## How split sdk is used?

miniapp server provides client with an an api `/v1/feature/{split-name}` to get treatments for a specific split. The [api handler](https://github.com/farshadahmadi/kosmo/blob/d4891a71075d7d671bc7f7e4419262fb50b110ba/internal/app/app.go#L102) uses [decisions package](internal/feature-flag/decisions) to fetch treatment. decision package itself uses [feature package](internal/feature-flag/features) that leverages split sdk to communicate with split.io.
