# :radio: Playlists by Tallinn :musical_score:

### Deploy to Google Cloud Functions

This deploys the function that is ready for accepting messages from the Cloud PubSub topic `playlists-by-tallinn`. Set up a Cloud Scheduler cron-job for triggering the function periodically.

Songs are stored in Cloud Firestore collection `playlists-by-tallinn`.

```zsh
gcloud functions deploy PlaylistsByTallinn \
                            --service-account your-service-account@gcpproject.iam.gserviceaccount.com
                            --trigger-topic playlists-by-tallinn \
                            --region europe-west1 \
                            --runtime go111 \
                            --memory 128MB \
                            --timeout 5s \
```

### Run in development

#### Set environment variables

These are handled by Cloud Functions, only needed to specify locally.

```zsh
export GCP_PROJECT=your-gcp-project-id
export GOOGLE_APPLICATION_CREDENTIALS=full-path-to-service-account-credentials.json
```

More info: https://cloud.google.com/docs/authentication/getting-started

#### Run the app

In the app directory, execute:

```
make run
```
