# Star Tracker

This service is designed to read `/api/v1/game/${GAME_ID}/feed/live` data and exports the star player data

## Deployment

(tidy this into a build/make step)

- `make`
- `docker push trelore/star-tracker:${VERSION}`
- `pachctl delete-pipeline star-tracker`
- `make pipeline` (make sure to update the `star-tracker.json` file)
