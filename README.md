# go-serverless

Go serverless functions examples with most popular Cloud Providers 

### Creating zip archive

```shell
go mod download

go build ./cmd/<aws|gcp>

zip -r handler ./<aws|gcp>
```

### GCP

For Google Cloud Functions you could use inline editor to try use `.zip` archive option.

Refer to the <a href="https://cloud.google.com/functions/docs/deploying/console">docs</a>

### AWS

For AWS Lambda you could use inline editor to try use `.zip` archive option.

Refer to the <a href="https://cloud.google.com/functions/docs/deploying/console">docs</a>
