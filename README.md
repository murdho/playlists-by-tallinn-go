# :musical_score: Playlists by Tallinn :radio:

[Cloud Function](https://cloud.google.com/functions/) keeping an eye on tracks played in [Raadio Tallinn](https://raadiotallinn.err.ee/) and storing them in [Cloud Firestore](https://cloud.google.com/firestore/).

I enjoy the music of Raadio Tallinn and created this function for personal use. The data from Firestore can be used in many ways, for example to add new tracks to a playlist in Spotify.

Also a note about the implementation. I tried lazy initialization mostly for trying it out, it seems to perform similarly without laziness as well. :rocket: Next time around I would not do the lazy init. again if not absolutely necessary. I prefer simple and readable code over crazy-optimized code in most of the cases.

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
                            --timeout 2s
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
